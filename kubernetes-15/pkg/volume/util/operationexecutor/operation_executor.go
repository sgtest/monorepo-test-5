/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package operationexecutor implements interfaces that enable execution of
// attach, detach, mount, and unmount operations with a
// nestedpendingoperations so that more than one operation is never triggered
// on the same volume for the same pod.
package operationexecutor

import (
	"strings"
	"time"

	"github.com/golang/glog"

	"k8s.io/apimachinery/pkg/types"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-15/pkg/api/v1"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-15/pkg/util/mount"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-15/pkg/volume"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-15/pkg/volume/util/nestedpendingoperations"
	volumetypes "github.com/sourcegraph/monorepo-test-1/kubernetes-15/pkg/volume/util/types"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-15/pkg/volume/util/volumehelper"
)

// OperationExecutor defines a set of operations for attaching, detaching,
// mounting, or unmounting a volume that are executed with a NewNestedPendingOperations which
// prevents more than one operation from being triggered on the same volume.
//
// These operations should be idempotent (for example, AttachVolume should
// still succeed if the volume is already attached to the node, etc.). However,
// they depend on the volume plugins to implement this behavior.
//
// Once an operation completes successfully, the actualStateOfWorld is updated
// to indicate the volume is attached/detached/mounted/unmounted.
//
// If the OperationExecutor fails to start the operation because, for example,
// an operation with the same UniqueVolumeName is already pending, a non-nil
// error is returned.
//
// Once the operation is started, since it is executed asynchronously,
// errors are simply logged and the goroutine is terminated without updating
// actualStateOfWorld (callers are responsible for retrying as needed).
//
// Some of these operations may result in calls to the API server; callers are
// responsible for rate limiting on errors.
type OperationExecutor interface {
	// AttachVolume attaches the volume to the node specified in volumeToAttach.
	// It then updates the actual state of the world to reflect that.
	AttachVolume(volumeToAttach VolumeToAttach, actualStateOfWorld ActualStateOfWorldAttacherUpdater) error

	// VerifyVolumesAreAttachedPerNode verifies the given list of volumes to see whether they are still attached to the node.
	// If any volume is not attached right now, it will update the actual state of the world to reflect that.
	// Note that this operation could be operated concurrently with other attach/detach operations.
	// In theory (but very unlikely in practise), race condition among these operations might mark volume as detached
	// even if it is attached. But reconciler can correct this in a short period of time.
	VerifyVolumesAreAttachedPerNode(AttachedVolumes []AttachedVolume, nodeName types.NodeName, actualStateOfWorld ActualStateOfWorldAttacherUpdater) error

	// VerifyVolumesAreAttached verifies volumes being used in entire cluster and if they are still attached to the node
	// If any volume is not attached right now, it will update actual state of world to reflect that.
	VerifyVolumesAreAttached(volumesToVerify map[types.NodeName][]AttachedVolume, actualStateOfWorld ActualStateOfWorldAttacherUpdater)

	// DetachVolume detaches the volume from the node specified in
	// volumeToDetach, and updates the actual state of the world to reflect
	// that. If verifySafeToDetach is set, a call is made to the fetch the node
	// object and it is used to verify that the volume does not exist in Node's
	// Status.VolumesInUse list (operation fails with error if it is).
	DetachVolume(volumeToDetach AttachedVolume, verifySafeToDetach bool, actualStateOfWorld ActualStateOfWorldAttacherUpdater) error

	// MountVolume mounts the volume to the pod specified in volumeToMount.
	// Specifically it will:
	// * Wait for the device to finish attaching (for attachable volumes only).
	// * Mount device to global mount path (for attachable volumes only).
	// * Update actual state of world to reflect volume is globally mounted (for
	//   attachable volumes only).
	// * Mount the volume to the pod specific path.
	// * Update actual state of world to reflect volume is mounted to the pod
	//   path.
	MountVolume(waitForAttachTimeout time.Duration, volumeToMount VolumeToMount, actualStateOfWorld ActualStateOfWorldMounterUpdater) error

	// UnmountVolume unmounts the volume from the pod specified in
	// volumeToUnmount and updates the actual state of the world to reflect that.
	UnmountVolume(volumeToUnmount MountedVolume, actualStateOfWorld ActualStateOfWorldMounterUpdater) error

	// UnmountDevice unmounts the volumes global mount path from the device (for
	// attachable volumes only, freeing it for detach. It then updates the
	// actual state of the world to reflect that.
	UnmountDevice(deviceToDetach AttachedVolume, actualStateOfWorld ActualStateOfWorldMounterUpdater, mounter mount.Interface) error

	// VerifyControllerAttachedVolume checks if the specified volume is present
	// in the specified nodes AttachedVolumes Status field. It uses kubeClient
	// to fetch the node object.
	// If the volume is found, the actual state of the world is updated to mark
	// the volume as attached.
	// If the volume does not implement the attacher interface, it is assumed to
	// be attached and the actual state of the world is updated accordingly.
	// If the volume is not found or there is an error (fetching the node
	// object, for example) then an error is returned which triggers exponential
	// back off on retries.
	VerifyControllerAttachedVolume(volumeToMount VolumeToMount, nodeName types.NodeName, actualStateOfWorld ActualStateOfWorldAttacherUpdater) error

	// IsOperationPending returns true if an operation for the given volumeName and podName is pending,
	// otherwise it returns false
	IsOperationPending(volumeName v1.UniqueVolumeName, podName volumetypes.UniquePodName) bool
}

