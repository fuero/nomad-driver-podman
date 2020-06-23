// Code generated by go-swagger; DO NOT EDIT.

package pods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewStartPodParams creates a new StartPodParams object
// with the default values initialized.
func NewStartPodParams() *StartPodParams {
	var ()
	return &StartPodParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartPodParamsWithTimeout creates a new StartPodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartPodParamsWithTimeout(timeout time.Duration) *StartPodParams {
	var ()
	return &StartPodParams{

		timeout: timeout,
	}
}

// NewStartPodParamsWithContext creates a new StartPodParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartPodParamsWithContext(ctx context.Context) *StartPodParams {
	var ()
	return &StartPodParams{

		Context: ctx,
	}
}

// NewStartPodParamsWithHTTPClient creates a new StartPodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartPodParamsWithHTTPClient(client *http.Client) *StartPodParams {
	var ()
	return &StartPodParams{
		HTTPClient: client,
	}
}

/*StartPodParams contains all the parameters to send to the API endpoint
for the start pod operation typically these are written to a http.Request
*/
type StartPodParams struct {

	/*Name
	  the name or ID of the pod

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start pod params
func (o *StartPodParams) WithTimeout(timeout time.Duration) *StartPodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start pod params
func (o *StartPodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start pod params
func (o *StartPodParams) WithContext(ctx context.Context) *StartPodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start pod params
func (o *StartPodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start pod params
func (o *StartPodParams) WithHTTPClient(client *http.Client) *StartPodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start pod params
func (o *StartPodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the start pod params
func (o *StartPodParams) WithName(name string) *StartPodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the start pod params
func (o *StartPodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *StartPodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}