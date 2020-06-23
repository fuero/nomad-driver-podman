// Code generated by go-swagger; DO NOT EDIT.

package images_libpod

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodPushImageReader is a Reader for the LibpodPushImage structure.
type LibpodPushImageReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *LibpodPushImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLibpodPushImageOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLibpodPushImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLibpodPushImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLibpodPushImageOK creates a LibpodPushImageOK with default headers values
func NewLibpodPushImageOK(writer io.Writer) *LibpodPushImageOK {
	return &LibpodPushImageOK{
		Payload: writer,
	}
}

/*LibpodPushImageOK handles this case with default header values.

no error
*/
type LibpodPushImageOK struct {
	Payload io.Writer
}

func (o *LibpodPushImageOK) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name:.*}/push][%d] libpodPushImageOK  %+v", 200, o.Payload)
}

func (o *LibpodPushImageOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *LibpodPushImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodPushImageNotFound creates a LibpodPushImageNotFound with default headers values
func NewLibpodPushImageNotFound() *LibpodPushImageNotFound {
	return &LibpodPushImageNotFound{}
}

/*LibpodPushImageNotFound handles this case with default header values.

No such image
*/
type LibpodPushImageNotFound struct {
	Payload *LibpodPushImageNotFoundBody
}

func (o *LibpodPushImageNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name:.*}/push][%d] libpodPushImageNotFound  %+v", 404, o.Payload)
}

func (o *LibpodPushImageNotFound) GetPayload() *LibpodPushImageNotFoundBody {
	return o.Payload
}

func (o *LibpodPushImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodPushImageNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLibpodPushImageInternalServerError creates a LibpodPushImageInternalServerError with default headers values
func NewLibpodPushImageInternalServerError() *LibpodPushImageInternalServerError {
	return &LibpodPushImageInternalServerError{}
}

/*LibpodPushImageInternalServerError handles this case with default header values.

Internal server error
*/
type LibpodPushImageInternalServerError struct {
	Payload *LibpodPushImageInternalServerErrorBody
}

func (o *LibpodPushImageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name:.*}/push][%d] libpodPushImageInternalServerError  %+v", 500, o.Payload)
}

func (o *LibpodPushImageInternalServerError) GetPayload() *LibpodPushImageInternalServerErrorBody {
	return o.Payload
}

func (o *LibpodPushImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(LibpodPushImageInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*LibpodPushImageInternalServerErrorBody libpod push image internal server error body
swagger:model LibpodPushImageInternalServerErrorBody
*/
type LibpodPushImageInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod push image internal server error body
func (o *LibpodPushImageInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodPushImageInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodPushImageInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res LibpodPushImageInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*LibpodPushImageNotFoundBody libpod push image not found body
swagger:model LibpodPushImageNotFoundBody
*/
type LibpodPushImageNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this libpod push image not found body
func (o *LibpodPushImageNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *LibpodPushImageNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LibpodPushImageNotFoundBody) UnmarshalBinary(b []byte) error {
	var res LibpodPushImageNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}