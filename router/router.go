package router

import (
	"ToDoListServer/handler/user"
	"ToDoListServer/handler/event"

	"github.com/labstack/echo"
	"ToDoListServer/handler"
)

func InitRouter(e *echo.Echo) {
	if e == nil {
		panic("[InitRouter], server couldn't be nil")
	}

	// user
	e.POST("/", user.Login)
	e.GET("/",handler.LoginHtml)
	e.POST("/register",user.Register)
	e.GET("/register",handler.RegisterHtml)
	// event
	e.POST("/v1/event/create", event.Create)
}
