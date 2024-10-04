/*
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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/policy/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// EgressLister helps list Egresses.
// All objects returned here must be treated as read-only.
type EgressLister interface {
	// List lists all Egresses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Egress, err error)
	// Egresses returns an object that can list and get Egresses.
	Egresses(namespace string) EgressNamespaceLister
	EgressListerExpansion
}

// egressLister implements the EgressLister interface.
type egressLister struct {
	listers.ResourceIndexer[*v1alpha1.Egress]
}

// NewEgressLister returns a new EgressLister.
func NewEgressLister(indexer cache.Indexer) EgressLister {
	return &egressLister{listers.New[*v1alpha1.Egress](indexer, v1alpha1.Resource("egress"))}
}

// Egresses returns an object that can list and get Egresses.
func (s *egressLister) Egresses(namespace string) EgressNamespaceLister {
	return egressNamespaceLister{listers.NewNamespaced[*v1alpha1.Egress](s.ResourceIndexer, namespace)}
}

// EgressNamespaceLister helps list and get Egresses.
// All objects returned here must be treated as read-only.
type EgressNamespaceLister interface {
	// List lists all Egresses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Egress, err error)
	// Get retrieves the Egress from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Egress, error)
	EgressNamespaceListerExpansion
}

// egressNamespaceLister implements the EgressNamespaceLister
// interface.
type egressNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.Egress]
}