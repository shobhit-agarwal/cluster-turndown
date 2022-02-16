/* Generated Source: Do Not Modify */
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kubecost/cluster-turndown/pkg/apis/turndownschedule/v1alpha1"
	scheme "github.com/kubecost/cluster-turndown/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TurndownSchedulesGetter has a method to return a TurndownScheduleInterface.
// A group's client should implement this interface.
type TurndownSchedulesGetter interface {
	TurndownSchedules() TurndownScheduleInterface
}

// TurndownScheduleInterface has methods to work with TurndownSchedule resources.
type TurndownScheduleInterface interface {
	Create(ctx context.Context, turndownSchedule *v1alpha1.TurndownSchedule, opts v1.CreateOptions) (*v1alpha1.TurndownSchedule, error)
	Update(ctx context.Context, turndownSchedule *v1alpha1.TurndownSchedule, opts v1.UpdateOptions) (*v1alpha1.TurndownSchedule, error)
	UpdateStatus(ctx context.Context, turndownSchedule *v1alpha1.TurndownSchedule, opts v1.UpdateOptions) (*v1alpha1.TurndownSchedule, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.TurndownSchedule, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.TurndownScheduleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TurndownSchedule, err error)
	TurndownScheduleExpansion
}

// turndownSchedules implements TurndownScheduleInterface
type turndownSchedules struct {
	client rest.Interface
}

// newTurndownSchedules returns a TurndownSchedules
func newTurndownSchedules(c *KubecostV1alpha1Client) *turndownSchedules {
	return &turndownSchedules{
		client: c.RESTClient(),
	}
}

// Get takes name of the turndownSchedule, and returns the corresponding turndownSchedule object, and an error if there is any.
func (c *turndownSchedules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TurndownSchedule, err error) {
	result = &v1alpha1.TurndownSchedule{}
	err = c.client.Get().
		Resource("turndownschedules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TurndownSchedules that match those selectors.
func (c *turndownSchedules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TurndownScheduleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TurndownScheduleList{}
	err = c.client.Get().
		Resource("turndownschedules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested turndownSchedules.
func (c *turndownSchedules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("turndownschedules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a turndownSchedule and creates it.  Returns the server's representation of the turndownSchedule, and an error, if there is any.
func (c *turndownSchedules) Create(ctx context.Context, turndownSchedule *v1alpha1.TurndownSchedule, opts v1.CreateOptions) (result *v1alpha1.TurndownSchedule, err error) {
	result = &v1alpha1.TurndownSchedule{}
	err = c.client.Post().
		Resource("turndownschedules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(turndownSchedule).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a turndownSchedule and updates it. Returns the server's representation of the turndownSchedule, and an error, if there is any.
func (c *turndownSchedules) Update(ctx context.Context, turndownSchedule *v1alpha1.TurndownSchedule, opts v1.UpdateOptions) (result *v1alpha1.TurndownSchedule, err error) {
	result = &v1alpha1.TurndownSchedule{}
	err = c.client.Put().
		Resource("turndownschedules").
		Name(turndownSchedule.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(turndownSchedule).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *turndownSchedules) UpdateStatus(ctx context.Context, turndownSchedule *v1alpha1.TurndownSchedule, opts v1.UpdateOptions) (result *v1alpha1.TurndownSchedule, err error) {
	result = &v1alpha1.TurndownSchedule{}
	err = c.client.Put().
		Resource("turndownschedules").
		Name(turndownSchedule.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(turndownSchedule).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the turndownSchedule and deletes it. Returns an error if one occurs.
func (c *turndownSchedules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("turndownschedules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *turndownSchedules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("turndownschedules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched turndownSchedule.
func (c *turndownSchedules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TurndownSchedule, err error) {
	result = &v1alpha1.TurndownSchedule{}
	err = c.client.Patch(pt).
		Resource("turndownschedules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
