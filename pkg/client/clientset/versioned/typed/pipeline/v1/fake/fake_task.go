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

// FakeTasks implements TaskInterface
type FakeTasks struct {
	Fake *FakeTektonV1
	ns   string
}

var tasksResource = schema.GroupVersionResource{Group: "tekton.dev", Version: "v1", Resource: "tasks"}

var tasksKind = schema.GroupVersionKind{Group: "tekton.dev", Version: "v1", Kind: "Task"}

// Get takes name of the task, and returns the corresponding task object, and an error if there is any.
func (c *FakeTasks) Get(ctx context.Context, name string, options v1.GetOptions) (result *pipelinev1.Task, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tasksResource, c.ns, name), &pipelinev1.Task{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pipelinev1.Task), err
}

// List takes label and field selectors, and returns the list of Tasks that match those selectors.
func (c *FakeTasks) List(ctx context.Context, opts v1.ListOptions) (result *pipelinev1.TaskList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tasksResource, tasksKind, c.ns, opts), &pipelinev1.TaskList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &pipelinev1.TaskList{ListMeta: obj.(*pipelinev1.TaskList).ListMeta}
	for _, item := range obj.(*pipelinev1.TaskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tasks.
func (c *FakeTasks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tasksResource, c.ns, opts))

}

// Create takes the representation of a task and creates it.  Returns the server's representation of the task, and an error, if there is any.
func (c *FakeTasks) Create(ctx context.Context, task *pipelinev1.Task, opts v1.CreateOptions) (result *pipelinev1.Task, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tasksResource, c.ns, task), &pipelinev1.Task{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pipelinev1.Task), err
}

// Update takes the representation of a task and updates it. Returns the server's representation of the task, and an error, if there is any.
func (c *FakeTasks) Update(ctx context.Context, task *pipelinev1.Task, opts v1.UpdateOptions) (result *pipelinev1.Task, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tasksResource, c.ns, task), &pipelinev1.Task{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pipelinev1.Task), err
}

// Delete takes name of the task and deletes it. Returns an error if one occurs.
func (c *FakeTasks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(tasksResource, c.ns, name, opts), &pipelinev1.Task{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTasks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tasksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &pipelinev1.TaskList{})
	return err
}

// Patch applies the patch and returns the patched task.
func (c *FakeTasks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *pipelinev1.Task, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tasksResource, c.ns, name, pt, data, subresources...), &pipelinev1.Task{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pipelinev1.Task), err
}
