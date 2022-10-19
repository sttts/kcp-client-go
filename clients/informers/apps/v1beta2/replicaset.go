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

package v1beta2

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v2"

	appsv1beta2 "k8s.io/api/apps/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamappsv1beta2informers "k8s.io/client-go/informers/apps/v1beta2"
	upstreamappsv1beta2listers "k8s.io/client-go/listers/apps/v1beta2"
	"k8s.io/client-go/tools/cache"

	clientset "github.com/kcp-dev/client-go/clients/clientset/versioned"
	"github.com/kcp-dev/client-go/clients/informers/internalinterfaces"
	appsv1beta2listers "github.com/kcp-dev/client-go/clients/listers/apps/v1beta2"
)

// ReplicaSetClusterInformer provides access to a shared informer and lister for
// ReplicaSets.
type ReplicaSetClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamappsv1beta2informers.ReplicaSetInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() appsv1beta2listers.ReplicaSetClusterLister
}

type replicaSetClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewReplicaSetClusterInformer constructs a new informer for ReplicaSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReplicaSetClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredReplicaSetClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredReplicaSetClusterInformer constructs a new informer for ReplicaSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReplicaSetClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1beta2().ReplicaSets().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1beta2().ReplicaSets().Watch(context.TODO(), options)
			},
		},
		&appsv1beta2.ReplicaSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *replicaSetClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredReplicaSetClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *replicaSetClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&appsv1beta2.ReplicaSet{}, f.defaultInformer)
}

func (f *replicaSetClusterInformer) Lister() appsv1beta2listers.ReplicaSetClusterLister {
	return appsv1beta2listers.NewReplicaSetClusterLister(f.Informer().GetIndexer())
}

func (f *replicaSetClusterInformer) Cluster(cluster logicalcluster.Name) upstreamappsv1beta2informers.ReplicaSetInformer {
	return &replicaSetInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type replicaSetInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamappsv1beta2listers.ReplicaSetLister
}

func (f *replicaSetInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *replicaSetInformer) Lister() upstreamappsv1beta2listers.ReplicaSetLister {
	return f.lister
}