// NewOperationExecutor returns a new instance of OperationExecutor.
func NewOperationExecutor(
	operationGenerator OperationGenerator) OperationExecutor {

	return &operationExecutor{
		pendingOperations: nestedpendingoperations.NewNestedPendingOperations(
			true /* exponentialBackOffOnError */),
		operationGenerator: operationGenerator,
	}
}

// ActualStateOfWorldMounterUpdater defines a set of operations updating the actual
// state of the world cache after successful mount/unmount.
type ActualStateOfWorldMounterUpdater interface {
	// Marks the specified volume as mounted to the specified pod
	MarkVolumeAsMounted(podName volumetypes.UniquePodName, podUID types.UID, volumeName v1.UniqueVolumeName, mounter volume.Mounter, outerVolumeSpecName string, volumeGidValue string) error

	// Marks the specified volume as unmounted from the specified pod
	MarkVolumeAsUnmounted(podName volumetypes.UniquePodName, volumeName v1.UniqueVolumeName) error

	// Marks the specified volume as having been globally mounted.
	MarkDeviceAsMounted(volumeName v1.UniqueVolumeName) error

	// Marks the specified volume as having its global mount unmounted.
	MarkDeviceAsUnmounted(volumeName v1.UniqueVolumeName) error
}

// ActualStateOfWorldAttacherUpdater defines a set of operations updating the
// actual state of the world cache after successful attach/detach/mount/unmount.
type ActualStateOfWorldAttacherUpdater interface {
	// Marks the specified volume as attached to the specified node.  If the
	// volume name is supplied, that volume name will be used.  If not, the
	// volume name is computed using the result from querying the plugin.
	//
	// TODO: in the future, we should be able to remove the volumeName
	// argument to this method -- since it is used only for attachable
	// volumes.  See issue 29695.
	MarkVolumeAsAttached(volumeName v1.UniqueVolumeName, volumeSpec *volume.Spec, nodeName types.NodeName, devicePath string) error

	// Marks the specified volume as detached from the specified node
	MarkVolumeAsDetached(volumeName v1.UniqueVolumeName, nodeName types.NodeName)

	// Marks desire to detach the specified volume (remove the volume from the node's
	// volumesToReportedAsAttached list)
	RemoveVolumeFromReportAsAttached(volumeName v1.UniqueVolumeName, nodeName types.NodeName) error

	// Unmarks the desire to detach for the specified volume (add the volume back to
	// the node's volumesToReportedAsAttached list)
	AddVolumeToReportAsAttached(volumeName v1.UniqueVolumeName, nodeName types.NodeName)
}

