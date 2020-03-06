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

// FakeSriovPFs implements SriovPFInterface
type FakeSriovPFs struct {
	Fake *FakeSriovV1
}

var sriovpfsResource = schema.GroupVersionResource{Group: "k8s.cni.cncf.io", Version: "v1", Resource: "sriovpfs"}

var sriovpfsKind = schema.GroupVersionKind{Group: "k8s.cni.cncf.io", Version: "v1", Kind: "SriovPF"}

// Get takes name of the sriovPF, and returns the corresponding sriovPF object, and an error if there is any.
func (c *FakeSriovPFs) Get(name string, options v1.GetOptions) (result *sriovv1.SriovPF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(sriovpfsResource, name), &sriovv1.SriovPF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*sriovv1.SriovPF), err
}

// List takes label and field selectors, and returns the list of SriovPFs that match those selectors.
func (c *FakeSriovPFs) List(opts v1.ListOptions) (result *sriovv1.SriovPFList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(sriovpfsResource, sriovpfsKind, opts), &sriovv1.SriovPFList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &sriovv1.SriovPFList{ListMeta: obj.(*sriovv1.SriovPFList).ListMeta}
	for _, item := range obj.(*sriovv1.SriovPFList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sriovPFs.
func (c *FakeSriovPFs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(sriovpfsResource, opts))
}

// Create takes the representation of a sriovPF and creates it.  Returns the server's representation of the sriovPF, and an error, if there is any.
func (c *FakeSriovPFs) Create(sriovPF *sriovv1.SriovPF) (result *sriovv1.SriovPF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(sriovpfsResource, sriovPF), &sriovv1.SriovPF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*sriovv1.SriovPF), err
}

// Update takes the representation of a sriovPF and updates it. Returns the server's representation of the sriovPF, and an error, if there is any.
func (c *FakeSriovPFs) Update(sriovPF *sriovv1.SriovPF) (result *sriovv1.SriovPF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(sriovpfsResource, sriovPF), &sriovv1.SriovPF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*sriovv1.SriovPF), err
}

// Delete takes name of the sriovPF and deletes it. Returns an error if one occurs.
func (c *FakeSriovPFs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(sriovpfsResource, name), &sriovv1.SriovPF{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSriovPFs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(sriovpfsResource, listOptions)

	_, err := c.Fake.Invokes(action, &sriovv1.SriovPFList{})
	return err
}

// Patch applies the patch and returns the patched sriovPF.
func (c *FakeSriovPFs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *sriovv1.SriovPF, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(sriovpfsResource, name, pt, data, subresources...), &sriovv1.SriovPF{})
	if obj == nil {
		return nil, err
	}
	return obj.(*sriovv1.SriovPF), err
}