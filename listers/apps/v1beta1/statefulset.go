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
	kcpcache "github.com/kcp-dev/apimachinery/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	appsv1beta1 "k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	appsv1beta1listers "k8s.io/client-go/listers/apps/v1beta1"
	"k8s.io/client-go/tools/cache"
)

// StatefulSetClusterLister can list StatefulSets across all workspaces, or scope down to a StatefulSetLister for one workspace.
// All objects returned here must be treated as read-only.
type StatefulSetClusterLister interface {
	// List lists all StatefulSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error)
	// Cluster returns a lister that can list and get StatefulSets in one workspace.
	Cluster(cluster logicalcluster.Name) appsv1beta1listers.StatefulSetLister
	StatefulSetClusterListerExpansion
}

type statefulSetClusterLister struct {
	indexer cache.Indexer
}

// NewStatefulSetClusterLister returns a new StatefulSetClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewStatefulSetClusterLister(indexer cache.Indexer) *statefulSetClusterLister {
	return &statefulSetClusterLister{indexer: indexer}
}

// List lists all StatefulSets in the indexer across all workspaces.
func (s *statefulSetClusterLister) List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*appsv1beta1.StatefulSet))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get StatefulSets.
func (s *statefulSetClusterLister) Cluster(cluster logicalcluster.Name) appsv1beta1listers.StatefulSetLister {
	return &statefulSetLister{indexer: s.indexer, cluster: cluster}
}

// statefulSetLister implements the appsv1beta1listers.StatefulSetLister interface.
type statefulSetLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all StatefulSets in the indexer for a workspace.
func (s *statefulSetLister) List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.cluster, selector, func(i interface{}) {
		ret = append(ret, i.(*appsv1beta1.StatefulSet))
	})
	return ret, err
}

// StatefulSets returns an object that can list and get StatefulSets in one namespace.
func (s *statefulSetLister) StatefulSets(namespace string) appsv1beta1listers.StatefulSetNamespaceLister {
	return &statefulSetNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// statefulSetNamespaceLister implements the appsv1beta1listers.StatefulSetNamespaceLister interface.
type statefulSetNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all StatefulSets in the indexer for a given workspace and namespace.
func (s *statefulSetNamespaceLister) List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.cluster, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*appsv1beta1.StatefulSet))
	})
	return ret, err
}

// Get retrieves the StatefulSet from the indexer for a given workspace, namespace and name.
func (s *statefulSetNamespaceLister) Get(name string) (*appsv1beta1.StatefulSet, error) {
	key := kcpcache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(appsv1beta1.Resource("StatefulSet"), name)
	}
	return obj.(*appsv1beta1.StatefulSet), nil
}
