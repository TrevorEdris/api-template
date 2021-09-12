package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/inconshreveable/log15"

	"github.com/TrevorEdris/api-template/app/util"
)

// Log adds the logger to the context of the request after adding a few additional contextual values.
func Log(log log15.Logger, h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := addLoggerToContext(r, log)
        h.ServeHTTP(w, r.WithContext(ctx))
    })
}

func addLoggerToContext(r *http.Request, log log15.Logger) context.Context {
    t := time.Now()
    requestDuration := func() float64 { return time.Since(t).Seconds() }
    ctx := log15.Ctx{
        "host": r.Host,
        "url": r.RequestURI,
        "method": r.Method,
        "request_duration": log15.Lazy{Fn: requestDuration},
    }
    l := log.New(ctx)
    return context.WithValue(r.Context(), util.LogContextKey, l)
}
