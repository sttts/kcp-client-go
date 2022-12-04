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

	coordinationv1beta1 "k8s.io/api/coordination/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationscoordinationv1beta1 "k8s.io/client-go/applyconfigurations/coordination/v1beta1"
	coordinationv1beta1client "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	"k8s.io/client-go/testing"

	kcpcoordinationv1beta1 "github.com/kcp-dev/client-go/kubernetes/typed/coordination/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var leasesResource = schema.GroupVersionResource{Group: "coordination.k8s.io", Version: "v1beta1", Resource: "leases"}
var leasesKind = schema.GroupVersionKind{Group: "coordination.k8s.io", Version: "v1beta1", Kind: "Lease"}

type leasesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *leasesClusterClient) Cluster(cluster logicalcluster.Path) kcpcoordinationv1beta1.LeasesNamespacer {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &leasesNamespacer{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of Leases that match those selectors across all clusters.
func (c *leasesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*coordinationv1beta1.LeaseList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(leasesResource, leasesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &coordinationv1beta1.LeaseList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &coordinationv1beta1.LeaseList{ListMeta: obj.(*coordinationv1beta1.LeaseList).ListMeta}
	for _, item := range obj.(*coordinationv1beta1.LeaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested Leases across all clusters.
func (c *leasesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(leasesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type leasesNamespacer struct {
	*kcptesting.Fake
	Cluster logicalcluster.Path
}

func (n *leasesNamespacer) Namespace(namespace string) coordinationv1beta1client.LeaseInterface {
	return &leasesClient{Fake: n.Fake, Cluster: n.Cluster, Namespace: namespace}
}

type leasesClient struct {
	*kcptesting.Fake
	Cluster   logicalcluster.Path
	Namespace string
}

func (c *leasesClient) Create(ctx context.Context, lease *coordinationv1beta1.Lease, opts metav1.CreateOptions) (*coordinationv1beta1.Lease, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(leasesResource, c.Cluster, c.Namespace, lease), &coordinationv1beta1.Lease{})
	if obj == nil {
		return nil, err
	}
	return obj.(*coordinationv1beta1.Lease), err
}

func (c *leasesClient) Update(ctx context.Context, lease *coordinationv1beta1.Lease, opts metav1.UpdateOptions) (*coordinationv1beta1.Lease, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(leasesResource, c.Cluster, c.Namespace, lease), &coordinationv1beta1.Lease{})
	if obj == nil {
		return nil, err
	}
	return obj.(*coordinationv1beta1.Lease), err
}

func (c *leasesClient) UpdateStatus(ctx context.Context, lease *coordinationv1beta1.Lease, opts metav1.UpdateOptions) (*coordinationv1beta1.Lease, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(leasesResource, c.Cluster, "status", c.Namespace, lease), &coordinationv1beta1.Lease{})
	if obj == nil {
		return nil, err
	}
	return obj.(*coordinationv1beta1.Lease), err
}

func (c *leasesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(leasesResource, c.Cluster, c.Namespace, name, opts), &coordinationv1beta1.Lease{})
	return err
}

func (c *leasesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(leasesResource, c.Cluster, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &coordinationv1beta1.LeaseList{})
	return err
}

func (c *leasesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*coordinationv1beta1.Lease, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(leasesResource, c.Cluster, c.Namespace, name), &coordinationv1beta1.Lease{})
	if obj == nil {
		return nil, err
	}
	return obj.(*coordinationv1beta1.Lease), err
}

// List takes label and field selectors, and returns the list of Leases that match those selectors.
func (c *leasesClient) List(ctx context.Context, opts metav1.ListOptions) (*coordinationv1beta1.LeaseList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(leasesResource, leasesKind, c.Cluster, c.Namespace, opts), &coordinationv1beta1.LeaseList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &coordinationv1beta1.LeaseList{ListMeta: obj.(*coordinationv1beta1.LeaseList).ListMeta}
	for _, item := range obj.(*coordinationv1beta1.LeaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *leasesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(leasesResource, c.Cluster, c.Namespace, opts))
}

func (c *leasesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*coordinationv1beta1.Lease, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(leasesResource, c.Cluster, c.Namespace, name, pt, data, subresources...), &coordinationv1beta1.Lease{})
	if obj == nil {
		return nil, err
	}
	return obj.(*coordinationv1beta1.Lease), err
}

func (c *leasesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationscoordinationv1beta1.LeaseApplyConfiguration, opts metav1.ApplyOptions) (*coordinationv1beta1.Lease, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(leasesResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data), &coordinationv1beta1.Lease{})
	if obj == nil {
		return nil, err
	}
	return obj.(*coordinationv1beta1.Lease), err
}

func (c *leasesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationscoordinationv1beta1.LeaseApplyConfiguration, opts metav1.ApplyOptions) (*coordinationv1beta1.Lease, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(leasesResource, c.Cluster, c.Namespace, *name, types.ApplyPatchType, data, "status"), &coordinationv1beta1.Lease{})
	if obj == nil {
		return nil, err
	}
	return obj.(*coordinationv1beta1.Lease), err
}
