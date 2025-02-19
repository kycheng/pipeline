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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	scheme "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VerificationPoliciesGetter has a method to return a VerificationPolicyInterface.
// A group's client should implement this interface.
type VerificationPoliciesGetter interface {
	VerificationPolicies(namespace string) VerificationPolicyInterface
}

// VerificationPolicyInterface has methods to work with VerificationPolicy resources.
type VerificationPolicyInterface interface {
	Create(ctx context.Context, verificationPolicy *v1alpha1.VerificationPolicy, opts v1.CreateOptions) (*v1alpha1.VerificationPolicy, error)
	Update(ctx context.Context, verificationPolicy *v1alpha1.VerificationPolicy, opts v1.UpdateOptions) (*v1alpha1.VerificationPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VerificationPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VerificationPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VerificationPolicy, err error)
	VerificationPolicyExpansion
}

// verificationPolicies implements VerificationPolicyInterface
type verificationPolicies struct {
	client rest.Interface
	ns     string
}

// newVerificationPolicies returns a VerificationPolicies
func newVerificationPolicies(c *TektonV1alpha1Client, namespace string) *verificationPolicies {
	return &verificationPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the verificationPolicy, and returns the corresponding verificationPolicy object, and an error if there is any.
func (c *verificationPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VerificationPolicy, err error) {
	result = &v1alpha1.VerificationPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("verificationpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VerificationPolicies that match those selectors.
func (c *verificationPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VerificationPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VerificationPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("verificationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested verificationPolicies.
func (c *verificationPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("verificationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a verificationPolicy and creates it.  Returns the server's representation of the verificationPolicy, and an error, if there is any.
func (c *verificationPolicies) Create(ctx context.Context, verificationPolicy *v1alpha1.VerificationPolicy, opts v1.CreateOptions) (result *v1alpha1.VerificationPolicy, err error) {
	result = &v1alpha1.VerificationPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("verificationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(verificationPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a verificationPolicy and updates it. Returns the server's representation of the verificationPolicy, and an error, if there is any.
func (c *verificationPolicies) Update(ctx context.Context, verificationPolicy *v1alpha1.VerificationPolicy, opts v1.UpdateOptions) (result *v1alpha1.VerificationPolicy, err error) {
	result = &v1alpha1.VerificationPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("verificationpolicies").
		Name(verificationPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(verificationPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the verificationPolicy and deletes it. Returns an error if one occurs.
func (c *verificationPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("verificationpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *verificationPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("verificationpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched verificationPolicy.
func (c *verificationPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VerificationPolicy, err error) {
	result = &v1alpha1.VerificationPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("verificationpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