// VolumeToAttach represents a volume that should be attached to a node.
type VolumeToAttach struct {
	// VolumeName is the unique identifier for the volume that should be
	// attached.
	VolumeName v1.UniqueVolumeName

	// VolumeSpec is a volume spec containing the specification for the volume
	// that should be attached.
	VolumeSpec *volume.Spec

	// NodeName is the identifier for the node that the volume should be
	// attached to.
	NodeName types.NodeName

	// scheduledPods is a map containing the set of pods that reference this
	// volume and are scheduled to the underlying node. The key in the map is
	// the name of the pod and the value is a pod object containing more
	// information about the pod.
	ScheduledPods []*v1.Pod
}

// VolumeToMount represents a volume that should be attached to this node and
// mounted to the PodName.
type VolumeToMount struct {
	// VolumeName is the unique identifier for the volume that should be
	// mounted.
	VolumeName v1.UniqueVolumeName

	// PodName is the unique identifier for the pod that the volume should be
	// mounted to after it is attached.
	PodName volumetypes.UniquePodName

	// VolumeSpec is a volume spec containing the specification for the volume
	// that should be mounted. Used to create NewMounter. Used to generate
	// InnerVolumeSpecName.
	VolumeSpec *volume.Spec

	// outerVolumeSpecName is the podSpec.Volume[x].Name of the volume. If the
	// volume was referenced through a persistent volume claim, this contains
	// the podSpec.Volume[x].Name of the persistent volume claim.
	OuterVolumeSpecName string

	// Pod to mount the volume to. Used to create NewMounter.
	Pod *v1.Pod

	// PluginIsAttachable indicates that the plugin for this volume implements
	// the volume.Attacher interface
	PluginIsAttachable bool

	// VolumeGidValue contains the value of the GID annotation, if present.
	VolumeGidValue string

	// DevicePath contains the path on the node where the volume is attached.
	// For non-attachable volumes this is empty.
	DevicePath string

	// ReportedInUse indicates that the volume was successfully added to the
	// VolumesInUse field in the node's status.
	ReportedInUse bool
}

// AttachedVolume represents a volume that is attached to a node.
type AttachedVolume struct {
	// VolumeName is the unique identifier for the volume that is attached.
	VolumeName v1.UniqueVolumeName

	// VolumeSpec is the volume spec containing the specification for the
	// volume that is attached.
	VolumeSpec *volume.Spec

	// NodeName is the identifier for the node that the volume is attached to.
	NodeName types.NodeName

	// PluginIsAttachable indicates that the plugin for this volume implements
	// the volume.Attacher interface
	PluginIsAttachable bool

	// DevicePath contains the path on the node where the volume is attached.
	// For non-attachable volumes this is empty.
	DevicePath string
}

