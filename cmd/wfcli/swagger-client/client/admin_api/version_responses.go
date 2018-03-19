// Code generated by go-swagger; DO NOT EDIT.

package admin_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fission/fission-workflows/cmd/wfcli/swagger-client/models"
)

// VersionReader is a Reader for the Version structure.
type VersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVersionOK creates a VersionOK with default headers values
func NewVersionOK() *VersionOK {
	return &VersionOK{}
}

/*VersionOK handles this case with default header values.

VersionOK version o k
*/
type VersionOK struct {
	Payload *models.ApiserverVersionResp
}

func (o *VersionOK) Error() string {
	return fmt.Sprintf("[GET /version][%d] versionOK  %+v", 200, o.Payload)
}

func (o *VersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApiserverVersionResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}