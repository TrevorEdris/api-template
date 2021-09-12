package api

import (
	"net/http"

	redocly "github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/inconshreveable/log15"

	"github.com/TrevorEdris/api-template/app/api/middleware"
	v1 "github.com/TrevorEdris/api-template/app/api/v1"
	"github.com/TrevorEdris/api-template/app/util"
)

const (
    openapiDocsPath = "/resources/docs/swagger.yaml"
)

type router struct {
    jwtIssuer string
    log log15.Logger
}

func NewRouter(jwtIssuer string, log log15.Logger) http.Handler {
    router := &router{
        jwtIssuer: "TODO",
        log: log,
    }
    return router.init()
}

func (router *router) init() http.Handler {
    r := mux.NewRouter().StrictSlash(true).UseEncodedPath()
    apiV1 := v1.New(util.GetLoggerFromRequestContext)

    // Serve the swagger docs via a file server, but rendered with redocly
    docsHandler := redocly.Redoc(redocly.RedocOpts{SpecURL: openapiDocsPath}, nil)

    // TODO: Implement auth
    addMiddleware := func(path string, h http.HandlerFunc) http.Handler {
        return middleware.Auth(router.jwtIssuer, &router.log, h)
    }

    r.Handle("/docs", docsHandler)
    r.Handle("/v1/health", addMiddleware("/v1/health", apiV1.Health)).Methods(http.MethodGet)
    r.Handle("/v1/generalkenobi", addMiddleware("/v1/generalkenobi", apiV1.GeneralKenobi)).Methods(http.MethodGet)

    return middleware.Log(router.log, r)
}
