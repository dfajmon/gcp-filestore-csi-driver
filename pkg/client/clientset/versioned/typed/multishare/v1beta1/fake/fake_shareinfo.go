/*
Copyright The Kubernetes Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "sigs.k8s.io/gcp-filestore-csi-driver/pkg/apis/multishare/v1beta1"
)

// FakeShareInfos implements ShareInfoInterface
type FakeShareInfos struct {
	Fake *FakeMultishareV1beta1
	ns   string
}

var shareinfosResource = schema.GroupVersionResource{Group: "multishare.filestore.csi.storage.gke.io", Version: "v1beta1", Resource: "shareinfos"}

var shareinfosKind = schema.GroupVersionKind{Group: "multishare.filestore.csi.storage.gke.io", Version: "v1beta1", Kind: "ShareInfo"}

// Get takes name of the shareInfo, and returns the corresponding shareInfo object, and an error if there is any.
func (c *FakeShareInfos) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ShareInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(shareinfosResource, c.ns, name), &v1beta1.ShareInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ShareInfo), err
}

// List takes label and field selectors, and returns the list of ShareInfos that match those selectors.
func (c *FakeShareInfos) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ShareInfoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(shareinfosResource, shareinfosKind, c.ns, opts), &v1beta1.ShareInfoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ShareInfoList{ListMeta: obj.(*v1beta1.ShareInfoList).ListMeta}
	for _, item := range obj.(*v1beta1.ShareInfoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested shareInfos.
func (c *FakeShareInfos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(shareinfosResource, c.ns, opts))

}

// Create takes the representation of a shareInfo and creates it.  Returns the server's representation of the shareInfo, and an error, if there is any.
func (c *FakeShareInfos) Create(ctx context.Context, shareInfo *v1beta1.ShareInfo, opts v1.CreateOptions) (result *v1beta1.ShareInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(shareinfosResource, c.ns, shareInfo), &v1beta1.ShareInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ShareInfo), err
}

// Update takes the representation of a shareInfo and updates it. Returns the server's representation of the shareInfo, and an error, if there is any.
func (c *FakeShareInfos) Update(ctx context.Context, shareInfo *v1beta1.ShareInfo, opts v1.UpdateOptions) (result *v1beta1.ShareInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(shareinfosResource, c.ns, shareInfo), &v1beta1.ShareInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ShareInfo), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeShareInfos) UpdateStatus(ctx context.Context, shareInfo *v1beta1.ShareInfo, opts v1.UpdateOptions) (*v1beta1.ShareInfo, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(shareinfosResource, "status", c.ns, shareInfo), &v1beta1.ShareInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ShareInfo), err
}

// Delete takes name of the shareInfo and deletes it. Returns an error if one occurs.
func (c *FakeShareInfos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(shareinfosResource, c.ns, name, opts), &v1beta1.ShareInfo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeShareInfos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(shareinfosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ShareInfoList{})
	return err
}

// Patch applies the patch and returns the patched shareInfo.
func (c *FakeShareInfos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ShareInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(shareinfosResource, c.ns, name, pt, data, subresources...), &v1beta1.ShareInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ShareInfo), err
}
