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

package fake

import (
	v1alpha1 "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/core/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterPropagatedVersions implements ClusterPropagatedVersionInterface
type FakeClusterPropagatedVersions struct {
	Fake *FakeCoreV1alpha1
}

var clusterpropagatedversionsResource = schema.GroupVersionResource{Group: "core.kubefed.io", Version: "v1alpha1", Resource: "clusterpropagatedversions"}

var clusterpropagatedversionsKind = schema.GroupVersionKind{Group: "core.kubefed.io", Version: "v1alpha1", Kind: "ClusterPropagatedVersion"}

// Get takes name of the clusterPropagatedVersion, and returns the corresponding clusterPropagatedVersion object, and an error if there is any.
func (c *FakeClusterPropagatedVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterPropagatedVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterpropagatedversionsResource, name), &v1alpha1.ClusterPropagatedVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPropagatedVersion), err
}

// List takes label and field selectors, and returns the list of ClusterPropagatedVersions that match those selectors.
func (c *FakeClusterPropagatedVersions) List(opts v1.ListOptions) (result *v1alpha1.ClusterPropagatedVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterpropagatedversionsResource, clusterpropagatedversionsKind, opts), &v1alpha1.ClusterPropagatedVersionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterPropagatedVersionList{}
	for _, item := range obj.(*v1alpha1.ClusterPropagatedVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterPropagatedVersions.
func (c *FakeClusterPropagatedVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterpropagatedversionsResource, opts))
}

// Create takes the representation of a clusterPropagatedVersion and creates it.  Returns the server's representation of the clusterPropagatedVersion, and an error, if there is any.
func (c *FakeClusterPropagatedVersions) Create(clusterPropagatedVersion *v1alpha1.ClusterPropagatedVersion) (result *v1alpha1.ClusterPropagatedVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterpropagatedversionsResource, clusterPropagatedVersion), &v1alpha1.ClusterPropagatedVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPropagatedVersion), err
}

// Update takes the representation of a clusterPropagatedVersion and updates it. Returns the server's representation of the clusterPropagatedVersion, and an error, if there is any.
func (c *FakeClusterPropagatedVersions) Update(clusterPropagatedVersion *v1alpha1.ClusterPropagatedVersion) (result *v1alpha1.ClusterPropagatedVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterpropagatedversionsResource, clusterPropagatedVersion), &v1alpha1.ClusterPropagatedVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPropagatedVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterPropagatedVersions) UpdateStatus(clusterPropagatedVersion *v1alpha1.ClusterPropagatedVersion) (*v1alpha1.ClusterPropagatedVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterpropagatedversionsResource, "status", clusterPropagatedVersion), &v1alpha1.ClusterPropagatedVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPropagatedVersion), err
}

// Delete takes name of the clusterPropagatedVersion and deletes it. Returns an error if one occurs.
func (c *FakeClusterPropagatedVersions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterpropagatedversionsResource, name), &v1alpha1.ClusterPropagatedVersion{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterPropagatedVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterpropagatedversionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterPropagatedVersionList{})
	return err
}

// Patch applies the patch and returns the patched clusterPropagatedVersion.
func (c *FakeClusterPropagatedVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterPropagatedVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterpropagatedversionsResource, name, data, subresources...), &v1alpha1.ClusterPropagatedVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPropagatedVersion), err
}
