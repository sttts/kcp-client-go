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
	"github.com/kcp-dev/logicalcluster/v3"

	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"

	kcpcorev1 "github.com/kcp-dev/client-go/kubernetes/typed/core/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpcorev1.CoreV1ClusterInterface = (*CoreV1ClusterClient)(nil)

type CoreV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *CoreV1ClusterClient) Cluster(cluster logicalcluster.Path) corev1.CoreV1Interface {
	if cluster == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &CoreV1Client{Fake: c.Fake, Cluster: cluster}
}

func (c *CoreV1ClusterClient) PersistentVolumes() kcpcorev1.PersistentVolumeClusterInterface {
	return &persistentVolumesClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) PersistentVolumeClaims() kcpcorev1.PersistentVolumeClaimClusterInterface {
	return &persistentVolumeClaimsClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Pods() kcpcorev1.PodClusterInterface {
	return &podsClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) PodTemplates() kcpcorev1.PodTemplateClusterInterface {
	return &podTemplatesClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) ReplicationControllers() kcpcorev1.ReplicationControllerClusterInterface {
	return &replicationControllersClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Services() kcpcorev1.ServiceClusterInterface {
	return &servicesClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) ServiceAccounts() kcpcorev1.ServiceAccountClusterInterface {
	return &serviceAccountsClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Endpoints() kcpcorev1.EndpointsClusterInterface {
	return &endpointsClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Nodes() kcpcorev1.NodeClusterInterface {
	return &nodesClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Namespaces() kcpcorev1.NamespaceClusterInterface {
	return &namespacesClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Events() kcpcorev1.EventClusterInterface {
	return &eventsClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) LimitRanges() kcpcorev1.LimitRangeClusterInterface {
	return &limitRangesClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) ResourceQuotas() kcpcorev1.ResourceQuotaClusterInterface {
	return &resourceQuotasClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Secrets() kcpcorev1.SecretClusterInterface {
	return &secretsClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) ConfigMaps() kcpcorev1.ConfigMapClusterInterface {
	return &configMapsClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) ComponentStatuses() kcpcorev1.ComponentStatusClusterInterface {
	return &componentStatusesClusterClient{Fake: c.Fake}
}

var _ corev1.CoreV1Interface = (*CoreV1Client)(nil)

type CoreV1Client struct {
	*kcptesting.Fake
	Cluster logicalcluster.Path
}

func (c *CoreV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *CoreV1Client) PersistentVolumes() corev1.PersistentVolumeInterface {
	return &persistentVolumesClient{Fake: c.Fake, Cluster: c.Cluster}
}

func (c *CoreV1Client) PersistentVolumeClaims(namespace string) corev1.PersistentVolumeClaimInterface {
	return &persistentVolumeClaimsClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *CoreV1Client) Pods(namespace string) corev1.PodInterface {
	return &podsClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *CoreV1Client) PodTemplates(namespace string) corev1.PodTemplateInterface {
	return &podTemplatesClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *CoreV1Client) ReplicationControllers(namespace string) corev1.ReplicationControllerInterface {
	return &replicationControllersClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *CoreV1Client) Services(namespace string) corev1.ServiceInterface {
	return &servicesClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *CoreV1Client) ServiceAccounts(namespace string) corev1.ServiceAccountInterface {
	return &serviceAccountsClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *CoreV1Client) Endpoints(namespace string) corev1.EndpointsInterface {
	return &endpointsClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *CoreV1Client) Nodes() corev1.NodeInterface {
	return &nodesClient{Fake: c.Fake, Cluster: c.Cluster}
}

func (c *CoreV1Client) Namespaces() corev1.NamespaceInterface {
	return &namespacesClient{Fake: c.Fake, Cluster: c.Cluster}
}

func (c *CoreV1Client) Events(namespace string) corev1.EventInterface {
	return &eventsClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *CoreV1Client) LimitRanges(namespace string) corev1.LimitRangeInterface {
	return &limitRangesClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *CoreV1Client) ResourceQuotas(namespace string) corev1.ResourceQuotaInterface {
	return &resourceQuotasClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *CoreV1Client) Secrets(namespace string) corev1.SecretInterface {
	return &secretsClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *CoreV1Client) ConfigMaps(namespace string) corev1.ConfigMapInterface {
	return &configMapsClient{Fake: c.Fake, Cluster: c.Cluster, Namespace: namespace}
}

func (c *CoreV1Client) ComponentStatuses() corev1.ComponentStatusInterface {
	return &componentStatusesClient{Fake: c.Fake, Cluster: c.Cluster}
}
