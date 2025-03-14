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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"
	time "time"

	apistoragev1beta1 "k8s.io/api/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	storagev1beta1 "k8s.io/client-go/listers/storage/v1beta1"
	cache "k8s.io/client-go/tools/cache"
)

// VolumeAttributesClassInformer provides access to a shared informer and lister for
// VolumeAttributesClasses.
type VolumeAttributesClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() storagev1beta1.VolumeAttributesClassLister
}

type volumeAttributesClassInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewVolumeAttributesClassInformer constructs a new informer for VolumeAttributesClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVolumeAttributesClassInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVolumeAttributesClassInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredVolumeAttributesClassInformer constructs a new informer for VolumeAttributesClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVolumeAttributesClassInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1beta1().VolumeAttributesClasses().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1beta1().VolumeAttributesClasses().Watch(context.TODO(), options)
			},
		},
		&apistoragev1beta1.VolumeAttributesClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *volumeAttributesClassInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVolumeAttributesClassInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *volumeAttributesClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apistoragev1beta1.VolumeAttributesClass{}, f.defaultInformer)
}

func (f *volumeAttributesClassInformer) Lister() storagev1beta1.VolumeAttributesClassLister {
	return storagev1beta1.NewVolumeAttributesClassLister(f.Informer().GetIndexer())
}
