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
	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/machine/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// VirtualMachineLister helps list VirtualMachines.
// All objects returned here must be treated as read-only.
type VirtualMachineLister interface {
	// List lists all VirtualMachines in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualMachine, err error)
	// VirtualMachines returns an object that can list and get VirtualMachines.
	VirtualMachines(namespace string) VirtualMachineNamespaceLister
	VirtualMachineListerExpansion
}

// virtualMachineLister implements the VirtualMachineLister interface.
type virtualMachineLister struct {
	listers.ResourceIndexer[*v1alpha1.VirtualMachine]
}

// NewVirtualMachineLister returns a new VirtualMachineLister.
func NewVirtualMachineLister(indexer cache.Indexer) VirtualMachineLister {
	return &virtualMachineLister{listers.New[*v1alpha1.VirtualMachine](indexer, v1alpha1.Resource("virtualmachine"))}
}

// VirtualMachines returns an object that can list and get VirtualMachines.
func (s *virtualMachineLister) VirtualMachines(namespace string) VirtualMachineNamespaceLister {
	return virtualMachineNamespaceLister{listers.NewNamespaced[*v1alpha1.VirtualMachine](s.ResourceIndexer, namespace)}
}

// VirtualMachineNamespaceLister helps list and get VirtualMachines.
// All objects returned here must be treated as read-only.
type VirtualMachineNamespaceLister interface {
	// List lists all VirtualMachines in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualMachine, err error)
	// Get retrieves the VirtualMachine from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VirtualMachine, error)
	VirtualMachineNamespaceListerExpansion
}

// virtualMachineNamespaceLister implements the VirtualMachineNamespaceLister
// interface.
type virtualMachineNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.VirtualMachine]
}