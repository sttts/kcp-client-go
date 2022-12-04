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

package v1

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	certificatesv1 "k8s.io/api/certificates/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationscertificatesv1 "k8s.io/client-go/applyconfigurations/certificates/v1"
	certificatesv1client "k8s.io/client-go/kubernetes/typed/certificates/v1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var certificateSigningRequestsResource = schema.GroupVersionResource{Group: "certificates.k8s.io", Version: "v1", Resource: "certificatesigningrequests"}
var certificateSigningRequestsKind = schema.GroupVersionKind{Group: "certificates.k8s.io", Version: "v1", Kind: "CertificateSigningRequest"}

type certificateSigningRequestsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *certificateSigningRequestsClusterClient) Cluster(cluster logicalcluster.Path) certificatesv1client.CertificateSigningRequestInterface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &certificateSigningRequestsClient{Fake: c.Fake, Cluster: cluster}
}

// List takes label and field selectors, and returns the list of CertificateSigningRequests that match those selectors across all clusters.
func (c *certificateSigningRequestsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*certificatesv1.CertificateSigningRequestList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(certificateSigningRequestsResource, certificateSigningRequestsKind, logicalcluster.Wildcard, opts), &certificatesv1.CertificateSigningRequestList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &certificatesv1.CertificateSigningRequestList{ListMeta: obj.(*certificatesv1.CertificateSigningRequestList).ListMeta}
	for _, item := range obj.(*certificatesv1.CertificateSigningRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested CertificateSigningRequests across all clusters.
func (c *certificateSigningRequestsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(certificateSigningRequestsResource, logicalcluster.Wildcard, opts))
}

type certificateSigningRequestsClient struct {
	*kcptesting.Fake
	Cluster logicalcluster.Path
}

func (c *certificateSigningRequestsClient) Create(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequest, opts metav1.CreateOptions) (*certificatesv1.CertificateSigningRequest, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(certificateSigningRequestsResource, c.Cluster, certificateSigningRequest), &certificatesv1.CertificateSigningRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*certificatesv1.CertificateSigningRequest), err
}

func (c *certificateSigningRequestsClient) Update(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequest, opts metav1.UpdateOptions) (*certificatesv1.CertificateSigningRequest, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(certificateSigningRequestsResource, c.Cluster, certificateSigningRequest), &certificatesv1.CertificateSigningRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*certificatesv1.CertificateSigningRequest), err
}

func (c *certificateSigningRequestsClient) UpdateStatus(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequest, opts metav1.UpdateOptions) (*certificatesv1.CertificateSigningRequest, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(certificateSigningRequestsResource, c.Cluster, "status", certificateSigningRequest), &certificatesv1.CertificateSigningRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*certificatesv1.CertificateSigningRequest), err
}

func (c *certificateSigningRequestsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(certificateSigningRequestsResource, c.Cluster, name, opts), &certificatesv1.CertificateSigningRequest{})
	return err
}

func (c *certificateSigningRequestsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(certificateSigningRequestsResource, c.Cluster, listOpts)

	_, err := c.Fake.Invokes(action, &certificatesv1.CertificateSigningRequestList{})
	return err
}

func (c *certificateSigningRequestsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*certificatesv1.CertificateSigningRequest, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(certificateSigningRequestsResource, c.Cluster, name), &certificatesv1.CertificateSigningRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*certificatesv1.CertificateSigningRequest), err
}

// List takes label and field selectors, and returns the list of CertificateSigningRequests that match those selectors.
func (c *certificateSigningRequestsClient) List(ctx context.Context, opts metav1.ListOptions) (*certificatesv1.CertificateSigningRequestList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(certificateSigningRequestsResource, certificateSigningRequestsKind, c.Cluster, opts), &certificatesv1.CertificateSigningRequestList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &certificatesv1.CertificateSigningRequestList{ListMeta: obj.(*certificatesv1.CertificateSigningRequestList).ListMeta}
	for _, item := range obj.(*certificatesv1.CertificateSigningRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *certificateSigningRequestsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(certificateSigningRequestsResource, c.Cluster, opts))
}

func (c *certificateSigningRequestsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*certificatesv1.CertificateSigningRequest, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(certificateSigningRequestsResource, c.Cluster, name, pt, data, subresources...), &certificatesv1.CertificateSigningRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*certificatesv1.CertificateSigningRequest), err
}

func (c *certificateSigningRequestsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationscertificatesv1.CertificateSigningRequestApplyConfiguration, opts metav1.ApplyOptions) (*certificatesv1.CertificateSigningRequest, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(certificateSigningRequestsResource, c.Cluster, *name, types.ApplyPatchType, data), &certificatesv1.CertificateSigningRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*certificatesv1.CertificateSigningRequest), err
}

func (c *certificateSigningRequestsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationscertificatesv1.CertificateSigningRequestApplyConfiguration, opts metav1.ApplyOptions) (*certificatesv1.CertificateSigningRequest, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(certificateSigningRequestsResource, c.Cluster, *name, types.ApplyPatchType, data, "status"), &certificatesv1.CertificateSigningRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*certificatesv1.CertificateSigningRequest), err
}

func (c *certificateSigningRequestsClient) UpdateApproval(ctx context.Context, certificateSigningRequestName string, certificateSigningRequest *certificatesv1.CertificateSigningRequest, opts metav1.UpdateOptions) (*certificatesv1.CertificateSigningRequest, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(certificateSigningRequestsResource, c.Cluster, "approval", certificateSigningRequest), &certificatesv1.CertificateSigningRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*certificatesv1.CertificateSigningRequest), err
}
