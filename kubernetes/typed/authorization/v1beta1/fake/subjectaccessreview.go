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

	"github.com/kcp-dev/logicalcluster/v3"

	authorizationv1beta1 "k8s.io/api/authorization/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	authorizationv1beta1client "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var subjectAccessReviewsResource = schema.GroupVersionResource{Group: "authorization.k8s.io", Version: "v1beta1", Resource: "subjectaccessreviews"}
var subjectAccessReviewsKind = schema.GroupVersionKind{Group: "authorization.k8s.io", Version: "v1beta1", Kind: "SubjectAccessReview"}

type subjectAccessReviewsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *subjectAccessReviewsClusterClient) Cluster(cluster logicalcluster.Path) authorizationv1beta1client.SubjectAccessReviewInterface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &subjectAccessReviewsClient{Fake: c.Fake, Cluster: cluster}
}

type subjectAccessReviewsClient struct {
	*kcptesting.Fake
	Cluster logicalcluster.Path
}

func (c *subjectAccessReviewsClient) Create(ctx context.Context, subjectAccessReview *authorizationv1beta1.SubjectAccessReview, opts metav1.CreateOptions) (*authorizationv1beta1.SubjectAccessReview, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(subjectAccessReviewsResource, c.Cluster, subjectAccessReview), &authorizationv1beta1.SubjectAccessReview{})
	if obj == nil {
		return nil, err
	}
	return obj.(*authorizationv1beta1.SubjectAccessReview), err
}
