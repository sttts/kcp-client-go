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
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	appsv1beta1 "k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsappsv1beta1 "k8s.io/client-go/applyconfigurations/apps/v1beta1"
	appsv1beta1client "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	"k8s.io/client-go/testing"

	kcpappsv1beta1 "github.com/kcp-dev/client-go/kubernetes/typed/apps/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var statefulSetsResource = schema.GroupVersionResource{Group: "apps", Version: "v1beta1", Resource: "statefulsets"}
var statefulSetsKind = schema.GroupVersionKind{Group: "apps", Version: "v1beta1", Kind: "StatefulSet"}

type statefulSetsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *statefulSetsClusterClient) Cluster(cluster logicalcluster.Path) kcpappsv1beta1.StatefulSetsNamespacer {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &statefulSetsNamespacer{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of StatefulSets that match those selectors across all clusters.
func (c *statefulSetsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*appsv1beta1.StatefulSetList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(statefulSetsResource, statefulSetsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &appsv1beta1.StatefulSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1beta1.StatefulSetList{ListMeta: obj.(*appsv1beta1.StatefulSetList).ListMeta}
	for _, item := range obj.(*appsv1beta1.StatefulSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested StatefulSets across all clusters.
func (c *statefulSetsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(statefulSetsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type statefulSetsNamespacer struct {
	*kcptesting.Fake
	Cluster logicalcluster.Path
}

func (n *statefulSetsNamespacer) Namespace(namespace string) appsv1beta1client.StatefulSetInterface {
	return &statefulSetsClient{Fake: n.Fake, Cluster: n.Cluster, Namespace: namespace}
}

type statefulSetsClient struct {
	*kcptesting.Fake
	Cluster   logicalcluster.Path
	Namespace string
}

func (c *statefulSetsClient) Create(ctx context.Context, statefulSet *appsv1beta1.StatefulSet, opts metav1.CreateOptions) (*appsv1beta1.StatefulSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(statefulSetsResource, c.Cluster, c.Namespace, statefulSet), &appsv1beta1.StatefulSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.StatefulSet), err
}

func (c *statefulSetsClient) Update(ctx context.Context, statefulSet *appsv1beta1.StatefulSet, opts metav1.UpdateOptions) (*appsv1beta1.StatefulSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(statefulSetsResource, c.Cluster, c.Namespace, statefulSet), &appsv1beta1.StatefulSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.StatefulSet), err
}

func (c *statefulSetsClient) UpdateStatus(ctx context.Context, statefulSet *appsv1beta1.StatefulSet, opts metav1.UpdateOptions) (*appsv1beta1.StatefulSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(statefulSetsResource, c.Cluster, "status", c.Namespace, statefulSet), &appsv1beta1.StatefulSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.StatefulSet), err
}

func (c *statefulSetsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(statefulSetsResource, c.Cluster, c.Namespace, name, opts), &appsv1beta1.StatefulSet{})
	return err
}

func (c *statefulSetsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(statefulSetsResource, c.Cluster, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &appsv1beta1.StatefulSetList{})
	return err
}

func (c *statefulSetsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*appsv1beta1.StatefulSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(statefulSetsResource, c.Cluster, c.Namespace, name), &appsv1beta1.StatefulSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.StatefulSet), err
}

// List takes label and field selectors, and returns the list of StatefulSets that match those selectors.
func (c *statefulSetsClient) List(ctx context.Context, opts metav1.ListOptions) (*appsv1beta1.StatefulSetList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(statefulSetsResource, statefulSetsKind, c.Cluster, c.Namespace, opts), &appsv1beta1.StatefulSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1beta1.StatefulSetList{ListMeta: obj.(*appsv1beta1.StatefulSetList).ListMeta}
	for _, item := range obj.(*appsv1beta1.StatefulSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *statefulSetsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(statefulSetsResource, c.Cluster, c.Namespace, opts))
}

func (c *statefulSetsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*appsv1beta1.StatefulSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(statefulSetsResource, c.Cluster, c.Namespace, name, pt, data, subresources...), &appsv1beta1.StatefulSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.StatefulSet), err
}

func (c *statefulSetsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsappsv1beta1.StatefulSetApplyConfiguration, opts metav1.ApplyOptions) (*appsv1beta1.StatefulSet, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(statefulSetsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data), &appsv1beta1.StatefulSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.StatefulSet), err
}

func (c *statefulSetsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsappsv1beta1.StatefulSetApplyConfiguration, opts metav1.ApplyOptions) (*appsv1beta1.StatefulSet, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(statefulSetsResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data, "status"), &appsv1beta1.StatefulSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1beta1.StatefulSet), err
}
