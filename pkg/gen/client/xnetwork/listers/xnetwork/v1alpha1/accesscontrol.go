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
	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/xnetwork/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// AccessControlLister helps list AccessControls.
// All objects returned here must be treated as read-only.
type AccessControlLister interface {
	// List lists all AccessControls in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AccessControl, err error)
	// AccessControls returns an object that can list and get AccessControls.
	AccessControls(namespace string) AccessControlNamespaceLister
	AccessControlListerExpansion
}

// accessControlLister implements the AccessControlLister interface.
type accessControlLister struct {
	listers.ResourceIndexer[*v1alpha1.AccessControl]
}

// NewAccessControlLister returns a new AccessControlLister.
func NewAccessControlLister(indexer cache.Indexer) AccessControlLister {
	return &accessControlLister{listers.New[*v1alpha1.AccessControl](indexer, v1alpha1.Resource("accesscontrol"))}
}

// AccessControls returns an object that can list and get AccessControls.
func (s *accessControlLister) AccessControls(namespace string) AccessControlNamespaceLister {
	return accessControlNamespaceLister{listers.NewNamespaced[*v1alpha1.AccessControl](s.ResourceIndexer, namespace)}
}

// AccessControlNamespaceLister helps list and get AccessControls.
// All objects returned here must be treated as read-only.
type AccessControlNamespaceLister interface {
	// List lists all AccessControls in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AccessControl, err error)
	// Get retrieves the AccessControl from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AccessControl, error)
	AccessControlNamespaceListerExpansion
}

// accessControlNamespaceLister implements the AccessControlNamespaceLister
// interface.
type accessControlNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.AccessControl]
}
