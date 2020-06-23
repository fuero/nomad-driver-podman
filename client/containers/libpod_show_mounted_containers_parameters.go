// Code generated by go-swagger; DO NOT EDIT.

package containers

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

// NewLibpodShowMountedContainersParams creates a new LibpodShowMountedContainersParams object
// with the default values initialized.
func NewLibpodShowMountedContainersParams() *LibpodShowMountedContainersParams {

	return &LibpodShowMountedContainersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLibpodShowMountedContainersParamsWithTimeout creates a new LibpodShowMountedContainersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLibpodShowMountedContainersParamsWithTimeout(timeout time.Duration) *LibpodShowMountedContainersParams {

	return &LibpodShowMountedContainersParams{

		timeout: timeout,
	}
}

// NewLibpodShowMountedContainersParamsWithContext creates a new LibpodShowMountedContainersParams object
// with the default values initialized, and the ability to set a context for a request
func NewLibpodShowMountedContainersParamsWithContext(ctx context.Context) *LibpodShowMountedContainersParams {

	return &LibpodShowMountedContainersParams{

		Context: ctx,
	}
}

// NewLibpodShowMountedContainersParamsWithHTTPClient creates a new LibpodShowMountedContainersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLibpodShowMountedContainersParamsWithHTTPClient(client *http.Client) *LibpodShowMountedContainersParams {

	return &LibpodShowMountedContainersParams{
		HTTPClient: client,
	}
}

/*LibpodShowMountedContainersParams contains all the parameters to send to the API endpoint
for the libpod show mounted containers operation typically these are written to a http.Request
*/
type LibpodShowMountedContainersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the libpod show mounted containers params
func (o *LibpodShowMountedContainersParams) WithTimeout(timeout time.Duration) *LibpodShowMountedContainersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the libpod show mounted containers params
func (o *LibpodShowMountedContainersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the libpod show mounted containers params
func (o *LibpodShowMountedContainersParams) WithContext(ctx context.Context) *LibpodShowMountedContainersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the libpod show mounted containers params
func (o *LibpodShowMountedContainersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the libpod show mounted containers params
func (o *LibpodShowMountedContainersParams) WithHTTPClient(client *http.Client) *LibpodShowMountedContainersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the libpod show mounted containers params
func (o *LibpodShowMountedContainersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LibpodShowMountedContainersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}