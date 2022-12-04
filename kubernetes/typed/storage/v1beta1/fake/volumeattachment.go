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

	storagev1beta1 "k8s.io/api/storage/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsstoragev1beta1 "k8s.io/client-go/applyconfigurations/storage/v1beta1"
	storagev1beta1client "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var volumeAttachmentsResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1beta1", Resource: "volumeattachments"}
var volumeAttachmentsKind = schema.GroupVersionKind{Group: "storage.k8s.io", Version: "v1beta1", Kind: "VolumeAttachment"}

type volumeAttachmentsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *volumeAttachmentsClusterClient) Cluster(cluster logicalcluster.Path) storagev1beta1client.VolumeAttachmentInterface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &volumeAttachmentsClient{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of VolumeAttachments that match those selectors across all clusters.
func (c *volumeAttachmentsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*storagev1beta1.VolumeAttachmentList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(volumeAttachmentsResource, volumeAttachmentsKind, logicalcluster.Wildcard, opts), &storagev1beta1.VolumeAttachmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1beta1.VolumeAttachmentList{ListMeta: obj.(*storagev1beta1.VolumeAttachmentList).ListMeta}
	for _, item := range obj.(*storagev1beta1.VolumeAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested VolumeAttachments across all clusters.
func (c *volumeAttachmentsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(volumeAttachmentsResource, logicalcluster.Wildcard, opts))
}

type volumeAttachmentsClient struct {
	*kcptesting.Fake
	Cluster logicalcluster.Path
}

func (c *volumeAttachmentsClient) Create(ctx context.Context, volumeAttachment *storagev1beta1.VolumeAttachment, opts metav1.CreateOptions) (*storagev1beta1.VolumeAttachment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(volumeAttachmentsResource, c.Cluster, volumeAttachment), &storagev1beta1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttachment), err
}

func (c *volumeAttachmentsClient) Update(ctx context.Context, volumeAttachment *storagev1beta1.VolumeAttachment, opts metav1.UpdateOptions) (*storagev1beta1.VolumeAttachment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(volumeAttachmentsResource, c.Cluster, volumeAttachment), &storagev1beta1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttachment), err
}

func (c *volumeAttachmentsClient) UpdateStatus(ctx context.Context, volumeAttachment *storagev1beta1.VolumeAttachment, opts metav1.UpdateOptions) (*storagev1beta1.VolumeAttachment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(volumeAttachmentsResource, c.Cluster, "status", volumeAttachment), &storagev1beta1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttachment), err
}

func (c *volumeAttachmentsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(volumeAttachmentsResource, c.Cluster, name, opts), &storagev1beta1.VolumeAttachment{})
	return err
}

func (c *volumeAttachmentsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(volumeAttachmentsResource, c.Cluster, listOpts)

	_, err := c.Fake.Invokes(action, &storagev1beta1.VolumeAttachmentList{})
	return err
}

func (c *volumeAttachmentsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*storagev1beta1.VolumeAttachment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(volumeAttachmentsResource, c.Cluster, name), &storagev1beta1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttachment), err
}

// List takes label and field selectors, and returns the list of VolumeAttachments that match those selectors.
func (c *volumeAttachmentsClient) List(ctx context.Context, opts metav1.ListOptions) (*storagev1beta1.VolumeAttachmentList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(volumeAttachmentsResource, volumeAttachmentsKind, c.Cluster, opts), &storagev1beta1.VolumeAttachmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1beta1.VolumeAttachmentList{ListMeta: obj.(*storagev1beta1.VolumeAttachmentList).ListMeta}
	for _, item := range obj.(*storagev1beta1.VolumeAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *volumeAttachmentsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(volumeAttachmentsResource, c.Cluster, opts))
}

func (c *volumeAttachmentsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*storagev1beta1.VolumeAttachment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(volumeAttachmentsResource, c.Cluster, name, pt, data, subresources...), &storagev1beta1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttachment), err
}

func (c *volumeAttachmentsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsstoragev1beta1.VolumeAttachmentApplyConfiguration, opts metav1.ApplyOptions) (*storagev1beta1.VolumeAttachment, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(volumeAttachmentsResource, c.Cluster, *name, types.ApplyPatchType, data), &storagev1beta1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttachment), err
}

func (c *volumeAttachmentsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsstoragev1beta1.VolumeAttachmentApplyConfiguration, opts metav1.ApplyOptions) (*storagev1beta1.VolumeAttachment, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(volumeAttachmentsResource, c.Cluster, *name, types.ApplyPatchType, data, "status"), &storagev1beta1.VolumeAttachment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttachment), err
}
