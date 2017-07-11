package router

import (
	"ToDoListServer/handler/user"
	"ToDoListServer/handler/event"

	"github.com/labstack/echo"
)

func InitRouter(e *echo.Echo) {
	if e == nil {
		panic("[InitRouter], server couldn't be nil")
	}

	// user
	e.POST("/v1/user/login", user.Login)

	// event
	e.POST("/v1/event/create", event.Create)
}
