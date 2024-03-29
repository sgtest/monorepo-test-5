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

package storage

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-10/pkg/api"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-10/pkg/apis/rbac"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-10/pkg/registry/cachesize"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-10/pkg/registry/rbac/role"
)

// REST implements a RESTStorage for Role
type REST struct {
	*genericregistry.Store
}

// NewREST returns a RESTStorage object that will work against Role objects.
func NewREST(optsGetter generic.RESTOptionsGetter) *REST {
	store := &genericregistry.Store{
		Copier:      api.Scheme,
		NewFunc:     func() runtime.Object { return &rbac.Role{} },
		NewListFunc: func() runtime.Object { return &rbac.RoleList{} },
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*rbac.Role).Name, nil
		},
		PredicateFunc:     role.Matcher,
		QualifiedResource: rbac.Resource("roles"),
		WatchCacheSize:    cachesize.GetWatchCacheSizeByResource("roles"),

		CreateStrategy: role.Strategy,
		UpdateStrategy: role.Strategy,
		DeleteStrategy: role.Strategy,
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: role.GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err) // TODO: Propagate error up
	}

	return &REST{store}
}
