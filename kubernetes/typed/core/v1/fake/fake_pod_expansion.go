/*
Copyright 2014 The Kubernetes Authors.
Modifications Copyright 2022 The KCP Authors.

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

package v1

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	v1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	restclient "k8s.io/client-go/rest"
	fakerest "k8s.io/client-go/rest/fake"

	core "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

func (c *podsClient) Bind(ctx context.Context, binding *v1.Binding, opts metav1.CreateOptions) error {
	action := core.CreateActionImpl{}
	action.Verb = "create"
	action.Namespace = binding.Namespace
	action.Resource = podsResource
	action.Subresource = "binding"
	action.Object = binding
	action.Cluster = c.Cluster

	_, err := c.Fake.Invokes(action, binding)
	return err
}

func (c *podsClient) GetBinding(name string) (result *v1.Binding, err error) {
	obj, err := c.Fake.
		Invokes(core.NewGetSubresourceAction(podsResource, c.Cluster, c.Namespace, "binding", name), &v1.Binding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Binding), err
}

func (c *podsClient) GetLogs(name string, opts *v1.PodLogOptions) *restclient.Request {
	action := core.GenericActionImpl{}
	action.Verb = "get"
	action.Namespace = c.Namespace
	action.Resource = podsResource
	action.Subresource = "log"
	action.Value = opts
	action.Cluster = c.Cluster

	_, _ = c.Fake.Invokes(action, &v1.Pod{})
	fakeClient := &fakerest.RESTClient{
		Client: fakerest.CreateHTTPClient(func(request *http.Request) (*http.Response, error) {
			resp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("fake logs")),
			}
			return resp, nil
		}),
		NegotiatedSerializer: scheme.Codecs.WithoutConversion(),
		GroupVersion:         podsKind.GroupVersion(),
		VersionedAPIPath:     fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/log", c.Namespace, name),
	}
	return fakeClient.Request()
}

func (c *podsClient) Evict(ctx context.Context, eviction *policyv1beta1.Eviction) error {
	return c.EvictV1beta1(ctx, eviction)
}

func (c *podsClient) EvictV1(ctx context.Context, eviction *policyv1.Eviction) error {
	action := core.CreateActionImpl{}
	action.Verb = "create"
	action.Namespace = c.Namespace
	action.Resource = podsResource
	action.Subresource = "eviction"
	action.Object = eviction
	action.Cluster = c.Cluster

	_, err := c.Fake.Invokes(action, eviction)
	return err
}

func (c *podsClient) EvictV1beta1(ctx context.Context, eviction *policyv1beta1.Eviction) error {
	action := core.CreateActionImpl{}
	action.Verb = "create"
	action.Namespace = c.Namespace
	action.Resource = podsResource
	action.Subresource = "eviction"
	action.Object = eviction
	action.Cluster = c.Cluster

	_, err := c.Fake.Invokes(action, eviction)
	return err
}

func (c *podsClient) ProxyGet(scheme, name, port, path string, params map[string]string) restclient.ResponseWrapper {
	return c.Fake.InvokesProxy(core.NewProxyGetAction(podsResource, c.Cluster, c.Namespace, scheme, name, port, path, params))
}
