package util

import (
	"context"
	"fmt"
	"net/http"

	"github.com/inconshreveable/log15"
)

const (
	// LogContextKey is the key used to insert the log15.Logger into a given context.Context.
	LogContextKey = "log-context-key"
)

var (
	// Create a default logger to use in the case of an error
	defaultlog = log15.New(log15.Ctx{"module": "util/log"})
)

// LogProvider is a custom type, defining a function that takes an http.Request and returns a log15.Logger.
type LogProvider func(r *http.Request) log15.Logger

// GetLoggerFromContext attempts to pull a log15.Logger out of the given context.Context, returning a default log15.Logger
// if the context does not have the expected key/value.
func GetLoggerFromContext(ctx context.Context) log15.Logger {

	// Attempt to pull the value of the LogContextKey out of the context
	value := ctx.Value(LogContextKey)
	if value == nil {
		defaultlog.Error("Context was missing log key; returning fresh logger", "logKey", LogContextKey)
		return freshLogger()
	}

	// Verify the value of the key matches the expected type of log15.Logger
	log, ok := value.(log15.Logger)
	if !ok {
		log.Error("Type of log key did not match expected type (log15.Logger)", "logKey", LogContextKey, "type", fmt.Sprintf("%T", log))
		return freshLogger()
	}

	return log
}

func freshLogger() log15.Logger {
	return log15.New(log15.Ctx{"module": "created-by-util/log"})
}

// GetLoggerFromRequestContext pulls a log15.Logger out of the context of an http.Request's context.
func GetLoggerFromRequestContext(r *http.Request) log15.Logger {
	return GetLoggerFromContext(r.Context())
}