// MountedVolume represents a volume that has successfully been mounted to a pod.
type MountedVolume struct {
	// PodName is the unique identifier of the pod mounted to.
	PodName volumetypes.UniquePodName

	// VolumeName is the unique identifier of the volume mounted to the pod.
	VolumeName v1.UniqueVolumeName

	// InnerVolumeSpecName is the volume.Spec.Name() of the volume. If the
	// volume was referenced through a persistent volume claims, this contains
	// the name of the bound persistent volume object.
	// It is the name that plugins use in their pod mount path, i.e.
	// /var/lib/kubelet/pods/{podUID}/volumes/{escapeQualifiedPluginName}/{innerVolumeSpecName}/
	// PVC example,
	//   apiVersion: v1
	//   kind: PersistentVolume
	//   metadata:
	//     name: pv0003				<- InnerVolumeSpecName
	//   spec:
	//     capacity:
	//       storage: 5Gi
	//     accessModes:
	//       - ReadWriteOnce
	//     persistentVolumeReclaimPolicy: Recycle
	//     nfs:
	//       path: /tmp
	//       server: 172.17.0.2
	// Non-PVC example:
	//   apiVersion: v1
	//   kind: Pod
	//   metadata:
	//     name: test-pd
	//   spec:
	//     containers:
	//     - image: gcr.io/google_containers/test-webserver
	//     	 name: test-container
	//     	 volumeMounts:
	//     	 - mountPath: /test-pd
	//     	   name: test-volume
	//     volumes:
	//     - name: test-volume			<- InnerVolumeSpecName
	//     	 gcePersistentDisk:
	//     	   pdName: my-data-disk
	//     	   fsType: ext4
	InnerVolumeSpecName string

	// outerVolumeSpecName is the podSpec.Volume[x].Name of the volume. If the
	// volume was referenced through a persistent volume claim, this contains
	// the podSpec.Volume[x].Name of the persistent volume claim.
	// PVC example:
	//   kind: Pod
	//   apiVersion: v1
	//   metadata:
	//     name: mypod
	//   spec:
	//     containers:
	//       - name: myfrontend
	//         image: dockerfile/nginx
	//         volumeMounts:
	//         - mountPath: "/var/www/html"
	//           name: mypd
	//     volumes:
	//       - name: mypd				<- OuterVolumeSpecName
	//         persistentVolumeClaim:
	//           claimName: myclaim
	// Non-PVC example:
	//   apiVersion: v1
	//   kind: Pod
	//   metadata:
	//     name: test-pd
	//   spec:
	//     containers:
	//     - image: gcr.io/google_containers/test-webserver
	//     	 name: test-container
	//     	 volumeMounts:
	//     	 - mountPath: /test-pd
	//     	   name: test-volume
	//     volumes:
	//     - name: test-volume			<- OuterVolumeSpecName
	//     	 gcePersistentDisk:
	//     	   pdName: my-data-disk
	//     	   fsType: ext4
	OuterVolumeSpecName string

	// PluginName is the "Unescaped Qualified" name of the volume plugin used to
	// mount and unmount this volume. It can be used to fetch the volume plugin
	// to unmount with, on demand. It is also the name that plugins use, though
	// escaped, in their pod mount path, i.e.
	// /var/lib/kubelet/pods/{podUID}/volumes/{escapeQualifiedPluginName}/{outerVolumeSpecName}/
	PluginName string

	// PodUID is the UID of the pod mounted to. It is also the string used by
	// plugins in their pod mount path, i.e.
	// /var/lib/kubelet/pods/{podUID}/volumes/{escapeQualifiedPluginName}/{outerVolumeSpecName}/
	PodUID types.UID

	// Mounter is the volume mounter used to mount this volume. It is required
	// by kubelet to create container.VolumeMap.
	Mounter volume.Mounter

	// VolumeGidValue contains the value of the GID annotation, if present.
	VolumeGidValue string
}

type operationExecutor struct {
	// pendingOperations keeps track of pending attach and detach operations so
	// multiple operations are not started on the same volume
	pendingOperations nestedpendingoperations.NestedPendingOperations

	// operationGenerator is an interface that provides implementations for
	// generating volume function
	operationGenerator OperationGenerator
}

func (oe *operationExecutor) IsOperationPending(volumeName v1.UniqueVolumeName, podName volumetypes.UniquePodName) bool {
	return oe.pendingOperations.IsOperationPending(volumeName, podName)
}

func (oe *operationExecutor) AttachVolume(
	volumeToAttach VolumeToAttach,
	actualStateOfWorld ActualStateOfWorldAttacherUpdater) error {
	attachFunc, err :=
		oe.operationGenerator.GenerateAttachVolumeFunc(volumeToAttach, actualStateOfWorld)
	if err != nil {
		return err
	}

	return oe.pendingOperations.Run(
		volumeToAttach.VolumeName, "" /* podName */, attachFunc)
}

