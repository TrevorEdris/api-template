package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/TrevorEdris/api-template/app/models/item"
	"github.com/TrevorEdris/api-template/app/util"
)

var (
    // ErrInvalidJSON indicates that the request body contained invalid json.
    ErrInvalidJSON = errors.New("invalid JSON")
)

// V1 is a container around the v1 endpoints.
type V1 struct {
    logProvider util.LogProvider
    items *item.Items
}

// New returns an instance of the V1 container.
func New(logProvider util.LogProvider, items *item.Items) *V1 {
    return &V1{
        logProvider: logProvider,
        items: items,
    }
}

// Health verifies the connectivity to any external services,
// responding with a 200 OK if everything is working as expected.
func (api *V1) Health(w http.ResponseWriter, r *http.Request) {
    // Check DB connectivity, ping external services, etc.
    // None to do for this API...
    response := newSimpleResponse("Healthy")
    api.logProvider(r).Debug("Responding", "response", response)
    api.response(w, r, http.StatusOK, response)
}

// GeneralKenobi is the enemy of General Grevious.
func (api *V1) GeneralKenobi(w http.ResponseWriter, r *http.Request) {
    response := newSimpleResponse("Hello there.")
    api.logProvider(r).Debug("Responding", "response", response)
    api.response(w, r, http.StatusOK, response)
}

// Echo simply returns the request body if it matches the expected format.
func (api *V1) Echo(w http.ResponseWriter, r *http.Request) {
    log := api.logProvider(r)
    req, err := newEchoRequest(w, r)
    if err != nil {
        log.Warn("Unable to unmarshal request body into expected format", "error", err)
        api.response(w, r, http.StatusBadRequest, newErrorResponse(http.StatusText(http.StatusBadRequest), err.Error()))
        return
    }
    log.Info("Received", "request", req)
    resp := newSimpleResponse(req.Body.Msg)
    api.response(w, r, http.StatusOK, resp)
}

// response provides a consistent mechanism for the API to return a JSON response.
func (api *V1) response(w http.ResponseWriter, r *http.Request, status int, response interface{}) {
    log := api.logProvider(r)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Error("Unable to encode response", "error", err)
    }
    api.logProvider(r).Info("Request completed")
}
