package util

import (
	"context"
	"fmt"
	"net/http"

	"github.com/inconshreveable/log15"
)

const (
    LogContextKey = "log-context-key"
)

var (
    // Create a default logger to use in the case of an error
    defaultlog = log15.New(log15.Ctx{"module": "util/log"})
)

type LogProvider func(r *http.Request) log15.Logger

func GetLoggerFromContext(ctx context.Context) log15.Logger {
    value := ctx.Value(LogContextKey)
    if value == nil {
        defaultlog.Error("Context was missing log key; returning fresh logger", "logKey", LogContextKey)
        return freshLogger()
    }

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

func GetLoggerFromRequestContext(r *http.Request) log15.Logger {
    return GetLoggerFromContext(r.Context())
}
