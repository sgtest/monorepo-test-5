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

package winuserspace

import (
	"net"

	"github.com/sourcegraph/monorepo-test-1/kubernetes-15/pkg/api"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-15/pkg/proxy"
)

// LoadBalancer is an interface for distributing incoming requests to service endpoints.
type LoadBalancer interface {
	// NextEndpoint returns the endpoint to handle a request for the given
	// service-port and source address.
	NextEndpoint(service proxy.ServicePortName, srcAddr net.Addr, sessionAffinityReset bool) (string, error)
	NewService(service proxy.ServicePortName, sessionAffinityType api.ServiceAffinity, stickyMaxAgeMinutes int) error
	DeleteService(service proxy.ServicePortName)
	CleanupStaleStickySessions(service proxy.ServicePortName)
}
