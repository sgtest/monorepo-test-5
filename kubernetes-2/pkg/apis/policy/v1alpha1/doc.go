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

// +k8s:defaulter-gen=TypeMeta

// Package policy is for any kind of policy object.  Suitable examples, even if
// they aren't all here, are PodDisruptionBudget, PodSecurityPolicy,
// NetworkPolicy, etc.
package v1alpha1 // import "github.com/sourcegraph/monorepo-test-1/kubernetes-2/pkg/apis/policy/v1alpha1"