func (oe *operationExecutor) DetachVolume(
	volumeToDetach AttachedVolume,
	verifySafeToDetach bool,
	actualStateOfWorld ActualStateOfWorldAttacherUpdater) error {
	detachFunc, err :=
		oe.operationGenerator.GenerateDetachVolumeFunc(volumeToDetach, verifySafeToDetach, actualStateOfWorld)
	if err != nil {
		return err
	}

	return oe.pendingOperations.Run(
		volumeToDetach.VolumeName, "" /* podName */, detachFunc)
}
func (oe *operationExecutor) VerifyVolumesAreAttached(
	attachedVolumes map[types.NodeName][]AttachedVolume,
	actualStateOfWorld ActualStateOfWorldAttacherUpdater) {

	// A map of plugin names and nodes on which they exist with volumes they manage
	bulkVerifyPluginsByNode := make(map[string]map[types.NodeName][]*volume.Spec)
	volumeSpecMapByPlugin := make(map[string]map[*volume.Spec]v1.UniqueVolumeName)

	for node, nodeAttachedVolumes := range attachedVolumes {
		for _, volumeAttached := range nodeAttachedVolumes {
			if volumeAttached.VolumeSpec == nil {
				glog.Errorf("VerifyVolumesAreAttached: nil spec for volume %s", volumeAttached.VolumeName)
				continue
			}
			volumePlugin, err :=
				oe.operationGenerator.GetVolumePluginMgr().FindPluginBySpec(volumeAttached.VolumeSpec)

			if err != nil || volumePlugin == nil {
				glog.Errorf(
					"VolumesAreAttached.FindPluginBySpec failed for volume %q (spec.Name: %q) on node %q with error: %v",
					volumeAttached.VolumeName,
					volumeAttached.VolumeSpec.Name(),
					volumeAttached.NodeName,
					err)
				continue
			}

			pluginName := volumePlugin.GetPluginName()

			if volumePlugin.SupportsBulkVolumeVerification() {
				pluginNodes, pluginNodesExist := bulkVerifyPluginsByNode[pluginName]

				if !pluginNodesExist {
					pluginNodes = make(map[types.NodeName][]*volume.Spec)
				}

				volumeSpecList, nodeExists := pluginNodes[node]
				if !nodeExists {
					volumeSpecList = []*volume.Spec{}
				}
				volumeSpecList = append(volumeSpecList, volumeAttached.VolumeSpec)
				pluginNodes[node] = volumeSpecList

				bulkVerifyPluginsByNode[pluginName] = pluginNodes
				volumeSpecMap, mapExists := volumeSpecMapByPlugin[pluginName]

				if !mapExists {
					volumeSpecMap = make(map[*volume.Spec]v1.UniqueVolumeName)
				}
				volumeSpecMap[volumeAttached.VolumeSpec] = volumeAttached.VolumeName
				volumeSpecMapByPlugin[pluginName] = volumeSpecMap
				continue
			}

			// If node doesn't support Bulk volume polling it is best to poll individually
			nodeError := oe.VerifyVolumesAreAttachedPerNode(nodeAttachedVolumes, node, actualStateOfWorld)
			if nodeError != nil {
				glog.Errorf("BulkVerifyVolumes.VerifyVolumesAreAttached verifying volumes on node %q with %v", node, nodeError)
			}
			break
		}
	}

	for pluginName, pluginNodeVolumes := range bulkVerifyPluginsByNode {
		bulkVerifyVolumeFunc, err := oe.operationGenerator.GenerateBulkVolumeVerifyFunc(
			pluginNodeVolumes,
			pluginName,
			volumeSpecMapByPlugin[pluginName],
			actualStateOfWorld)
		if err != nil {
			glog.Errorf("BulkVerifyVolumes.GenerateBulkVolumeVerifyFunc  error bulk verifying volumes for plugin %q with  %v", pluginName, err)
		}
		// Ugly hack to ensure - we don't do parallel bulk polling of same volume plugin
		uniquePluginName := v1.UniqueVolumeName(pluginName)
		err = oe.pendingOperations.Run(uniquePluginName, "" /* Pod Name */, bulkVerifyVolumeFunc)
		if err != nil {
			glog.Errorf("BulkVerifyVolumes.Run Error bulk volume verification for plugin %q  with %v", pluginName, err)
		}
	}
}

