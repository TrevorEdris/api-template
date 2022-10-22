package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	// JSONResponse defines the information required to construct a JSON HTTP response.
	JSONResponse struct {
		StatusCode int
		Headers    map[string]string
		RequestID  string
		Path       string
		URL        string
		Context    echo.Context
		ToURL      func(name string, params ...interface{}) string
		Body       interface{}
	}

	errResponse struct {
		Message string `json:"message"`
	}
)

// NewJSONResponse creates a new instance of a JSONResponse struct given an echo context.
func NewJSONResponse(ctx echo.Context) JSONResponse {
	return JSONResponse{
		Context:    ctx,
		ToURL:      ctx.Echo().Reverse,
		Path:       ctx.Request().URL.Path,
		URL:        ctx.Request().URL.String(),
		StatusCode: http.StatusOK,
		Headers:    make(map[string]string),
		RequestID:  ctx.Response().Header().Get(echo.HeaderXRequestID),
	}
}
