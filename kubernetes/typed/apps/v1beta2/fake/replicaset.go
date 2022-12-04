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
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	appsv1beta2 "k8s.io/api/apps/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsappsv1beta2 "k8s.io/client-go/applyconfigurations/apps/v1beta2"
	appsv1beta2client "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	"k8s.io/client-go/testing"

	kcpappsv1beta2 "github.com/kcp-dev/client-go/kubernetes/typed/apps/v1beta2"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var replicaSetsResource = schema.GroupVersionResource{Group: "apps", Version: "v1beta2", Resource: "replicasets"}
var replicaSetsKind = schema.GroupVersionKind{Group: "apps", Version: "v1beta2", Kind: "ReplicaSet"}

type replicaSetsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *replicaSetsClusterClient) Cluster(cluster logicalcluster.Path) kcpappsv1beta2.ReplicaSetsNamespacer {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &replicaSetsNamespacer{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of ReplicaSets that match those selectors across all clusters.
func (c *replicaSetsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*appsv1beta2.ReplicaSetList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(replicaSetsResource, replicaSetsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &appsv1beta2.ReplicaSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1beta2.ReplicaSetList{ListMeta: obj.(*appsv1beta2.ReplicaSetList).ListMeta}
	for _, item := range obj.(*appsv1beta2.ReplicaSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ReplicaSets across all clusters.
func (c *replicaSetsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(replicaSetsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type replicaSetsNamespacer struct {
	*kcptesting.Fake
	Cluster logicalcluster.Path
}

func (n *replicaSetsNamespacer) Namespace(namespace string) appsv1beta2client.ReplicaSetInterface {
	return &replicaSetsClient{Fake: n.Fake, Cluster: n.Cluster, Namespace: namespace}
}

type replicaSetsClient struct {
	*kcptesting.Fake
	Cluster   logicalcluster.Path
	Namespace string
}

func (c *replicaSetsClient) Create(ctx context.Context, replicaSet *appsv1beta2.ReplicaSet, opts metav1.CreateOptions) (*appsv1beta2.ReplicaSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(replicaSetsResource, c.Cluster, c.Namespace, replicaSet), &appsv1beta2.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta2.ReplicaSet), err
}

func (c *replicaSetsClient) Update(ctx context.Context, replicaSet *appsv1beta2.ReplicaSet, opts metav1.UpdateOptions) (*appsv1beta2.ReplicaSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(replicaSetsResource, c.Cluster, c.Namespace, replicaSet), &appsv1beta2.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta2.ReplicaSet), err
}

func (c *replicaSetsClient) UpdateStatus(ctx context.Context, replicaSet *appsv1beta2.ReplicaSet, opts metav1.UpdateOptions) (*appsv1beta2.ReplicaSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(replicaSetsResource, c.Cluster, "status", c.Namespace, replicaSet), &appsv1beta2.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta2.ReplicaSet), err
}

func (c *replicaSetsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(replicaSetsResource, c.Cluster, c.Namespace, name, opts), &appsv1beta2.ReplicaSet{})
	return err
}

func (c *replicaSetsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(replicaSetsResource, c.Cluster, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &appsv1beta2.ReplicaSetList{})
	return err
}

func (c *replicaSetsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*appsv1beta2.ReplicaSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(replicaSetsResource, c.Cluster, c.Namespace, name), &appsv1beta2.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta2.ReplicaSet), err
}

// List takes label and field selectors, and returns the list of ReplicaSets that match those selectors.
func (c *replicaSetsClient) List(ctx context.Context, opts metav1.ListOptions) (*appsv1beta2.ReplicaSetList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(replicaSetsResource, replicaSetsKind, c.Cluster, c.Namespace, opts), &appsv1beta2.ReplicaSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1beta2.ReplicaSetList{ListMeta: obj.(*appsv1beta2.ReplicaSetList).ListMeta}
	for _, item := range obj.(*appsv1beta2.ReplicaSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *replicaSetsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(replicaSetsResource, c.Cluster, c.Namespace, opts))
}

func (c *replicaSetsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*appsv1beta2.ReplicaSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(replicaSetsResource, c.Cluster, c.Namespace, name, pt, data, subresources...), &appsv1beta2.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta2.ReplicaSet), err
}

func (c *replicaSetsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsappsv1beta2.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions) (*appsv1beta2.ReplicaSet, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(replicaSetsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data), &appsv1beta2.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta2.ReplicaSet), err
}

func (c *replicaSetsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsappsv1beta2.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions) (*appsv1beta2.ReplicaSet, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(replicaSetsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data, "status"), &appsv1beta2.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta2.ReplicaSet), err
}
