// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListContainersReader is a Reader for the ListContainers structure.
type ListContainersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListContainersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListContainersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListContainersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListContainersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListContainersOK creates a ListContainersOK with default headers values
func NewListContainersOK() *ListContainersOK {
	return &ListContainersOK{}
}

/*ListContainersOK handles this case with default header values.

List Containers
*/
type ListContainersOK struct {
	Payload interface{}
}

func (o *ListContainersOK) Error() string {
	return fmt.Sprintf("[GET /containers/json][%d] listContainersOK  %+v", 200, o.Payload)
}

func (o *ListContainersOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ListContainersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListContainersBadRequest creates a ListContainersBadRequest with default headers values
func NewListContainersBadRequest() *ListContainersBadRequest {
	return &ListContainersBadRequest{}
}

/*ListContainersBadRequest handles this case with default header values.

Bad parameter in request
*/
type ListContainersBadRequest struct {
	Payload *ListContainersBadRequestBody
}

func (o *ListContainersBadRequest) Error() string {
	return fmt.Sprintf("[GET /containers/json][%d] listContainersBadRequest  %+v", 400, o.Payload)
}

func (o *ListContainersBadRequest) GetPayload() *ListContainersBadRequestBody {
	return o.Payload
}

func (o *ListContainersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListContainersBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListContainersInternalServerError creates a ListContainersInternalServerError with default headers values
func NewListContainersInternalServerError() *ListContainersInternalServerError {
	return &ListContainersInternalServerError{}
}

/*ListContainersInternalServerError handles this case with default header values.

Internal server error
*/
type ListContainersInternalServerError struct {
	Payload *ListContainersInternalServerErrorBody
}

func (o *ListContainersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/json][%d] listContainersInternalServerError  %+v", 500, o.Payload)
}

func (o *ListContainersInternalServerError) GetPayload() *ListContainersInternalServerErrorBody {
	return o.Payload
}

func (o *ListContainersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListContainersInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListContainersBadRequestBody list containers bad request body
swagger:model ListContainersBadRequestBody
*/
type ListContainersBadRequestBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this list containers bad request body
func (o *ListContainersBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListContainersBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListContainersBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ListContainersBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListContainersInternalServerErrorBody list containers internal server error body
swagger:model ListContainersInternalServerErrorBody
*/
type ListContainersInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this list containers internal server error body
func (o *ListContainersInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListContainersInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListContainersInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ListContainersInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}