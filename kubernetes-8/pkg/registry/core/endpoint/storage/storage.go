/*
Copyright 2015 The Kubernetes Authors.

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

package storage

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-8/pkg/api"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-8/pkg/registry/cachesize"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-8/pkg/registry/core/endpoint"
)

type REST struct {
	*genericregistry.Store
}

// NewREST returns a RESTStorage object that will work against endpoints.
func NewREST(optsGetter generic.RESTOptionsGetter) *REST {
	store := &genericregistry.Store{
		Copier:      api.Scheme,
		NewFunc:     func() runtime.Object { return &api.Endpoints{} },
		NewListFunc: func() runtime.Object { return &api.EndpointsList{} },
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*api.Endpoints).Name, nil
		},
		PredicateFunc:     endpoint.MatchEndpoints,
		QualifiedResource: api.Resource("endpoints"),
		WatchCacheSize:    cachesize.GetWatchCacheSizeByResource("endpoints"),

		CreateStrategy: endpoint.Strategy,
		UpdateStrategy: endpoint.Strategy,
		DeleteStrategy: endpoint.Strategy,
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: endpoint.GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err) // TODO: Propagate error up
	}
	return &REST{store}
}

// Implement ShortNamesProvider
var _ rest.ShortNamesProvider = &REST{}

// ShortNames implements the ShortNamesProvider interface. Returns a list of short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"ep"}
}