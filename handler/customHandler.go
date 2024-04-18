package handler

import (
	"github.com/Nithin-Valiyaveedu/go-htmx-echo/views/home"
	"github.com/labstack/echo"
)

type UserHandler struct {
}

func (h UserHandler) ShowHandler(c echo.Context) error {
	return home.Show().Render(c.Request().Context(), c.Response())
}
