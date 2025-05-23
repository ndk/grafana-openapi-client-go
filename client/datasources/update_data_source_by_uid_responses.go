// Code generated by go-swagger; DO NOT EDIT.

package datasources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// UpdateDataSourceByUIDReader is a Reader for the UpdateDataSourceByUID structure.
type UpdateDataSourceByUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDataSourceByUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDataSourceByUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateDataSourceByUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDataSourceByUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateDataSourceByUIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateDataSourceByUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /datasources/uid/{uid}] updateDataSourceByUID", response, response.Code())
	}
}

// NewUpdateDataSourceByUIDOK creates a UpdateDataSourceByUIDOK with default headers values
func NewUpdateDataSourceByUIDOK() *UpdateDataSourceByUIDOK {
	return &UpdateDataSourceByUIDOK{}
}

/*
UpdateDataSourceByUIDOK describes a response with status code 200, with default header values.

(empty)
*/
type UpdateDataSourceByUIDOK struct {
	Payload *models.UpdateDataSourceByUIDOKBody
}

// IsSuccess returns true when this update data source by Uid Ok response has a 2xx status code
func (o *UpdateDataSourceByUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update data source by Uid Ok response has a 3xx status code
func (o *UpdateDataSourceByUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update data source by Uid Ok response has a 4xx status code
func (o *UpdateDataSourceByUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update data source by Uid Ok response has a 5xx status code
func (o *UpdateDataSourceByUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update data source by Uid Ok response a status code equal to that given
func (o *UpdateDataSourceByUIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update data source by Uid Ok response
func (o *UpdateDataSourceByUIDOK) Code() int {
	return 200
}

func (o *UpdateDataSourceByUIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidOk %s", 200, payload)
}

func (o *UpdateDataSourceByUIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidOk %s", 200, payload)
}

func (o *UpdateDataSourceByUIDOK) GetPayload() *models.UpdateDataSourceByUIDOKBody {
	return o.Payload
}

func (o *UpdateDataSourceByUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDataSourceByUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDataSourceByUIDUnauthorized creates a UpdateDataSourceByUIDUnauthorized with default headers values
func NewUpdateDataSourceByUIDUnauthorized() *UpdateDataSourceByUIDUnauthorized {
	return &UpdateDataSourceByUIDUnauthorized{}
}

/*
UpdateDataSourceByUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateDataSourceByUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update data source by Uid unauthorized response has a 2xx status code
func (o *UpdateDataSourceByUIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update data source by Uid unauthorized response has a 3xx status code
func (o *UpdateDataSourceByUIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update data source by Uid unauthorized response has a 4xx status code
func (o *UpdateDataSourceByUIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update data source by Uid unauthorized response has a 5xx status code
func (o *UpdateDataSourceByUIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update data source by Uid unauthorized response a status code equal to that given
func (o *UpdateDataSourceByUIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update data source by Uid unauthorized response
func (o *UpdateDataSourceByUIDUnauthorized) Code() int {
	return 401
}

func (o *UpdateDataSourceByUIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidUnauthorized %s", 401, payload)
}

func (o *UpdateDataSourceByUIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidUnauthorized %s", 401, payload)
}

func (o *UpdateDataSourceByUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateDataSourceByUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDataSourceByUIDForbidden creates a UpdateDataSourceByUIDForbidden with default headers values
func NewUpdateDataSourceByUIDForbidden() *UpdateDataSourceByUIDForbidden {
	return &UpdateDataSourceByUIDForbidden{}
}

/*
UpdateDataSourceByUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateDataSourceByUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update data source by Uid forbidden response has a 2xx status code
func (o *UpdateDataSourceByUIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update data source by Uid forbidden response has a 3xx status code
func (o *UpdateDataSourceByUIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update data source by Uid forbidden response has a 4xx status code
func (o *UpdateDataSourceByUIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update data source by Uid forbidden response has a 5xx status code
func (o *UpdateDataSourceByUIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update data source by Uid forbidden response a status code equal to that given
func (o *UpdateDataSourceByUIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update data source by Uid forbidden response
func (o *UpdateDataSourceByUIDForbidden) Code() int {
	return 403
}

func (o *UpdateDataSourceByUIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidForbidden %s", 403, payload)
}

func (o *UpdateDataSourceByUIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidForbidden %s", 403, payload)
}

func (o *UpdateDataSourceByUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateDataSourceByUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDataSourceByUIDConflict creates a UpdateDataSourceByUIDConflict with default headers values
func NewUpdateDataSourceByUIDConflict() *UpdateDataSourceByUIDConflict {
	return &UpdateDataSourceByUIDConflict{}
}

/*
UpdateDataSourceByUIDConflict describes a response with status code 409, with default header values.

ConflictError
*/
type UpdateDataSourceByUIDConflict struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update data source by Uid conflict response has a 2xx status code
func (o *UpdateDataSourceByUIDConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update data source by Uid conflict response has a 3xx status code
func (o *UpdateDataSourceByUIDConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update data source by Uid conflict response has a 4xx status code
func (o *UpdateDataSourceByUIDConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update data source by Uid conflict response has a 5xx status code
func (o *UpdateDataSourceByUIDConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update data source by Uid conflict response a status code equal to that given
func (o *UpdateDataSourceByUIDConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update data source by Uid conflict response
func (o *UpdateDataSourceByUIDConflict) Code() int {
	return 409
}

func (o *UpdateDataSourceByUIDConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidConflict %s", 409, payload)
}

func (o *UpdateDataSourceByUIDConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidConflict %s", 409, payload)
}

func (o *UpdateDataSourceByUIDConflict) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateDataSourceByUIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDataSourceByUIDInternalServerError creates a UpdateDataSourceByUIDInternalServerError with default headers values
func NewUpdateDataSourceByUIDInternalServerError() *UpdateDataSourceByUIDInternalServerError {
	return &UpdateDataSourceByUIDInternalServerError{}
}

/*
UpdateDataSourceByUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateDataSourceByUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update data source by Uid internal server error response has a 2xx status code
func (o *UpdateDataSourceByUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update data source by Uid internal server error response has a 3xx status code
func (o *UpdateDataSourceByUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update data source by Uid internal server error response has a 4xx status code
func (o *UpdateDataSourceByUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update data source by Uid internal server error response has a 5xx status code
func (o *UpdateDataSourceByUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update data source by Uid internal server error response a status code equal to that given
func (o *UpdateDataSourceByUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update data source by Uid internal server error response
func (o *UpdateDataSourceByUIDInternalServerError) Code() int {
	return 500
}

func (o *UpdateDataSourceByUIDInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidInternalServerError %s", 500, payload)
}

func (o *UpdateDataSourceByUIDInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidInternalServerError %s", 500, payload)
}

func (o *UpdateDataSourceByUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateDataSourceByUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
