//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v2"

	internalv1alpha1 "k8s.io/api/apiserverinternal/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreaminternalv1alpha1informers "k8s.io/client-go/informers/apiserverinternal/v1alpha1"
	upstreaminternalv1alpha1listers "k8s.io/client-go/listers/apiserverinternal/v1alpha1"
	"k8s.io/client-go/tools/cache"

	clientset "github.com/kcp-dev/client-go/clients/clientset/versioned"
	"github.com/kcp-dev/client-go/clients/informers/internalinterfaces"
	internalv1alpha1listers "github.com/kcp-dev/client-go/clients/listers/apiserverinternal/v1alpha1"
)

// StorageVersionClusterInformer provides access to a shared informer and lister for
// StorageVersions.
type StorageVersionClusterInformer interface {
	Cluster(logicalcluster.Name) upstreaminternalv1alpha1informers.StorageVersionInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() internalv1alpha1listers.StorageVersionClusterLister
}

type storageVersionClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewStorageVersionClusterInformer constructs a new informer for StorageVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageVersionClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredStorageVersionClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredStorageVersionClusterInformer constructs a new informer for StorageVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageVersionClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InternalV1alpha1().StorageVersions().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InternalV1alpha1().StorageVersions().Watch(context.TODO(), options)
			},
		},
		&internalv1alpha1.StorageVersion{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageVersionClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredStorageVersionClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *storageVersionClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&internalv1alpha1.StorageVersion{}, f.defaultInformer)
}

func (f *storageVersionClusterInformer) Lister() internalv1alpha1listers.StorageVersionClusterLister {
	return internalv1alpha1listers.NewStorageVersionClusterLister(f.Informer().GetIndexer())
}

func (f *storageVersionClusterInformer) Cluster(cluster logicalcluster.Name) upstreaminternalv1alpha1informers.StorageVersionInformer {
	return &storageVersionInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type storageVersionInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreaminternalv1alpha1listers.StorageVersionLister
}

func (f *storageVersionInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *storageVersionInformer) Lister() upstreaminternalv1alpha1listers.StorageVersionLister {
	return f.lister
}
