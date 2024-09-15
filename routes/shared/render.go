package shared

import (
	"net/http"

	"chatProject/routes"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func Page(ctx echo.Context, pageContent templ.Component, scripts templ.Component) error {
	layout := routes.Layout(pageContent, scripts)
	return Render(ctx, http.StatusOK, layout)
}
