package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	headerKeyContentType       = "Content-Type"
	headerValueApplicationJSON = "application/json"
	maxBodyBytes               = 1048576 // 1MiB
)

var (
	// ErrInvalidContentType indicates the request had an invalid Content-Type.
	ErrInvalidContentType = errors.New("invalid Content-Type")
)

// EchoRequest is the expected request body to the /echo endpoint.
type EchoRequest struct {
	Body struct {
		Msg string `json:"message"`
	}
	// Additional fields can go here if necessary,
	// such as metadata about the request.
}

// newEchoRequest validates the http request, parsing it into an EchoRequest.
func newEchoRequest(w http.ResponseWriter, r *http.Request) (*EchoRequest, error) {
	req := EchoRequest{}

	// Ensure the Content-Type: application/json header was provided
	err := validateMediaType(r, headerValueApplicationJSON)
	if err != nil {
		return &req, err
	}

	// Limit the size of the request body
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	// Don't allow unspecified fields in the request body
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	// Parse the body into the body of the EchoRequest
	err = dec.Decode(&req.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidJSON, err)
	}

	return &req, nil
}

// validateMediaType validates the Content-Type header was provided and matches expectedContentType.
func validateMediaType(r *http.Request, expectedContentType string) error {
	contentType := r.Header.Get(headerKeyContentType)
	if contentType != expectedContentType {
		return fmt.Errorf("%w: %s, expected %s", ErrInvalidContentType, contentType, expectedContentType)
	}
	return nil
}
