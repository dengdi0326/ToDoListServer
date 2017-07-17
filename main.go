package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"ToDoListServer/handler"
	"ToDoListServer/handler/event"
	"ToDoListServer/router"
)

func main() {
	e := echo.New()
	e.Renderer = handler.Temlate()
	e.Renderer = handler.Temlate()
	e.Use(middleware.Recover())

	router.InitRouter(e)

	e.Start(":7100")
}
