package index

import (
	"github.com/labstack/echo/v4"
)

type route struct {
}

func ConfigureRoutes(e *echo.Echo) {
	routes := route{}

	e.GET("/", routes.index)
}

func (route *route) index(e echo.Context) error {
	return e.HTML(200, "<h1>Salam!</h1>")
}
