package main

import (
	"ToDoListServer/router"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())

	router.InitRouter(e)

	e.Start(":7100")
}
