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

	flowcontrolv1beta2 "k8s.io/api/flowcontrol/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	upstreamflowcontrolv1beta2informers "k8s.io/client-go/informers/flowcontrol/v1beta2"
	upstreamflowcontrolv1beta2listers "k8s.io/client-go/listers/flowcontrol/v1beta2"
	"k8s.io/client-go/tools/cache"

	clientset "github.com/kcp-dev/client-go/clients/clientset/versioned"
	"github.com/kcp-dev/client-go/clients/informers/internalinterfaces"
	flowcontrolv1beta2listers "github.com/kcp-dev/client-go/clients/listers/flowcontrol/v1beta2"
)

// FlowSchemaClusterInformer provides access to a shared informer and lister for
// FlowSchemas.
type FlowSchemaClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamflowcontrolv1beta2informers.FlowSchemaInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() flowcontrolv1beta2listers.FlowSchemaClusterLister
}

type flowSchemaClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewFlowSchemaClusterInformer constructs a new informer for FlowSchema type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFlowSchemaClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredFlowSchemaClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredFlowSchemaClusterInformer constructs a new informer for FlowSchema type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFlowSchemaClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1beta2().FlowSchemas().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowcontrolV1beta2().FlowSchemas().Watch(context.TODO(), options)
			},
		},
		&flowcontrolv1beta2.FlowSchema{},
		resyncPeriod,
		indexers,
	)
}

func (f *flowSchemaClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredFlowSchemaClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
	},
		f.tweakListOptions,
	)
}

func (f *flowSchemaClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&flowcontrolv1beta2.FlowSchema{}, f.defaultInformer)
}

func (f *flowSchemaClusterInformer) Lister() flowcontrolv1beta2listers.FlowSchemaClusterLister {
	return flowcontrolv1beta2listers.NewFlowSchemaClusterLister(f.Informer().GetIndexer())
}

func (f *flowSchemaClusterInformer) Cluster(cluster logicalcluster.Name) upstreamflowcontrolv1beta2informers.FlowSchemaInformer {
	return &flowSchemaInformer{
		informer: f.Informer().Cluster(cluster),
		lister:   f.Lister().Cluster(cluster),
	}
}

type flowSchemaInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamflowcontrolv1beta2listers.FlowSchemaLister
}

func (f *flowSchemaInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *flowSchemaInformer) Lister() upstreamflowcontrolv1beta2listers.FlowSchemaLister {
	return f.lister
}
