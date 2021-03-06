/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	v1 "kube-inferencer/pkg/apis/inferencer/v1"
	scheme "kube-inferencer/pkg/generate/clientset/versioned/scheme"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InferencersGetter has a method to return a InferencerInterface.
// A group's client should implement this interface.
type InferencersGetter interface {
	Inferencers(namespace string) InferencerInterface
}

// InferencerInterface has methods to work with Inferencer resources.
type InferencerInterface interface {
	Create(ctx context.Context, inferencer *v1.Inferencer, opts metav1.CreateOptions) (*v1.Inferencer, error)
	Update(ctx context.Context, inferencer *v1.Inferencer, opts metav1.UpdateOptions) (*v1.Inferencer, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Inferencer, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.InferencerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Inferencer, err error)
	InferencerExpansion
}

// inferencers implements InferencerInterface
type inferencers struct {
	client rest.Interface
	ns     string
}

// newInferencers returns a Inferencers
func newInferencers(c *InferencerV1Client, namespace string) *inferencers {
	return &inferencers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the inferencer, and returns the corresponding inferencer object, and an error if there is any.
func (c *inferencers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Inferencer, err error) {
	result = &v1.Inferencer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("inferencers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Inferencers that match those selectors.
func (c *inferencers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.InferencerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.InferencerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("inferencers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested inferencers.
func (c *inferencers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("inferencers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a inferencer and creates it.  Returns the server's representation of the inferencer, and an error, if there is any.
func (c *inferencers) Create(ctx context.Context, inferencer *v1.Inferencer, opts metav1.CreateOptions) (result *v1.Inferencer, err error) {
	result = &v1.Inferencer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("inferencers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(inferencer).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a inferencer and updates it. Returns the server's representation of the inferencer, and an error, if there is any.
func (c *inferencers) Update(ctx context.Context, inferencer *v1.Inferencer, opts metav1.UpdateOptions) (result *v1.Inferencer, err error) {
	result = &v1.Inferencer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("inferencers").
		Name(inferencer.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(inferencer).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the inferencer and deletes it. Returns an error if one occurs.
func (c *inferencers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("inferencers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *inferencers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("inferencers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched inferencer.
func (c *inferencers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Inferencer, err error) {
	result = &v1.Inferencer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("inferencers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
