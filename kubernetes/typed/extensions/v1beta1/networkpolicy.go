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

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	extensionsv1beta1client "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

// NetworkPoliciesClusterGetter has a method to return a NetworkPolicyClusterInterface.
// A group's cluster client should implement this interface.
type NetworkPoliciesClusterGetter interface {
	NetworkPolicies() NetworkPolicyClusterInterface
}

// NetworkPolicyClusterInterface can operate on NetworkPolicies across all clusters,
// or scope down to one cluster and return a NetworkPoliciesNamespacer.
type NetworkPolicyClusterInterface interface {
	Cluster(logicalcluster.Path) NetworkPoliciesNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*extensionsv1beta1.NetworkPolicyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type networkPoliciesClusterInterface struct {
	clientCache kcpclient.Cache[*extensionsv1beta1client.ExtensionsV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *networkPoliciesClusterInterface) Cluster(path logicalcluster.Path) NetworkPoliciesNamespacer {
	if path == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &networkPoliciesNamespacer{clientCache: c.clientCache, path: path}
}

// List returns the entire collection of all NetworkPolicies across all clusters.
func (c *networkPoliciesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*extensionsv1beta1.NetworkPolicyList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).NetworkPolicies(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all NetworkPolicies across all clusters.
func (c *networkPoliciesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).NetworkPolicies(metav1.NamespaceAll).Watch(ctx, opts)
}

// NetworkPoliciesNamespacer can scope to objects within a namespace, returning a extensionsv1beta1client.NetworkPolicyInterface.
type NetworkPoliciesNamespacer interface {
	Namespace(string) extensionsv1beta1client.NetworkPolicyInterface
}

type networkPoliciesNamespacer struct {
	clientCache kcpclient.Cache[*extensionsv1beta1client.ExtensionsV1beta1Client]
	path        logicalcluster.Path
}

func (n *networkPoliciesNamespacer) Namespace(namespace string) extensionsv1beta1client.NetworkPolicyInterface {
	return n.clientCache.ClusterOrDie(n.path).NetworkPolicies(namespace)
}
