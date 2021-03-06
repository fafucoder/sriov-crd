/*
Copyright 2019 The Lijiao Authors.

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
	sriovv1 "github.com/fafucoder/sriov-crd/pkg/apis/sriov/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVFs implements VFInterface
type FakeVFs struct {
	Fake *FakeK8sCniCncfIoV1
}

var vfsResource = schema.GroupVersionResource{Group: "k8s.cni.cncf.io", Version: "v1", Resource: "vfs"}

var vfsKind = schema.GroupVersionKind{Group: "k8s.cni.cncf.io", Version: "v1", Kind: "VF"}

// Get takes name of the vF, and returns the corresponding vF object, and an error if there is any.
func (c *FakeVFs) Get(name string, options v1.GetOptions) (result *sriovv1.VF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(vfsResource, name), &sriovv1.VF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*sriovv1.VF), err
}

// List takes label and field selectors, and returns the list of VFs that match those selectors.
func (c *FakeVFs) List(opts v1.ListOptions) (result *sriovv1.VFList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(vfsResource, vfsKind, opts), &sriovv1.VFList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &sriovv1.VFList{ListMeta: obj.(*sriovv1.VFList).ListMeta}
	for _, item := range obj.(*sriovv1.VFList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vFs.
func (c *FakeVFs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(vfsResource, opts))
}

// Create takes the representation of a vF and creates it.  Returns the server's representation of the vF, and an error, if there is any.
func (c *FakeVFs) Create(vF *sriovv1.VF) (result *sriovv1.VF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(vfsResource, vF), &sriovv1.VF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*sriovv1.VF), err
}

// Update takes the representation of a vF and updates it. Returns the server's representation of the vF, and an error, if there is any.
func (c *FakeVFs) Update(vF *sriovv1.VF) (result *sriovv1.VF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(vfsResource, vF), &sriovv1.VF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*sriovv1.VF), err
}

// Delete takes name of the vF and deletes it. Returns an error if one occurs.
func (c *FakeVFs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(vfsResource, name), &sriovv1.VF{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVFs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(vfsResource, listOptions)

	_, err := c.Fake.Invokes(action, &sriovv1.VFList{})
	return err
}

// Patch applies the patch and returns the patched vF.
func (c *FakeVFs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *sriovv1.VF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(vfsResource, name, pt, data, subresources...), &sriovv1.VF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*sriovv1.VF), err
}
