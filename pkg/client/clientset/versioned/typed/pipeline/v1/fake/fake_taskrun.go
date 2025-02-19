/*
Copyright 2023 The Tekton Authors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	pipelinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTaskRuns implements TaskRunInterface
type FakeTaskRuns struct {
	Fake *FakeTektonV1
	ns   string
}

var taskrunsResource = schema.GroupVersionResource{Group: "tekton.dev", Version: "v1", Resource: "taskruns"}

var taskrunsKind = schema.GroupVersionKind{Group: "tekton.dev", Version: "v1", Kind: "TaskRun"}

// Get takes name of the taskRun, and returns the corresponding taskRun object, and an error if there is any.
func (c *FakeTaskRuns) Get(ctx context.Context, name string, options v1.GetOptions) (result *pipelinev1.TaskRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(taskrunsResource, c.ns, name), &pipelinev1.TaskRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pipelinev1.TaskRun), err
}

// List takes label and field selectors, and returns the list of TaskRuns that match those selectors.
func (c *FakeTaskRuns) List(ctx context.Context, opts v1.ListOptions) (result *pipelinev1.TaskRunList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(taskrunsResource, taskrunsKind, c.ns, opts), &pipelinev1.TaskRunList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &pipelinev1.TaskRunList{ListMeta: obj.(*pipelinev1.TaskRunList).ListMeta}
	for _, item := range obj.(*pipelinev1.TaskRunList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested taskRuns.
func (c *FakeTaskRuns) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(taskrunsResource, c.ns, opts))

}

// Create takes the representation of a taskRun and creates it.  Returns the server's representation of the taskRun, and an error, if there is any.
func (c *FakeTaskRuns) Create(ctx context.Context, taskRun *pipelinev1.TaskRun, opts v1.CreateOptions) (result *pipelinev1.TaskRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(taskrunsResource, c.ns, taskRun), &pipelinev1.TaskRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pipelinev1.TaskRun), err
}

// Update takes the representation of a taskRun and updates it. Returns the server's representation of the taskRun, and an error, if there is any.
func (c *FakeTaskRuns) Update(ctx context.Context, taskRun *pipelinev1.TaskRun, opts v1.UpdateOptions) (result *pipelinev1.TaskRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(taskrunsResource, c.ns, taskRun), &pipelinev1.TaskRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pipelinev1.TaskRun), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTaskRuns) UpdateStatus(ctx context.Context, taskRun *pipelinev1.TaskRun, opts v1.UpdateOptions) (*pipelinev1.TaskRun, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(taskrunsResource, "status", c.ns, taskRun), &pipelinev1.TaskRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pipelinev1.TaskRun), err
}

// Delete takes name of the taskRun and deletes it. Returns an error if one occurs.
func (c *FakeTaskRuns) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(taskrunsResource, c.ns, name, opts), &pipelinev1.TaskRun{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTaskRuns) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(taskrunsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &pipelinev1.TaskRunList{})
	return err
}

// Patch applies the patch and returns the patched taskRun.
func (c *FakeTaskRuns) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *pipelinev1.TaskRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(taskrunsResource, c.ns, name, pt, data, subresources...), &pipelinev1.TaskRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pipelinev1.TaskRun), err
}
