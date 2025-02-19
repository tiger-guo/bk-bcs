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
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-gamestatefulset-operator/pkg/apis/tkex/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	autoscaling "k8s.io/kubernetes/pkg/apis/autoscaling"
)

// FakeGameStatefulSets implements GameStatefulSetInterface
type FakeGameStatefulSets struct {
	Fake *FakeTkexV1alpha1
	ns   string
}

var gamestatefulsetsResource = schema.GroupVersionResource{Group: "tkex.tencent.com", Version: "v1alpha1", Resource: "gamestatefulsets"}

var gamestatefulsetsKind = schema.GroupVersionKind{Group: "tkex.tencent.com", Version: "v1alpha1", Kind: "GameStatefulSet"}

// Get takes name of the gameStatefulSet, and returns the corresponding gameStatefulSet object, and an error if there is any.
func (c *FakeGameStatefulSets) Get(name string, options v1.GetOptions) (result *v1alpha1.GameStatefulSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gamestatefulsetsResource, c.ns, name), &v1alpha1.GameStatefulSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameStatefulSet), err
}

// List takes label and field selectors, and returns the list of GameStatefulSets that match those selectors.
func (c *FakeGameStatefulSets) List(opts v1.ListOptions) (result *v1alpha1.GameStatefulSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gamestatefulsetsResource, gamestatefulsetsKind, c.ns, opts), &v1alpha1.GameStatefulSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GameStatefulSetList{ListMeta: obj.(*v1alpha1.GameStatefulSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.GameStatefulSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gameStatefulSets.
func (c *FakeGameStatefulSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gamestatefulsetsResource, c.ns, opts))

}

// Create takes the representation of a gameStatefulSet and creates it.  Returns the server's representation of the gameStatefulSet, and an error, if there is any.
func (c *FakeGameStatefulSets) Create(gameStatefulSet *v1alpha1.GameStatefulSet) (result *v1alpha1.GameStatefulSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gamestatefulsetsResource, c.ns, gameStatefulSet), &v1alpha1.GameStatefulSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameStatefulSet), err
}

// Update takes the representation of a gameStatefulSet and updates it. Returns the server's representation of the gameStatefulSet, and an error, if there is any.
func (c *FakeGameStatefulSets) Update(gameStatefulSet *v1alpha1.GameStatefulSet) (result *v1alpha1.GameStatefulSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gamestatefulsetsResource, c.ns, gameStatefulSet), &v1alpha1.GameStatefulSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameStatefulSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGameStatefulSets) UpdateStatus(gameStatefulSet *v1alpha1.GameStatefulSet) (*v1alpha1.GameStatefulSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gamestatefulsetsResource, "status", c.ns, gameStatefulSet), &v1alpha1.GameStatefulSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameStatefulSet), err
}

// Delete takes name of the gameStatefulSet and deletes it. Returns an error if one occurs.
func (c *FakeGameStatefulSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gamestatefulsetsResource, c.ns, name), &v1alpha1.GameStatefulSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGameStatefulSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gamestatefulsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GameStatefulSetList{})
	return err
}

// Patch applies the patch and returns the patched gameStatefulSet.
func (c *FakeGameStatefulSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GameStatefulSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gamestatefulsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GameStatefulSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameStatefulSet), err
}

// GetScale takes name of the gameStatefulSet, and returns the corresponding scale object, and an error if there is any.
func (c *FakeGameStatefulSets) GetScale(gameStatefulSetName string, options v1.GetOptions) (result *autoscaling.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceAction(gamestatefulsetsResource, c.ns, "scale", gameStatefulSetName), &autoscaling.Scale{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscaling.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeGameStatefulSets) UpdateScale(gameStatefulSetName string, scale *autoscaling.Scale) (result *autoscaling.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gamestatefulsetsResource, "scale", c.ns, scale), &autoscaling.Scale{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscaling.Scale), err
}
