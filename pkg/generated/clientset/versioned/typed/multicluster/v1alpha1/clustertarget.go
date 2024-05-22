/*
 * Copyright 2020 The Multicluster-Scheduler Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "admiralty.io/multicluster-scheduler/pkg/apis/multicluster/v1alpha1"
	scheme "admiralty.io/multicluster-scheduler/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterTargetsGetter has a method to return a ClusterTargetInterface.
// A group's client should implement this interface.
type ClusterTargetsGetter interface {
	ClusterTargets() ClusterTargetInterface
}

// ClusterTargetInterface has methods to work with ClusterTarget resources.
type ClusterTargetInterface interface {
	Create(ctx context.Context, clusterTarget *v1alpha1.ClusterTarget, opts v1.CreateOptions) (*v1alpha1.ClusterTarget, error)
	Update(ctx context.Context, clusterTarget *v1alpha1.ClusterTarget, opts v1.UpdateOptions) (*v1alpha1.ClusterTarget, error)
	UpdateStatus(ctx context.Context, clusterTarget *v1alpha1.ClusterTarget, opts v1.UpdateOptions) (*v1alpha1.ClusterTarget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterTarget, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterTargetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterTarget, err error)
	ClusterTargetExpansion
}

// clusterTargets implements ClusterTargetInterface
type clusterTargets struct {
	client rest.Interface
}

// newClusterTargets returns a ClusterTargets
func newClusterTargets(c *MulticlusterV1alpha1Client) *clusterTargets {
	return &clusterTargets{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterTarget, and returns the corresponding clusterTarget object, and an error if there is any.
func (c *clusterTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterTarget, err error) {
	result = &v1alpha1.ClusterTarget{}
	err = c.client.Get().
		Resource("clustertargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterTargets that match those selectors.
func (c *clusterTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterTargetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterTargetList{}
	err = c.client.Get().
		Resource("clustertargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterTargets.
func (c *clusterTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clustertargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterTarget and creates it.  Returns the server's representation of the clusterTarget, and an error, if there is any.
func (c *clusterTargets) Create(ctx context.Context, clusterTarget *v1alpha1.ClusterTarget, opts v1.CreateOptions) (result *v1alpha1.ClusterTarget, err error) {
	result = &v1alpha1.ClusterTarget{}
	err = c.client.Post().
		Resource("clustertargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTarget).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterTarget and updates it. Returns the server's representation of the clusterTarget, and an error, if there is any.
func (c *clusterTargets) Update(ctx context.Context, clusterTarget *v1alpha1.ClusterTarget, opts v1.UpdateOptions) (result *v1alpha1.ClusterTarget, err error) {
	result = &v1alpha1.ClusterTarget{}
	err = c.client.Put().
		Resource("clustertargets").
		Name(clusterTarget.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTarget).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterTargets) UpdateStatus(ctx context.Context, clusterTarget *v1alpha1.ClusterTarget, opts v1.UpdateOptions) (result *v1alpha1.ClusterTarget, err error) {
	result = &v1alpha1.ClusterTarget{}
	err = c.client.Put().
		Resource("clustertargets").
		Name(clusterTarget.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterTarget).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterTarget and deletes it. Returns an error if one occurs.
func (c *clusterTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustertargets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustertargets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterTarget.
func (c *clusterTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterTarget, err error) {
	result = &v1alpha1.ClusterTarget{}
	err = c.client.Patch(pt).
		Resource("clustertargets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
