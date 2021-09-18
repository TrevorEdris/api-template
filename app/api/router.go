package api

import (
	"net/http"

	redocly "github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/inconshreveable/log15"

	"github.com/TrevorEdris/api-template/app/api/middleware"
	v1 "github.com/TrevorEdris/api-template/app/api/v1"
	"github.com/TrevorEdris/api-template/app/models/item"
	"github.com/TrevorEdris/api-template/app/util"
)

const (
	openapiDocsPath = "/resources/docs/swagger.yaml"
)

type router struct {
	jwtIssuer string
	log       log15.Logger
}

// Models is a list of all the models the API iteracts with
type Models struct {
	Items *item.Items
}

// NewRouter creates an instance of a router, returning the HTTP handler.
func NewRouter(jwtIssuer string, log log15.Logger, models Models) http.Handler {
	router := &router{
		jwtIssuer: "TODO",
		log:       log,
	}
	return router.init(models)
}

func (router *router) init(models Models) http.Handler {
	r := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	apiV1 := v1.New(util.GetLoggerFromRequestContext, models.Items)

	// Serve the swagger docs via a file server, but rendered with redocly
	docsHandler := redocly.Redoc(redocly.RedocOpts{SpecURL: openapiDocsPath}, nil)

	// TODO: Implement auth
	addMiddleware := func(path string, h http.HandlerFunc) http.Handler {
		return middleware.Auth(router.jwtIssuer, &router.log, h)
	}

	r.Handle("/docs", docsHandler)
	r.Handle("/v1/health", addMiddleware("/v1/health", apiV1.Health)).Methods(http.MethodGet)
	r.Handle("/v1/generalkenobi", addMiddleware("/v1/generalkenobi", apiV1.GeneralKenobi)).Methods(http.MethodGet)
	r.Handle("/v1/echo", addMiddleware("/v1/echo", apiV1.Echo)).Methods(http.MethodPost)

	return middleware.Log(router.log, r)
}
