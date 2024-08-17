package index

import (
	"chatProject/routes/shared"

	"github.com/labstack/echo/v4"
)

type route struct {
}

func ConfigureRoutes(e *echo.Echo) {
	routes := route{}

	e.GET("/", routes.index)

}

func (route *route) index(ctx echo.Context) error {
	return shared.Page(ctx, hello("Seymur"), nil)
}