func (oe *operationExecutor) VerifyVolumesAreAttachedPerNode(
	attachedVolumes []AttachedVolume,
	nodeName types.NodeName,
	actualStateOfWorld ActualStateOfWorldAttacherUpdater) error {
	volumesAreAttachedFunc, err :=
		oe.operationGenerator.GenerateVolumesAreAttachedFunc(attachedVolumes, nodeName, actualStateOfWorld)
	if err != nil {
		return err
	}
	// Give an empty UniqueVolumeName so that this operation could be executed concurrently.
	return oe.pendingOperations.Run("" /* volumeName */, "" /* podName */, volumesAreAttachedFunc)
}

func (oe *operationExecutor) MountVolume(
	waitForAttachTimeout time.Duration,
	volumeToMount VolumeToMount,
	actualStateOfWorld ActualStateOfWorldMounterUpdater) error {
	mountFunc, err := oe.operationGenerator.GenerateMountVolumeFunc(
		waitForAttachTimeout, volumeToMount, actualStateOfWorld)
	if err != nil {
		return err
	}

	podName := nestedpendingoperations.EmptyUniquePodName
	// TODO: remove this -- not necessary
	if !volumeToMount.PluginIsAttachable {
		// Non-attachable volume plugins can execute mount for multiple pods
		// referencing the same volume in parallel
		podName = volumehelper.GetUniquePodName(volumeToMount.Pod)
	}

	return oe.pendingOperations.Run(
		volumeToMount.VolumeName, podName, mountFunc)
}

func (oe *operationExecutor) UnmountVolume(
	volumeToUnmount MountedVolume,
	actualStateOfWorld ActualStateOfWorldMounterUpdater) error {

	unmountFunc, err :=
		oe.operationGenerator.GenerateUnmountVolumeFunc(volumeToUnmount, actualStateOfWorld)
	if err != nil {
		return err
	}

	// All volume plugins can execute mount for multiple pods referencing the
	// same volume in parallel
	podName := volumetypes.UniquePodName(volumeToUnmount.PodUID)

	return oe.pendingOperations.Run(
		volumeToUnmount.VolumeName, podName, unmountFunc)
}

func (oe *operationExecutor) UnmountDevice(
	deviceToDetach AttachedVolume,
	actualStateOfWorld ActualStateOfWorldMounterUpdater,
	mounter mount.Interface) error {
	unmountDeviceFunc, err :=
		oe.operationGenerator.GenerateUnmountDeviceFunc(deviceToDetach, actualStateOfWorld, mounter)
	if err != nil {
		return err
	}

	return oe.pendingOperations.Run(
		deviceToDetach.VolumeName, "" /* podName */, unmountDeviceFunc)
}

func (oe *operationExecutor) VerifyControllerAttachedVolume(
	volumeToMount VolumeToMount,
	nodeName types.NodeName,
	actualStateOfWorld ActualStateOfWorldAttacherUpdater) error {
	verifyControllerAttachedVolumeFunc, err :=
		oe.operationGenerator.GenerateVerifyControllerAttachedVolumeFunc(volumeToMount, nodeName, actualStateOfWorld)
	if err != nil {
		return err
	}

	return oe.pendingOperations.Run(
		volumeToMount.VolumeName, "" /* podName */, verifyControllerAttachedVolumeFunc)
}

// TODO: this is a workaround for the unmount device issue caused by gci mounter.
// In GCI cluster, if gci mounter is used for mounting, the container started by mounter
// script will cause additional mounts created in the container. Since these mounts are
// irrelavant to the original mounts, they should be not considered when checking the
// mount references. Current solution is to filter out those mount paths that contain
// the string of original mount path.
// Plan to work on better approach to solve this issue.

func hasMountRefs(mountPath string, mountRefs []string) bool {
	count := 0
	for _, ref := range mountRefs {
		if !strings.Contains(ref, mountPath) {
			count = count + 1
		}
	}
	return count > 0
}
