package middleware

import (
	"net/http"

	"github.com/inconshreveable/log15"
)

func Auth(jwtIssuer string, log *log15.Logger, h http.Handler) http.Handler {
    // TODO: Implement some auth
    return h
}
