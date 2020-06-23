// Code generated by go-swagger; DO NOT EDIT.

package images_compat

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

// NewTagImageParams creates a new TagImageParams object
// with the default values initialized.
func NewTagImageParams() *TagImageParams {
	var ()
	return &TagImageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTagImageParamsWithTimeout creates a new TagImageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTagImageParamsWithTimeout(timeout time.Duration) *TagImageParams {
	var ()
	return &TagImageParams{

		timeout: timeout,
	}
}

// NewTagImageParamsWithContext creates a new TagImageParams object
// with the default values initialized, and the ability to set a context for a request
func NewTagImageParamsWithContext(ctx context.Context) *TagImageParams {
	var ()
	return &TagImageParams{

		Context: ctx,
	}
}

// NewTagImageParamsWithHTTPClient creates a new TagImageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTagImageParamsWithHTTPClient(client *http.Client) *TagImageParams {
	var ()
	return &TagImageParams{
		HTTPClient: client,
	}
}

/*TagImageParams contains all the parameters to send to the API endpoint
for the tag image operation typically these are written to a http.Request
*/
type TagImageParams struct {

	/*Name
	  the name or ID of the container

	*/
	Name string
	/*Repo
	  the repository to tag in

	*/
	Repo *string
	/*Tag
	  the name of the new tag

	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tag image params
func (o *TagImageParams) WithTimeout(timeout time.Duration) *TagImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tag image params
func (o *TagImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tag image params
func (o *TagImageParams) WithContext(ctx context.Context) *TagImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tag image params
func (o *TagImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tag image params
func (o *TagImageParams) WithHTTPClient(client *http.Client) *TagImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tag image params
func (o *TagImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the tag image params
func (o *TagImageParams) WithName(name string) *TagImageParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the tag image params
func (o *TagImageParams) SetName(name string) {
	o.Name = name
}

// WithRepo adds the repo to the tag image params
func (o *TagImageParams) WithRepo(repo *string) *TagImageParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the tag image params
func (o *TagImageParams) SetRepo(repo *string) {
	o.Repo = repo
}

// WithTag adds the tag to the tag image params
func (o *TagImageParams) WithTag(tag *string) *TagImageParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the tag image params
func (o *TagImageParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *TagImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name:.*
	if err := r.SetPathParam("name:.*", o.Name); err != nil {
		return err
	}

	if o.Repo != nil {

		// query param repo
		var qrRepo string
		if o.Repo != nil {
			qrRepo = *o.Repo
		}
		qRepo := qrRepo
		if qRepo != "" {
			if err := r.SetQueryParam("repo", qRepo); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}