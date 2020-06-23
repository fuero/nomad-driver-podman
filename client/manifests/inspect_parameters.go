// Code generated by go-swagger; DO NOT EDIT.

package manifests

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

// NewInspectParams creates a new InspectParams object
// with the default values initialized.
func NewInspectParams() *InspectParams {
	var ()
	return &InspectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInspectParamsWithTimeout creates a new InspectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInspectParamsWithTimeout(timeout time.Duration) *InspectParams {
	var ()
	return &InspectParams{

		timeout: timeout,
	}
}

// NewInspectParamsWithContext creates a new InspectParams object
// with the default values initialized, and the ability to set a context for a request
func NewInspectParamsWithContext(ctx context.Context) *InspectParams {
	var ()
	return &InspectParams{

		Context: ctx,
	}
}

// NewInspectParamsWithHTTPClient creates a new InspectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInspectParamsWithHTTPClient(client *http.Client) *InspectParams {
	var ()
	return &InspectParams{
		HTTPClient: client,
	}
}

/*InspectParams contains all the parameters to send to the API endpoint
for the inspect operation typically these are written to a http.Request
*/
type InspectParams struct {

	/*Name
	  the name or ID of the manifest

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the inspect params
func (o *InspectParams) WithTimeout(timeout time.Duration) *InspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the inspect params
func (o *InspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the inspect params
func (o *InspectParams) WithContext(ctx context.Context) *InspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the inspect params
func (o *InspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the inspect params
func (o *InspectParams) WithHTTPClient(client *http.Client) *InspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the inspect params
func (o *InspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the inspect params
func (o *InspectParams) WithName(name string) *InspectParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the inspect params
func (o *InspectParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *InspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name:.*
	if err := r.SetPathParam("name:.*", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}