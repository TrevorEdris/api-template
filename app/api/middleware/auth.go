package middleware

import (
	"net/http"

	"github.com/inconshreveable/log15"
)

// Auth inserts JWT Authorization middleware into the lifecycle of an HTTP request.
func Auth(jwtIssuer string, log *log15.Logger, h http.Handler) http.Handler {
	// TODO: Implement some auth
	return h
}
