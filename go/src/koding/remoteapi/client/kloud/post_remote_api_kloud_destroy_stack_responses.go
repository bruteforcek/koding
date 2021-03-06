package kloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIKloudDestroyStackReader is a Reader for the PostRemoteAPIKloudDestroyStack structure.
type PostRemoteAPIKloudDestroyStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIKloudDestroyStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIKloudDestroyStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIKloudDestroyStackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIKloudDestroyStackOK creates a PostRemoteAPIKloudDestroyStackOK with default headers values
func NewPostRemoteAPIKloudDestroyStackOK() *PostRemoteAPIKloudDestroyStackOK {
	return &PostRemoteAPIKloudDestroyStackOK{}
}

/*PostRemoteAPIKloudDestroyStackOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIKloudDestroyStackOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIKloudDestroyStackOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/Kloud.destroyStack][%d] postRemoteApiKloudDestroyStackOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIKloudDestroyStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIKloudDestroyStackUnauthorized creates a PostRemoteAPIKloudDestroyStackUnauthorized with default headers values
func NewPostRemoteAPIKloudDestroyStackUnauthorized() *PostRemoteAPIKloudDestroyStackUnauthorized {
	return &PostRemoteAPIKloudDestroyStackUnauthorized{}
}

/*PostRemoteAPIKloudDestroyStackUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIKloudDestroyStackUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIKloudDestroyStackUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/Kloud.destroyStack][%d] postRemoteApiKloudDestroyStackUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIKloudDestroyStackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
