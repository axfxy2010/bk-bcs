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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// FederatedClusterRoles returns a FederatedClusterRoleInformer.
	FederatedClusterRoles() FederatedClusterRoleInformer
	// FederatedClusterRoleBindings returns a FederatedClusterRoleBindingInformer.
	FederatedClusterRoleBindings() FederatedClusterRoleBindingInformer
	// FederatedConfigMaps returns a FederatedConfigMapInformer.
	FederatedConfigMaps() FederatedConfigMapInformer
	// FederatedDaemonSets returns a FederatedDaemonSetInformer.
	FederatedDaemonSets() FederatedDaemonSetInformer
	// FederatedDeployments returns a FederatedDeploymentInformer.
	FederatedDeployments() FederatedDeploymentInformer
	// FederatedEndpointses returns a FederatedEndpointsInformer.
	FederatedEndpointses() FederatedEndpointsInformer
	// FederatedIngresses returns a FederatedIngressInformer.
	FederatedIngresses() FederatedIngressInformer
	// FederatedJobs returns a FederatedJobInformer.
	FederatedJobs() FederatedJobInformer
	// FederatedNamespaces returns a FederatedNamespaceInformer.
	FederatedNamespaces() FederatedNamespaceInformer
	// FederatedReplicaSets returns a FederatedReplicaSetInformer.
	FederatedReplicaSets() FederatedReplicaSetInformer
	// FederatedReplicationControllers returns a FederatedReplicationControllerInformer.
	FederatedReplicationControllers() FederatedReplicationControllerInformer
	// FederatedRoles returns a FederatedRoleInformer.
	FederatedRoles() FederatedRoleInformer
	// FederatedRoleBindings returns a FederatedRoleBindingInformer.
	FederatedRoleBindings() FederatedRoleBindingInformer
	// FederatedSecrets returns a FederatedSecretInformer.
	FederatedSecrets() FederatedSecretInformer
	// FederatedServices returns a FederatedServiceInformer.
	FederatedServices() FederatedServiceInformer
	// FederatedServiceAccounts returns a FederatedServiceAccountInformer.
	FederatedServiceAccounts() FederatedServiceAccountInformer
	// FederatedStatefulSets returns a FederatedStatefulSetInformer.
	FederatedStatefulSets() FederatedStatefulSetInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// FederatedClusterRoles returns a FederatedClusterRoleInformer.
func (v *version) FederatedClusterRoles() FederatedClusterRoleInformer {
	return &federatedClusterRoleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedClusterRoleBindings returns a FederatedClusterRoleBindingInformer.
func (v *version) FederatedClusterRoleBindings() FederatedClusterRoleBindingInformer {
	return &federatedClusterRoleBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedConfigMaps returns a FederatedConfigMapInformer.
func (v *version) FederatedConfigMaps() FederatedConfigMapInformer {
	return &federatedConfigMapInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedDaemonSets returns a FederatedDaemonSetInformer.
func (v *version) FederatedDaemonSets() FederatedDaemonSetInformer {
	return &federatedDaemonSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedDeployments returns a FederatedDeploymentInformer.
func (v *version) FederatedDeployments() FederatedDeploymentInformer {
	return &federatedDeploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedEndpointses returns a FederatedEndpointsInformer.
func (v *version) FederatedEndpointses() FederatedEndpointsInformer {
	return &federatedEndpointsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedIngresses returns a FederatedIngressInformer.
func (v *version) FederatedIngresses() FederatedIngressInformer {
	return &federatedIngressInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedJobs returns a FederatedJobInformer.
func (v *version) FederatedJobs() FederatedJobInformer {
	return &federatedJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedNamespaces returns a FederatedNamespaceInformer.
func (v *version) FederatedNamespaces() FederatedNamespaceInformer {
	return &federatedNamespaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedReplicaSets returns a FederatedReplicaSetInformer.
func (v *version) FederatedReplicaSets() FederatedReplicaSetInformer {
	return &federatedReplicaSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedReplicationControllers returns a FederatedReplicationControllerInformer.
func (v *version) FederatedReplicationControllers() FederatedReplicationControllerInformer {
	return &federatedReplicationControllerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedRoles returns a FederatedRoleInformer.
func (v *version) FederatedRoles() FederatedRoleInformer {
	return &federatedRoleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedRoleBindings returns a FederatedRoleBindingInformer.
func (v *version) FederatedRoleBindings() FederatedRoleBindingInformer {
	return &federatedRoleBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedSecrets returns a FederatedSecretInformer.
func (v *version) FederatedSecrets() FederatedSecretInformer {
	return &federatedSecretInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedServices returns a FederatedServiceInformer.
func (v *version) FederatedServices() FederatedServiceInformer {
	return &federatedServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedServiceAccounts returns a FederatedServiceAccountInformer.
func (v *version) FederatedServiceAccounts() FederatedServiceAccountInformer {
	return &federatedServiceAccountInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedStatefulSets returns a FederatedStatefulSetInformer.
func (v *version) FederatedStatefulSets() FederatedStatefulSetInformer {
	return &federatedStatefulSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
