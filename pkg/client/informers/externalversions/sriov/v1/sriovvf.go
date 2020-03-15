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
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	sriovv1 "github.com/fafucoder/sriov-crd/pkg/apis/sriov/v1"
	versioned "github.com/fafucoder/sriov-crd/pkg/client/clientset/versioned"
	internalinterfaces "github.com/fafucoder/sriov-crd/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/fafucoder/sriov-crd/pkg/client/listers/sriov/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SriovVFInformer provides access to a shared informer and lister for
// SriovVFs.
type SriovVFInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SriovVFLister
}

type sriovVFInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewSriovVFInformer constructs a new informer for SriovVF type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSriovVFInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSriovVFInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredSriovVFInformer constructs a new informer for SriovVF type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSriovVFInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().SriovVFs().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().SriovVFs().Watch(options)
			},
		},
		&sriovv1.SriovVF{},
		resyncPeriod,
		indexers,
	)
}

func (f *sriovVFInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSriovVFInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sriovVFInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sriovv1.SriovVF{}, f.defaultInformer)
}

func (f *sriovVFInformer) Lister() v1.SriovVFLister {
	return v1.NewSriovVFLister(f.Informer().GetIndexer())
}
