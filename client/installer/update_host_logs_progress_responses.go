// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// UpdateHostLogsProgressReader is a Reader for the UpdateHostLogsProgress structure.
type UpdateHostLogsProgressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateHostLogsProgressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateHostLogsProgressNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateHostLogsProgressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateHostLogsProgressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateHostLogsProgressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewUpdateHostLogsProgressMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateHostLogsProgressConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateHostLogsProgressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUpdateHostLogsProgressServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateHostLogsProgressNoContent creates a UpdateHostLogsProgressNoContent with default headers values
func NewUpdateHostLogsProgressNoContent() *UpdateHostLogsProgressNoContent {
	return &UpdateHostLogsProgressNoContent{}
}

/*UpdateHostLogsProgressNoContent handles this case with default header values.

Update cluster install progress.
*/
type UpdateHostLogsProgressNoContent struct {
}

func (o *UpdateHostLogsProgressNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/clusters/{cluster_id}/hosts/{host_id}/logs_progress][%d] updateHostLogsProgressNoContent ", 204)
}

func (o *UpdateHostLogsProgressNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateHostLogsProgressUnauthorized creates a UpdateHostLogsProgressUnauthorized with default headers values
func NewUpdateHostLogsProgressUnauthorized() *UpdateHostLogsProgressUnauthorized {
	return &UpdateHostLogsProgressUnauthorized{}
}

/*UpdateHostLogsProgressUnauthorized handles this case with default header values.

Unauthorized.
*/
type UpdateHostLogsProgressUnauthorized struct {
	Payload *models.InfraError
}

func (o *UpdateHostLogsProgressUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/clusters/{cluster_id}/hosts/{host_id}/logs_progress][%d] updateHostLogsProgressUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateHostLogsProgressUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UpdateHostLogsProgressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHostLogsProgressForbidden creates a UpdateHostLogsProgressForbidden with default headers values
func NewUpdateHostLogsProgressForbidden() *UpdateHostLogsProgressForbidden {
	return &UpdateHostLogsProgressForbidden{}
}

/*UpdateHostLogsProgressForbidden handles this case with default header values.

Forbidden.
*/
type UpdateHostLogsProgressForbidden struct {
	Payload *models.InfraError
}

func (o *UpdateHostLogsProgressForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/clusters/{cluster_id}/hosts/{host_id}/logs_progress][%d] updateHostLogsProgressForbidden  %+v", 403, o.Payload)
}

func (o *UpdateHostLogsProgressForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UpdateHostLogsProgressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHostLogsProgressNotFound creates a UpdateHostLogsProgressNotFound with default headers values
func NewUpdateHostLogsProgressNotFound() *UpdateHostLogsProgressNotFound {
	return &UpdateHostLogsProgressNotFound{}
}

/*UpdateHostLogsProgressNotFound handles this case with default header values.

Error.
*/
type UpdateHostLogsProgressNotFound struct {
	Payload *models.Error
}

func (o *UpdateHostLogsProgressNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/clusters/{cluster_id}/hosts/{host_id}/logs_progress][%d] updateHostLogsProgressNotFound  %+v", 404, o.Payload)
}

func (o *UpdateHostLogsProgressNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHostLogsProgressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHostLogsProgressMethodNotAllowed creates a UpdateHostLogsProgressMethodNotAllowed with default headers values
func NewUpdateHostLogsProgressMethodNotAllowed() *UpdateHostLogsProgressMethodNotAllowed {
	return &UpdateHostLogsProgressMethodNotAllowed{}
}

/*UpdateHostLogsProgressMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type UpdateHostLogsProgressMethodNotAllowed struct {
	Payload *models.Error
}

func (o *UpdateHostLogsProgressMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /v1/clusters/{cluster_id}/hosts/{host_id}/logs_progress][%d] updateHostLogsProgressMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *UpdateHostLogsProgressMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHostLogsProgressMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHostLogsProgressConflict creates a UpdateHostLogsProgressConflict with default headers values
func NewUpdateHostLogsProgressConflict() *UpdateHostLogsProgressConflict {
	return &UpdateHostLogsProgressConflict{}
}

/*UpdateHostLogsProgressConflict handles this case with default header values.

Error.
*/
type UpdateHostLogsProgressConflict struct {
	Payload *models.Error
}

func (o *UpdateHostLogsProgressConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/clusters/{cluster_id}/hosts/{host_id}/logs_progress][%d] updateHostLogsProgressConflict  %+v", 409, o.Payload)
}

func (o *UpdateHostLogsProgressConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHostLogsProgressConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHostLogsProgressInternalServerError creates a UpdateHostLogsProgressInternalServerError with default headers values
func NewUpdateHostLogsProgressInternalServerError() *UpdateHostLogsProgressInternalServerError {
	return &UpdateHostLogsProgressInternalServerError{}
}

/*UpdateHostLogsProgressInternalServerError handles this case with default header values.

Error.
*/
type UpdateHostLogsProgressInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateHostLogsProgressInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/clusters/{cluster_id}/hosts/{host_id}/logs_progress][%d] updateHostLogsProgressInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateHostLogsProgressInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHostLogsProgressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHostLogsProgressServiceUnavailable creates a UpdateHostLogsProgressServiceUnavailable with default headers values
func NewUpdateHostLogsProgressServiceUnavailable() *UpdateHostLogsProgressServiceUnavailable {
	return &UpdateHostLogsProgressServiceUnavailable{}
}

/*UpdateHostLogsProgressServiceUnavailable handles this case with default header values.

Unavailable.
*/
type UpdateHostLogsProgressServiceUnavailable struct {
	Payload *models.Error
}

func (o *UpdateHostLogsProgressServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /v1/clusters/{cluster_id}/hosts/{host_id}/logs_progress][%d] updateHostLogsProgressServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UpdateHostLogsProgressServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHostLogsProgressServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
