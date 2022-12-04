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
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	rbacv1alpha1client "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
)

// RolesClusterGetter has a method to return a RoleClusterInterface.
// A group's cluster client should implement this interface.
type RolesClusterGetter interface {
	Roles() RoleClusterInterface
}

// RoleClusterInterface can operate on Roles across all clusters,
// or scope down to one cluster and return a RolesNamespacer.
type RoleClusterInterface interface {
	Cluster(logicalcluster.Path) RolesNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*rbacv1alpha1.RoleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type rolesClusterInterface struct {
	clientCache kcpclient.Cache[*rbacv1alpha1client.RbacV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *rolesClusterInterface) Cluster(path logicalcluster.Path) RolesNamespacer {
	if path == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &rolesNamespacer{clientCache: c.clientCache, path: path}
}

// List returns the entire collection of all Roles across all clusters.
func (c *rolesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1alpha1.RoleList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Roles(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all Roles across all clusters.
func (c *rolesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Roles(metav1.NamespaceAll).Watch(ctx, opts)
}

// RolesNamespacer can scope to objects within a namespace, returning a rbacv1alpha1client.RoleInterface.
type RolesNamespacer interface {
	Namespace(string) rbacv1alpha1client.RoleInterface
}

type rolesNamespacer struct {
	clientCache kcpclient.Cache[*rbacv1alpha1client.RbacV1alpha1Client]
	path        logicalcluster.Path
}

func (n *rolesNamespacer) Namespace(namespace string) rbacv1alpha1client.RoleInterface {
	return n.clientCache.ClusterOrDie(n.path).Roles(namespace)
}
