// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/models"
)

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateUserCreated creates a CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {
	return &CreateUserCreated{}
}

/*CreateUserCreated handles this case with default header values.

User created
*/
type CreateUserCreated struct {
	Payload *models.V4GenericResponse
}

func (o *CreateUserCreated) Error() string {
	return fmt.Sprintf("[PUT /v4/users/{email}/][%d] createUserCreated  %+v", 201, o.Payload)
}

func (o *CreateUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserBadRequest creates a CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {
	return &CreateUserBadRequest{}
}

/*CreateUserBadRequest handles this case with default header values.

User already exists
*/
type CreateUserBadRequest struct {
	Payload *models.V4GenericResponse
}

func (o *CreateUserBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v4/users/{email}/][%d] createUserBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserUnauthorized creates a CreateUserUnauthorized with default headers values
func NewCreateUserUnauthorized() *CreateUserUnauthorized {
	return &CreateUserUnauthorized{}
}

/*CreateUserUnauthorized handles this case with default header values.

Permission denied
*/
type CreateUserUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *CreateUserUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v4/users/{email}/][%d] createUserUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserDefault creates a CreateUserDefault with default headers values
func NewCreateUserDefault(code int) *CreateUserDefault {
	return &CreateUserDefault{
		_statusCode: code,
	}
}

/*CreateUserDefault handles this case with default header values.

Error
*/
type CreateUserDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the create user default response
func (o *CreateUserDefault) Code() int {
	return o._statusCode
}

func (o *CreateUserDefault) Error() string {
	return fmt.Sprintf("[PUT /v4/users/{email}/][%d] createUser default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
