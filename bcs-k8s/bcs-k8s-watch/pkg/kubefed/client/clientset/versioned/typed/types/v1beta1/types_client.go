/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"
	"bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/clientset/versioned/scheme"

	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type TypesV1beta1Interface interface {
	RESTClient() rest.Interface
	FederatedClusterRolesGetter
	FederatedClusterRoleBindingsGetter
	FederatedConfigMapsGetter
	FederatedDaemonSetsGetter
	FederatedDeploymentsGetter
	FederatedEndpointsesGetter
	FederatedIngressesGetter
	FederatedJobsGetter
	FederatedNamespacesGetter
	FederatedReplicaSetsGetter
	FederatedReplicationControllersGetter
	FederatedRolesGetter
	FederatedRoleBindingsGetter
	FederatedSecretsGetter
	FederatedServicesGetter
	FederatedServiceAccountsGetter
	FederatedStatefulSetsGetter
}

// TypesV1beta1Client is used to interact with features provided by the types.kubefed.io group.
type TypesV1beta1Client struct {
	restClient rest.Interface
}

func (c *TypesV1beta1Client) FederatedClusterRoles(namespace string) FederatedClusterRoleInterface {
	return newFederatedClusterRoles(c, namespace)
}

func (c *TypesV1beta1Client) FederatedClusterRoleBindings(namespace string) FederatedClusterRoleBindingInterface {
	return newFederatedClusterRoleBindings(c, namespace)
}

func (c *TypesV1beta1Client) FederatedConfigMaps(namespace string) FederatedConfigMapInterface {
	return newFederatedConfigMaps(c, namespace)
}

func (c *TypesV1beta1Client) FederatedDaemonSets(namespace string) FederatedDaemonSetInterface {
	return newFederatedDaemonSets(c, namespace)
}

func (c *TypesV1beta1Client) FederatedDeployments(namespace string) FederatedDeploymentInterface {
	return newFederatedDeployments(c, namespace)
}

func (c *TypesV1beta1Client) FederatedEndpointses(namespace string) FederatedEndpointsInterface {
	return newFederatedEndpointses(c, namespace)
}

func (c *TypesV1beta1Client) FederatedIngresses(namespace string) FederatedIngressInterface {
	return newFederatedIngresses(c, namespace)
}

func (c *TypesV1beta1Client) FederatedJobs(namespace string) FederatedJobInterface {
	return newFederatedJobs(c, namespace)
}

func (c *TypesV1beta1Client) FederatedNamespaces(namespace string) FederatedNamespaceInterface {
	return newFederatedNamespaces(c, namespace)
}

func (c *TypesV1beta1Client) FederatedReplicaSets(namespace string) FederatedReplicaSetInterface {
	return newFederatedReplicaSets(c, namespace)
}

func (c *TypesV1beta1Client) FederatedReplicationControllers(namespace string) FederatedReplicationControllerInterface {
	return newFederatedReplicationControllers(c, namespace)
}

func (c *TypesV1beta1Client) FederatedRoles(namespace string) FederatedRoleInterface {
	return newFederatedRoles(c, namespace)
}

func (c *TypesV1beta1Client) FederatedRoleBindings(namespace string) FederatedRoleBindingInterface {
	return newFederatedRoleBindings(c, namespace)
}

func (c *TypesV1beta1Client) FederatedSecrets(namespace string) FederatedSecretInterface {
	return newFederatedSecrets(c, namespace)
}

func (c *TypesV1beta1Client) FederatedServices(namespace string) FederatedServiceInterface {
	return newFederatedServices(c, namespace)
}

func (c *TypesV1beta1Client) FederatedServiceAccounts(namespace string) FederatedServiceAccountInterface {
	return newFederatedServiceAccounts(c, namespace)
}

func (c *TypesV1beta1Client) FederatedStatefulSets(namespace string) FederatedStatefulSetInterface {
	return newFederatedStatefulSets(c, namespace)
}

// NewForConfig creates a new TypesV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*TypesV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TypesV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new TypesV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TypesV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TypesV1beta1Client for the given RESTClient.
func New(c rest.Interface) *TypesV1beta1Client {
	return &TypesV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TypesV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
