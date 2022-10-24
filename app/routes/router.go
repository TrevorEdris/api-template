package routes

import (
	"net/http"
	"strings"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	"github.com/TrevorEdris/api-template/app/controller"
	"github.com/TrevorEdris/api-template/app/middleware"
	"github.com/TrevorEdris/api-template/app/services"
)

// BuildRouter builds the HTTP router for all endpoint handlers.
func BuildRouter(c *services.Container) {
	g := c.Web.Group("")

	// Force HTTPS if enabled
	if c.Config.HTTP.TLS.Enabled {
		g.Use(echomw.HTTPSRedirect())
	}

	// Use various other middleware functions here
	g.Use(
		echomw.RemoveTrailingSlashWithConfig(echomw.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		}),
		//echomw.Recover(),
		echomw.Secure(),
		echomw.RequestID(),
		echomw.Gzip(),
		middleware.LogRequestID(),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Timeout: c.Config.App.Timeout,
		}),
	)

	ctr := controller.NewController(c)

	defaultRoutes(c, g, ctr)

	itemGroup := g.Group("/items")
	itemRoutes(c, itemGroup, ctr)

	p := prometheus.NewPrometheus(strings.ReplaceAll(c.Config.App.Name, "-", "_"), nil)
	p.Use(c.Web)
}

// TODO: Add /health, /ready
// TODO: Add OpenTelemetry wrappers
func defaultRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	hello := Hello{Controller: ctr}
	g.GET("/", hello.Get).Name = "helloworld"
	// TODO: Consider changing names. This was previously contained in a package, so "Wrap" is now ambigious.
	Wrap(c.Web)
}

// itemRoutes defines the mapping for the Item-related handlers.
// Example
// GET http://localhost:8080/item/1234 --> maps to the item.Get function.
func itemRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	item := Item{Controller: ctr}

	// Create a group of routes where the json content type is enforced
	requireJSON := g.Group("", middleware.EnforceContentType(middleware.JSON))
	g.GET("/:id", item.Get).Name = "itemget"
	g.DELETE("/:id", item.Delete).Name = "itemdelete"
	requireJSON.POST("", item.Post).Name = "itempost"
	requireJSON.PUT("/:id", item.Put).Name = "itemput"
}
