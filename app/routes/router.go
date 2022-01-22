package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	"github.com/TrevorEdris/api-template/app/controller"
	"github.com/TrevorEdris/api-template/app/middleware"
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
		middleware.LogRequestID(),
		echomw.TimeoutWithConfig(echomw.TimeoutConfig{
			Timeout: c.Config.App.Timeout,
		}),
	)

	ctr := controller.NewController(c)

	defaultRoutes(c, g, ctr)

	// Create a group of routes where the json content type is enforced
	jsonGroup := g.Group("", middleware.EnforceContentType(middleware.JSON))
	itemRoutes(c, jsonGroup, ctr)
}

func defaultRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	hello := Hello{Controller: ctr}
	g.GET("/", hello.Get).Name = "helloworld"
}

func itemRoutes(c *services.Container, g *echo.Group, ctr controller.Controller) {
	item := Item{Controller: ctr}
	g.GET("/item/:id", item.Get).Name = "itemget"
	g.POST("/item", item.Post).Name = "itempost"
	g.PUT("/item/:id", item.Put).Name = "itemput"
	g.DELETE("/item/:id", item.Delete).Name = "itemdelete"
}
