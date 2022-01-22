package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	"github.com/TrevorEdris/api-template/app/controller"
	"github.com/TrevorEdris/api-template/app/services"
)

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
		echomw.Recover(),
		echomw.Secure(),
		echomw.RequestID(),
		echomw.Gzip(),
		// middleware.LogRequestID(),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Timeout: c.Config.App.Timeout,
		}),
	)

	ctr := controller.NewController(c)

	defaultRoutes(c, g, ctr)
}

func defaultRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	hello := Hello{Controller: ctr}
	g.GET("/", hello.Get).Name = "helloworld"
}
