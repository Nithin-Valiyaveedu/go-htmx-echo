package main

import (
	"github.com/Nithin-Valiyaveedu/go-htmx-echo/handler"
	"github.com/labstack/echo"
)

func main() {
	app := echo.New()
	handlerr := handler.UserHandler{}
	app.GET("/test", handlerr.ShowHandler)
	app.Start(":3000")
}
