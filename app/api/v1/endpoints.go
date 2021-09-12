package v1

import (
	"encoding/json"
	"net/http"

	"github.com/TrevorEdris/api-template/app/util"
)

// V1 is a container around the v1 endpoints
type V1 struct {
    logProvider util.LogProvider
}

type simpleResponse struct {
    Msg string `json:"msg"`
}

type errorResponse struct {
    Status string `json:"status"`
    Error string `json:"error"`
}

// New returns an instance of the V1 container.
func New(logProvider util.LogProvider) *V1 {
    return &V1{logProvider}
}

// Health verifies the connectivity to any external services,
// responding with a 200 OK if everything is working as expected.
func (api *V1) Health(w http.ResponseWriter, r *http.Request) {
    // Check DB connectivity, ping external services, etc.
    // None to do for this API...
    response := simpleResponse{
        Msg: "Healthy",
    }
    api.logProvider(r).Info("Responding with 'Healthy'")
    api.response(w, r, http.StatusOK, response)
}

// GeneralKenobi is the enemy of General Grevious.
func (api *V1) GeneralKenobi(w http.ResponseWriter, r *http.Request) {
    response := simpleResponse{
        Msg: "Hello there.",
    }
    api.logProvider(r).Info("Responding with 'Hello there.'")
    api.response(w, r, http.StatusOK, response)
}

// response provides a consistent mechanism for the API to return a JSON response.
func (api *V1) response(w http.ResponseWriter, r *http.Request, status int, response interface{}) {
    log := api.logProvider(r)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Error("Unable to encode response", "error", err)
    }
}
