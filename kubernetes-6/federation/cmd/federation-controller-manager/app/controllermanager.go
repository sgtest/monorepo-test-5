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

// Package app implements a server that runs a set of active
// components.  This includes cluster controller

package app

import (
	"net"
	"net/http"
	"net/http/pprof"
	goruntime "runtime"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/server/healthz"
	utilflag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/client-go/dynamic"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	federationclientset "github.com/sourcegraph/monorepo-test-1/kubernetes-6/federation/client/clientset_generated/federation_clientset"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-6/federation/cmd/federation-controller-manager/app/options"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-6/federation/pkg/dnsprovider"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-6/federation/pkg/federatedtypes"
	clustercontroller "github.com/sourcegraph/monorepo-test-1/kubernetes-6/federation/pkg/federation-controller/cluster"
	configmapcontroller "github.com/sourcegraph/monorepo-test-1/kubernetes-6/federation/pkg/federation-controller/configmap"
	daemonsetcontroller "github.com/sourcegraph/monorepo-test-1/kubernetes-6/federation/pkg/federation-controller/daemonset"
	deploymentcontroller "github.com/sourcegraph/monorepo-test-1/kubernetes-6/federation/pkg/federation-controller/deployment"
	ingresscontroller "github.com/sourcegraph/monorepo-test-1/kubernetes-6/federation/pkg/federation-controller/ingress"
	namespacecontroller "github.com/sourcegraph/monorepo-test-1/kubernetes-6/federation/pkg/federation-controller/namespace"
	replicasetcontroller "github.com/sourcegraph/monorepo-test-1/kubernetes-6/federation/pkg/federation-controller/replicaset"
	servicecontroller "github.com/sourcegraph/monorepo-test-1/kubernetes-6/federation/pkg/federation-controller/service"
	synccontroller "github.com/sourcegraph/monorepo-test-1/kubernetes-6/federation/pkg/federation-controller/sync"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-6/pkg/util/configz"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-6/pkg/version"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

// NewControllerManagerCommand creates a *cobra.Command object with default parameters
func NewControllerManagerCommand() *cobra.Command {
	s := options.NewCMServer()
	s.AddFlags(pflag.CommandLine)
	cmd := &cobra.Command{
		Use: "federation-controller-manager",
		Long: `The federation controller manager is a daemon that embeds
the core control loops shipped with federation. In applications of robotics and
automation, a control loop is a non-terminating loop that regulates the state of
the system. In federation, a controller is a control loop that watches the shared
state of the federation cluster through the apiserver and makes changes attempting
to move the current state towards the desired state. Examples of controllers that
ship with federation today is the cluster controller.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}

// Run runs the CMServer.  This should never exit.
func Run(s *options.CMServer) error {
	glog.Infof("%+v", version.Get())
	if c, err := configz.New("componentconfig"); err == nil {
		c.Set(s.ControllerManagerConfiguration)
	} else {
		glog.Errorf("unable to register configz: %s", err)
	}

	restClientCfg, err := clientcmd.BuildConfigFromFlags(s.Master, s.Kubeconfig)
	if err != nil || restClientCfg == nil {
		glog.V(2).Infof("Couldn't build the rest client config from flags: %v", err)
		return err
	}

	// Override restClientCfg qps/burst settings from flags
	restClientCfg.QPS = s.APIServerQPS
	restClientCfg.Burst = s.APIServerBurst

	go func() {
		mux := http.NewServeMux()
		healthz.InstallHandler(mux)
		if s.EnableProfiling {
			mux.HandleFunc("/debug/pprof/", pprof.Index)
			mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
			mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
			if s.EnableContentionProfiling {
				goruntime.SetBlockProfileRate(1)
			}
		}
		mux.Handle("/metrics", prometheus.Handler())

		server := &http.Server{
			Addr:    net.JoinHostPort(s.Address, strconv.Itoa(s.Port)),
			Handler: mux,
		}
		glog.Fatal(server.ListenAndServe())
	}()

	run := func() {
		err := StartControllers(s, restClientCfg)
		glog.Fatalf("error running controllers: %v", err)
		panic("unreachable")
	}
	run()
	panic("unreachable")
}

func StartControllers(s *options.CMServer, restClientCfg *restclient.Config) error {
	stopChan := wait.NeverStop
	minimizeLatency := false

	discoveryClient := discovery.NewDiscoveryClientForConfigOrDie(restClientCfg)
	serverResources, err := discoveryClient.ServerResources()
	if err != nil {
		glog.Fatalf("Could not find resources from API Server: %v", err)
	}

	clustercontroller.StartClusterController(restClientCfg, stopChan, s.ClusterMonitorPeriod.Duration)

	if controllerEnabled(s.Controllers, serverResources, servicecontroller.ControllerName, servicecontroller.RequiredResources, true) {
		dns, err := dnsprovider.InitDnsProvider(s.DnsProvider, s.DnsConfigFile)
		if err != nil {
			glog.Fatalf("Cloud provider could not be initialized: %v", err)
		}
		glog.Infof("Loading client config for service controller %q", servicecontroller.UserAgentName)
		scClientset := federationclientset.NewForConfigOrDie(restclient.AddUserAgent(restClientCfg, servicecontroller.UserAgentName))
		servicecontroller := servicecontroller.New(scClientset, dns, s.FederationName, s.ServiceDnsSuffix, s.ZoneName, s.ZoneID)
		glog.Infof("Running service controller")
		if err := servicecontroller.Run(s.ConcurrentServiceSyncs, wait.NeverStop); err != nil {
			glog.Fatalf("Failed to start service controller: %v", err)
		}
	}

	if controllerEnabled(s.Controllers, serverResources, namespacecontroller.ControllerName, namespacecontroller.RequiredResources, true) {
		glog.Infof("Loading client config for namespace controller %q", "namespace-controller")
		nsClientset := federationclientset.NewForConfigOrDie(restclient.AddUserAgent(restClientCfg, "namespace-controller"))
		namespaceController := namespacecontroller.NewNamespaceController(nsClientset, dynamic.NewDynamicClientPool(restclient.AddUserAgent(restClientCfg, "namespace-controller")))
		glog.Infof("Running namespace controller")
		namespaceController.Run(wait.NeverStop)
	}

	for kind, federatedType := range federatedtypes.FederatedTypes() {
		if controllerEnabled(s.Controllers, serverResources, federatedType.ControllerName, federatedType.RequiredResources, true) {
			synccontroller.StartFederationSyncController(kind, federatedType.AdapterFactory, restClientCfg, stopChan, minimizeLatency)
		}
	}

	if controllerEnabled(s.Controllers, serverResources, configmapcontroller.ControllerName, configmapcontroller.RequiredResources, true) {
		configmapcontrollerClientset := federationclientset.NewForConfigOrDie(restclient.AddUserAgent(restClientCfg, "configmap-controller"))
		configmapcontroller := configmapcontroller.NewConfigMapController(configmapcontrollerClientset)
		configmapcontroller.Run(wait.NeverStop)
	}

	if controllerEnabled(s.Controllers, serverResources, daemonsetcontroller.ControllerName, daemonsetcontroller.RequiredResources, true) {
		daemonsetcontrollerClientset := federationclientset.NewForConfigOrDie(restclient.AddUserAgent(restClientCfg, "daemonset-controller"))
		daemonsetcontroller := daemonsetcontroller.NewDaemonSetController(daemonsetcontrollerClientset)
		daemonsetcontroller.Run(wait.NeverStop)
	}

	if controllerEnabled(s.Controllers, serverResources, replicasetcontroller.ControllerName, replicasetcontroller.RequiredResources, true) {
		replicaSetClientset := federationclientset.NewForConfigOrDie(restclient.AddUserAgent(restClientCfg, replicasetcontroller.UserAgentName))
		replicaSetController := replicasetcontroller.NewReplicaSetController(replicaSetClientset)
		go replicaSetController.Run(s.ConcurrentReplicaSetSyncs, wait.NeverStop)
	}

	if controllerEnabled(s.Controllers, serverResources, deploymentcontroller.ControllerName, deploymentcontroller.RequiredResources, true) {
		deploymentClientset := federationclientset.NewForConfigOrDie(restclient.AddUserAgent(restClientCfg, deploymentcontroller.UserAgentName))
		deploymentController := deploymentcontroller.NewDeploymentController(deploymentClientset)
		// TODO: rename s.ConcurentReplicaSetSyncs
		go deploymentController.Run(s.ConcurrentReplicaSetSyncs, wait.NeverStop)
	}

	if controllerEnabled(s.Controllers, serverResources, ingresscontroller.ControllerName, ingresscontroller.RequiredResources, true) {
		glog.Infof("Loading client config for ingress controller %q", "ingress-controller")
		ingClientset := federationclientset.NewForConfigOrDie(restclient.AddUserAgent(restClientCfg, "ingress-controller"))
		ingressController := ingresscontroller.NewIngressController(ingClientset)
		glog.Infof("Running ingress controller")
		ingressController.Run(wait.NeverStop)
	}

	select {}
}

func controllerEnabled(controllers utilflag.ConfigurationMap, serverResources []*metav1.APIResourceList, controller string, requiredResources []schema.GroupVersionResource, defaultValue bool) bool {
	controllerConfig, ok := controllers[controller]
	if ok {
		if controllerConfig == "false" {
			glog.Infof("%s controller disabled by config", controller)
			return false
		}
		if controllerConfig == "true" {
			if !hasRequiredResources(serverResources, requiredResources) {
				glog.Fatalf("%s controller enabled explicitly but API Server does not have required resources", controller)
				panic("unreachable")
			}
			return true
		}
	} else if defaultValue {
		if !hasRequiredResources(serverResources, requiredResources) {
			glog.Warningf("%s controller disabled because API Server does not have required resources", controller)
			return false
		}
	}
	return defaultValue
}

func hasRequiredResources(serverResources []*metav1.APIResourceList, requiredResources []schema.GroupVersionResource) bool {
	for _, resource := range requiredResources {
		found := false
		for _, serverResource := range serverResources {
			if serverResource.GroupVersion == resource.GroupVersion().String() {
				for _, apiResource := range serverResource.APIResources {
					if apiResource.Name == resource.Resource {
						found = true
						break
					}
				}
			}
		}
		if !found {
			return false
		}
	}
	return true
}
