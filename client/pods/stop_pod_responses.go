// Code generated by go-swagger; DO NOT EDIT.

package pods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/pascomnet/nomad-driver-podman/models"
)

// StopPodReader is a Reader for the StopPod structure.
type StopPodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopPodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopPodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewStopPodNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewStopPodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopPodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopPodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopPodOK creates a StopPodOK with default headers values
func NewStopPodOK() *StopPodOK {
	return &StopPodOK{}
}

/*StopPodOK handles this case with default header values.

Stop pod
*/
type StopPodOK struct {
	Payload *models.PodStopReport
}

func (o *StopPodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] stopPodOK  %+v", 200, o.Payload)
}

func (o *StopPodOK) GetPayload() *models.PodStopReport {
	return o.Payload
}

func (o *StopPodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodStopReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopPodNotModified creates a StopPodNotModified with default headers values
func NewStopPodNotModified() *StopPodNotModified {
	return &StopPodNotModified{}
}

/*StopPodNotModified handles this case with default header values.

Pod already stopped
*/
type StopPodNotModified struct {
	Payload *StopPodNotModifiedBody
}

func (o *StopPodNotModified) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] stopPodNotModified  %+v", 304, o.Payload)
}

func (o *StopPodNotModified) GetPayload() *StopPodNotModifiedBody {
	return o.Payload
}

func (o *StopPodNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StopPodNotModifiedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopPodBadRequest creates a StopPodBadRequest with default headers values
func NewStopPodBadRequest() *StopPodBadRequest {
	return &StopPodBadRequest{}
}

/*StopPodBadRequest handles this case with default header values.

Bad parameter in request
*/
type StopPodBadRequest struct {
	Payload *StopPodBadRequestBody
}

func (o *StopPodBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] stopPodBadRequest  %+v", 400, o.Payload)
}

func (o *StopPodBadRequest) GetPayload() *StopPodBadRequestBody {
	return o.Payload
}

func (o *StopPodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StopPodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopPodNotFound creates a StopPodNotFound with default headers values
func NewStopPodNotFound() *StopPodNotFound {
	return &StopPodNotFound{}
}

/*StopPodNotFound handles this case with default header values.

No such pod
*/
type StopPodNotFound struct {
	Payload *StopPodNotFoundBody
}

func (o *StopPodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] stopPodNotFound  %+v", 404, o.Payload)
}

func (o *StopPodNotFound) GetPayload() *StopPodNotFoundBody {
	return o.Payload
}

func (o *StopPodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StopPodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopPodInternalServerError creates a StopPodInternalServerError with default headers values
func NewStopPodInternalServerError() *StopPodInternalServerError {
	return &StopPodInternalServerError{}
}

/*StopPodInternalServerError handles this case with default header values.

Internal server error
*/
type StopPodInternalServerError struct {
	Payload *StopPodInternalServerErrorBody
}

func (o *StopPodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] stopPodInternalServerError  %+v", 500, o.Payload)
}

func (o *StopPodInternalServerError) GetPayload() *StopPodInternalServerErrorBody {
	return o.Payload
}

func (o *StopPodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StopPodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StopPodBadRequestBody stop pod bad request body
swagger:model StopPodBadRequestBody
*/
type StopPodBadRequestBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this stop pod bad request body
func (o *StopPodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StopPodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StopPodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res StopPodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StopPodInternalServerErrorBody stop pod internal server error body
swagger:model StopPodInternalServerErrorBody
*/
type StopPodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this stop pod internal server error body
func (o *StopPodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StopPodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StopPodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res StopPodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StopPodNotFoundBody stop pod not found body
swagger:model StopPodNotFoundBody
*/
type StopPodNotFoundBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this stop pod not found body
func (o *StopPodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StopPodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StopPodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res StopPodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StopPodNotModifiedBody stop pod not modified body
swagger:model StopPodNotModifiedBody
*/
type StopPodNotModifiedBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this stop pod not modified body
func (o *StopPodNotModifiedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StopPodNotModifiedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StopPodNotModifiedBody) UnmarshalBinary(b []byte) error {
	var res StopPodNotModifiedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}