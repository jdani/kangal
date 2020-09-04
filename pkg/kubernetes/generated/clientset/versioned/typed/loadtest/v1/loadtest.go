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
	"time"

	v1 "github.com/hellofresh/kangal/pkg/kubernetes/apis/loadtest/v1"
	scheme "github.com/hellofresh/kangal/pkg/kubernetes/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LoadTestsGetter has a method to return a LoadTestInterface.
// A group's client should implement this interface.
type LoadTestsGetter interface {
	LoadTests() LoadTestInterface
}

// LoadTestInterface has methods to work with LoadTest resources.
type LoadTestInterface interface {
	Create(ctx context.Context, loadTest *v1.LoadTest, opts metav1.CreateOptions) (*v1.LoadTest, error)
	Update(ctx context.Context, loadTest *v1.LoadTest, opts metav1.UpdateOptions) (*v1.LoadTest, error)
	UpdateStatus(ctx context.Context, loadTest *v1.LoadTest, opts metav1.UpdateOptions) (*v1.LoadTest, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.LoadTest, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.LoadTestList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LoadTest, err error)
	LoadTestExpansion
}

// loadTests implements LoadTestInterface
type loadTests struct {
	client rest.Interface
}

// newLoadTests returns a LoadTests
func newLoadTests(c *KangalV1Client) *loadTests {
	return &loadTests{
		client: c.RESTClient(),
	}
}

// Get takes name of the loadTest, and returns the corresponding loadTest object, and an error if there is any.
func (c *loadTests) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.LoadTest, err error) {
	result = &v1.LoadTest{}
	err = c.client.Get().
		Resource("loadtests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LoadTests that match those selectors.
func (c *loadTests) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LoadTestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.LoadTestList{}
	err = c.client.Get().
		Resource("loadtests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested loadTests.
func (c *loadTests) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("loadtests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a loadTest and creates it.  Returns the server's representation of the loadTest, and an error, if there is any.
func (c *loadTests) Create(ctx context.Context, loadTest *v1.LoadTest, opts metav1.CreateOptions) (result *v1.LoadTest, err error) {
	result = &v1.LoadTest{}
	err = c.client.Post().
		Resource("loadtests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loadTest).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a loadTest and updates it. Returns the server's representation of the loadTest, and an error, if there is any.
func (c *loadTests) Update(ctx context.Context, loadTest *v1.LoadTest, opts metav1.UpdateOptions) (result *v1.LoadTest, err error) {
	result = &v1.LoadTest{}
	err = c.client.Put().
		Resource("loadtests").
		Name(loadTest.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loadTest).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *loadTests) UpdateStatus(ctx context.Context, loadTest *v1.LoadTest, opts metav1.UpdateOptions) (result *v1.LoadTest, err error) {
	result = &v1.LoadTest{}
	err = c.client.Put().
		Resource("loadtests").
		Name(loadTest.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loadTest).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the loadTest and deletes it. Returns an error if one occurs.
func (c *loadTests) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("loadtests").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *loadTests) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("loadtests").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched loadTest.
func (c *loadTests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LoadTest, err error) {
	result = &v1.LoadTest{}
	err = c.client.Patch(pt).
		Resource("loadtests").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
