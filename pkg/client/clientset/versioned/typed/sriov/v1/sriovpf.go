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

package v1

import (
	"time"

	v1 "github.com/fafucoder/sriov-crd/pkg/apis/sriov/v1"
	scheme "github.com/fafucoder/sriov-crd/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SriovPFsGetter has a method to return a SriovPFInterface.
// A group's client should implement this interface.
type SriovPFsGetter interface {
	SriovPFs() SriovPFInterface
}

// SriovPFInterface has methods to work with SriovPF resources.
type SriovPFInterface interface {
	Create(*v1.SriovPF) (*v1.SriovPF, error)
	Update(*v1.SriovPF) (*v1.SriovPF, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.SriovPF, error)
	List(opts metav1.ListOptions) (*v1.SriovPFList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SriovPF, err error)
	SriovPFExpansion
}

// sriovPFs implements SriovPFInterface
type sriovPFs struct {
	client rest.Interface
}

// newSriovPFs returns a SriovPFs
func newSriovPFs(c *SriovV1Client) *sriovPFs {
	return &sriovPFs{
		client: c.RESTClient(),
	}
}

// Get takes name of the sriovPF, and returns the corresponding sriovPF object, and an error if there is any.
func (c *sriovPFs) Get(name string, options metav1.GetOptions) (result *v1.SriovPF, err error) {
	result = &v1.SriovPF{}
	err = c.client.Get().
		Resource("sriovpfs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SriovPFs that match those selectors.
func (c *sriovPFs) List(opts metav1.ListOptions) (result *v1.SriovPFList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SriovPFList{}
	err = c.client.Get().
		Resource("sriovpfs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sriovPFs.
func (c *sriovPFs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("sriovpfs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sriovPF and creates it.  Returns the server's representation of the sriovPF, and an error, if there is any.
func (c *sriovPFs) Create(sriovPF *v1.SriovPF) (result *v1.SriovPF, err error) {
	result = &v1.SriovPF{}
	err = c.client.Post().
		Resource("sriovpfs").
		Body(sriovPF).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sriovPF and updates it. Returns the server's representation of the sriovPF, and an error, if there is any.
func (c *sriovPFs) Update(sriovPF *v1.SriovPF) (result *v1.SriovPF, err error) {
	result = &v1.SriovPF{}
	err = c.client.Put().
		Resource("sriovpfs").
		Name(sriovPF.Name).
		Body(sriovPF).
		Do().
		Into(result)
	return
}

// Delete takes name of the sriovPF and deletes it. Returns an error if one occurs.
func (c *sriovPFs) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("sriovpfs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sriovPFs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("sriovpfs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sriovPF.
func (c *sriovPFs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SriovPF, err error) {
	result = &v1.SriovPF{}
	err = c.client.Patch(pt).
		Resource("sriovpfs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}