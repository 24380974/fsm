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

// FakeZipkins implements ZipkinInterface
type FakeZipkins struct {
	Fake *FakeExtensionV1alpha1
	ns   string
}

var zipkinsResource = v1alpha1.SchemeGroupVersion.WithResource("zipkins")

var zipkinsKind = v1alpha1.SchemeGroupVersion.WithKind("Zipkin")

// Get takes name of the zipkin, and returns the corresponding zipkin object, and an error if there is any.
func (c *FakeZipkins) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Zipkin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(zipkinsResource, c.ns, name), &v1alpha1.Zipkin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Zipkin), err
}

// List takes label and field selectors, and returns the list of Zipkins that match those selectors.
func (c *FakeZipkins) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ZipkinList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(zipkinsResource, zipkinsKind, c.ns, opts), &v1alpha1.ZipkinList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ZipkinList{ListMeta: obj.(*v1alpha1.ZipkinList).ListMeta}
	for _, item := range obj.(*v1alpha1.ZipkinList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested zipkins.
func (c *FakeZipkins) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(zipkinsResource, c.ns, opts))

}

// Create takes the representation of a zipkin and creates it.  Returns the server's representation of the zipkin, and an error, if there is any.
func (c *FakeZipkins) Create(ctx context.Context, zipkin *v1alpha1.Zipkin, opts v1.CreateOptions) (result *v1alpha1.Zipkin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(zipkinsResource, c.ns, zipkin), &v1alpha1.Zipkin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Zipkin), err
}

// Update takes the representation of a zipkin and updates it. Returns the server's representation of the zipkin, and an error, if there is any.
func (c *FakeZipkins) Update(ctx context.Context, zipkin *v1alpha1.Zipkin, opts v1.UpdateOptions) (result *v1alpha1.Zipkin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(zipkinsResource, c.ns, zipkin), &v1alpha1.Zipkin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Zipkin), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeZipkins) UpdateStatus(ctx context.Context, zipkin *v1alpha1.Zipkin, opts v1.UpdateOptions) (*v1alpha1.Zipkin, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(zipkinsResource, "status", c.ns, zipkin), &v1alpha1.Zipkin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Zipkin), err
}

// Delete takes name of the zipkin and deletes it. Returns an error if one occurs.
func (c *FakeZipkins) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(zipkinsResource, c.ns, name, opts), &v1alpha1.Zipkin{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeZipkins) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(zipkinsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ZipkinList{})
	return err
}

// Patch applies the patch and returns the patched zipkin.
func (c *FakeZipkins) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Zipkin, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(zipkinsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Zipkin{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Zipkin), err
}
