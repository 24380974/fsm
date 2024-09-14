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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/extension/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeListenerFilters implements ListenerFilterInterface
type FakeListenerFilters struct {
	Fake *FakeExtensionV1alpha1
	ns   string
}

var listenerfiltersResource = v1alpha1.SchemeGroupVersion.WithResource("listenerfilters")

var listenerfiltersKind = v1alpha1.SchemeGroupVersion.WithKind("ListenerFilter")

// Get takes name of the listenerFilter, and returns the corresponding listenerFilter object, and an error if there is any.
func (c *FakeListenerFilters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ListenerFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(listenerfiltersResource, c.ns, name), &v1alpha1.ListenerFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerFilter), err
}

// List takes label and field selectors, and returns the list of ListenerFilters that match those selectors.
func (c *FakeListenerFilters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ListenerFilterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(listenerfiltersResource, listenerfiltersKind, c.ns, opts), &v1alpha1.ListenerFilterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ListenerFilterList{ListMeta: obj.(*v1alpha1.ListenerFilterList).ListMeta}
	for _, item := range obj.(*v1alpha1.ListenerFilterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested listenerFilters.
func (c *FakeListenerFilters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(listenerfiltersResource, c.ns, opts))

}

// Create takes the representation of a listenerFilter and creates it.  Returns the server's representation of the listenerFilter, and an error, if there is any.
func (c *FakeListenerFilters) Create(ctx context.Context, listenerFilter *v1alpha1.ListenerFilter, opts v1.CreateOptions) (result *v1alpha1.ListenerFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(listenerfiltersResource, c.ns, listenerFilter), &v1alpha1.ListenerFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerFilter), err
}

// Update takes the representation of a listenerFilter and updates it. Returns the server's representation of the listenerFilter, and an error, if there is any.
func (c *FakeListenerFilters) Update(ctx context.Context, listenerFilter *v1alpha1.ListenerFilter, opts v1.UpdateOptions) (result *v1alpha1.ListenerFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(listenerfiltersResource, c.ns, listenerFilter), &v1alpha1.ListenerFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerFilter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeListenerFilters) UpdateStatus(ctx context.Context, listenerFilter *v1alpha1.ListenerFilter, opts v1.UpdateOptions) (*v1alpha1.ListenerFilter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(listenerfiltersResource, "status", c.ns, listenerFilter), &v1alpha1.ListenerFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerFilter), err
}

// Delete takes name of the listenerFilter and deletes it. Returns an error if one occurs.
func (c *FakeListenerFilters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(listenerfiltersResource, c.ns, name, opts), &v1alpha1.ListenerFilter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeListenerFilters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(listenerfiltersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ListenerFilterList{})
	return err
}

// Patch applies the patch and returns the patched listenerFilter.
func (c *FakeListenerFilters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ListenerFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(listenerfiltersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ListenerFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerFilter), err
}
