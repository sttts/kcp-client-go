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

package v1beta1

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v2"

	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamnetworkingv1beta1informers "k8s.io/client-go/informers/networking/v1beta1"
	upstreamnetworkingv1beta1listers "k8s.io/client-go/listers/networking/v1beta1"
	"k8s.io/client-go/tools/cache"

	clientset "github.com/kcp-dev/client-go/clients/clientset/versioned"
	"github.com/kcp-dev/client-go/clients/informers/internalinterfaces"
	networkingv1beta1listers "github.com/kcp-dev/client-go/clients/listers/networking/v1beta1"
)

// IngressClassClusterInformer provides access to a shared informer and lister for
// IngressClasses.
type IngressClassClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamnetworkingv1beta1informers.IngressClassInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() networkingv1beta1listers.IngressClassClusterLister
}

type ingressClassClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewIngressClassClusterInformer constructs a new informer for IngressClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIngressClassClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredIngressClassClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredIngressClassClusterInformer constructs a new informer for IngressClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIngressClassClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1beta1().IngressClasses().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1beta1().IngressClasses().Watch(context.TODO(), options)
			},
		},
		&networkingv1beta1.IngressClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *ingressClassClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredIngressClassClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *ingressClassClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&networkingv1beta1.IngressClass{}, f.defaultInformer)
}

func (f *ingressClassClusterInformer) Lister() networkingv1beta1listers.IngressClassClusterLister {
	return networkingv1beta1listers.NewIngressClassClusterLister(f.Informer().GetIndexer())
}

func (f *ingressClassClusterInformer) Cluster(cluster logicalcluster.Name) upstreamnetworkingv1beta1informers.IngressClassInformer {
	return &ingressClassInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type ingressClassInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamnetworkingv1beta1listers.IngressClassLister
}

func (f *ingressClassInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *ingressClassInformer) Lister() upstreamnetworkingv1beta1listers.IngressClassLister {
	return f.lister
}
