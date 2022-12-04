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
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	flowcontrolv1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	flowcontrolv1alpha1listers "k8s.io/client-go/listers/flowcontrol/v1alpha1"
	"k8s.io/client-go/tools/cache"
)

// FlowSchemaClusterLister can list FlowSchemas across all workspaces, or scope down to a FlowSchemaLister for one workspace.
// All objects returned here must be treated as read-only.
type FlowSchemaClusterLister interface {
	// List lists all FlowSchemas in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*flowcontrolv1alpha1.FlowSchema, err error)
	// Cluster returns a lister that can list and get FlowSchemas in one workspace.
	Cluster(cluster logicalcluster.Name) flowcontrolv1alpha1listers.FlowSchemaLister
	FlowSchemaClusterListerExpansion
}

type flowSchemaClusterLister struct {
	indexer cache.Indexer
}

// NewFlowSchemaClusterLister returns a new FlowSchemaClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewFlowSchemaClusterLister(indexer cache.Indexer) *flowSchemaClusterLister {
	return &flowSchemaClusterLister{indexer: indexer}
}

// List lists all FlowSchemas in the indexer across all workspaces.
func (s *flowSchemaClusterLister) List(selector labels.Selector) (ret []*flowcontrolv1alpha1.FlowSchema, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*flowcontrolv1alpha1.FlowSchema))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get FlowSchemas.
func (s *flowSchemaClusterLister) Cluster(cluster logicalcluster.Name) flowcontrolv1alpha1listers.FlowSchemaLister {
	return &flowSchemaLister{indexer: s.indexer, cluster: cluster}
}

// flowSchemaLister implements the flowcontrolv1alpha1listers.FlowSchemaLister interface.
type flowSchemaLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all FlowSchemas in the indexer for a workspace.
func (s *flowSchemaLister) List(selector labels.Selector) (ret []*flowcontrolv1alpha1.FlowSchema, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*flowcontrolv1alpha1.FlowSchema))
	})
	return ret, err
}

// Get retrieves the FlowSchema from the indexer for a given workspace and name.
func (s *flowSchemaLister) Get(name string) (*flowcontrolv1alpha1.FlowSchema, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(flowcontrolv1alpha1.Resource("FlowSchema"), name)
	}
	return obj.(*flowcontrolv1alpha1.FlowSchema), nil
}
