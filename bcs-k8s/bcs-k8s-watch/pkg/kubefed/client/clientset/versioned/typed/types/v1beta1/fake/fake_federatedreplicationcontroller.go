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
	v1beta1 "bk-bcs/bcs-k8s/bcs-k8s-watch/pkg/kubefed/apis/types/v1beta1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedReplicationControllers implements FederatedReplicationControllerInterface
type FakeFederatedReplicationControllers struct {
	Fake *FakeTypesV1beta1
	ns   string
}

var federatedreplicationcontrollersResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federatedreplicationcontrollers"}

var federatedreplicationcontrollersKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedReplicationController"}

// Get takes name of the federatedReplicationController, and returns the corresponding federatedReplicationController object, and an error if there is any.
func (c *FakeFederatedReplicationControllers) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedReplicationController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedreplicationcontrollersResource, c.ns, name), &v1beta1.FederatedReplicationController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedReplicationController), err
}

// List takes label and field selectors, and returns the list of FederatedReplicationControllers that match those selectors.
func (c *FakeFederatedReplicationControllers) List(opts v1.ListOptions) (result *v1beta1.FederatedReplicationControllerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedreplicationcontrollersResource, federatedreplicationcontrollersKind, c.ns, opts), &v1beta1.FederatedReplicationControllerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedReplicationControllerList{}
	for _, item := range obj.(*v1beta1.FederatedReplicationControllerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedReplicationControllers.
func (c *FakeFederatedReplicationControllers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedreplicationcontrollersResource, c.ns, opts))

}

// Create takes the representation of a federatedReplicationController and creates it.  Returns the server's representation of the federatedReplicationController, and an error, if there is any.
func (c *FakeFederatedReplicationControllers) Create(federatedReplicationController *v1beta1.FederatedReplicationController) (result *v1beta1.FederatedReplicationController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedreplicationcontrollersResource, c.ns, federatedReplicationController), &v1beta1.FederatedReplicationController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedReplicationController), err
}

// Update takes the representation of a federatedReplicationController and updates it. Returns the server's representation of the federatedReplicationController, and an error, if there is any.
func (c *FakeFederatedReplicationControllers) Update(federatedReplicationController *v1beta1.FederatedReplicationController) (result *v1beta1.FederatedReplicationController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedreplicationcontrollersResource, c.ns, federatedReplicationController), &v1beta1.FederatedReplicationController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedReplicationController), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedReplicationControllers) UpdateStatus(federatedReplicationController *v1beta1.FederatedReplicationController) (*v1beta1.FederatedReplicationController, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedreplicationcontrollersResource, "status", c.ns, federatedReplicationController), &v1beta1.FederatedReplicationController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedReplicationController), err
}

// Delete takes name of the federatedReplicationController and deletes it. Returns an error if one occurs.
func (c *FakeFederatedReplicationControllers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedreplicationcontrollersResource, c.ns, name), &v1beta1.FederatedReplicationController{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedReplicationControllers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedreplicationcontrollersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedReplicationControllerList{})
	return err
}

// Patch applies the patch and returns the patched federatedReplicationController.
func (c *FakeFederatedReplicationControllers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedReplicationController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedreplicationcontrollersResource, c.ns, name, data, subresources...), &v1beta1.FederatedReplicationController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedReplicationController), err
}
