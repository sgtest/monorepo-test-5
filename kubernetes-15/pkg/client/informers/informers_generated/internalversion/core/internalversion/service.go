/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	api "github.com/sourcegraph/monorepo-test-1/kubernetes-15/pkg/api"
	internalclientset "github.com/sourcegraph/monorepo-test-1/kubernetes-15/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/sourcegraph/monorepo-test-1/kubernetes-15/pkg/client/informers/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/sourcegraph/monorepo-test-1/kubernetes-15/pkg/client/listers/core/internalversion"
	time "time"
)

// ServiceInformer provides access to a shared informer and lister for
// Services.
type ServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ServiceLister
}

type serviceInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newServiceInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Core().Services(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Core().Services(v1.NamespaceAll).Watch(options)
			},
		},
		&api.Service{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *serviceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&api.Service{}, newServiceInformer)
}

func (f *serviceInformer) Lister() internalversion.ServiceLister {
	return internalversion.NewServiceLister(f.Informer().GetIndexer())
}
