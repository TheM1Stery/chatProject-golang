package index

import (
	"chatProject/routes/shared"
	"net/http"

	"github.com/labstack/echo/v4"
)

type route struct {
}

func ConfigureRoutes(e *echo.Echo) {
	routes := route{}

	e.GET("/", routes.index)
}

func (route *route) index(c echo.Context) error {
	return shared.Render(c, http.StatusOK, hello("Maksud"))
}
