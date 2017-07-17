package router

import (
	"ToDoListServer/handler/event"
	"ToDoListServer/handler/user"

	"ToDoListServer/handler"
	"github.com/labstack/echo"
)

func InitRouter(e *echo.Echo) {
	if e == nil {
		panic("[InitRouter], server couldn't be nil")
	}

	// user
	e.POST("/", user.Login)
	e.GET("/", handler.LoginHtml)
	e.POST("/register", user.Register)
	e.GET("/register", handler.RegisterHtml)
	//event
	e.GET("/show", handler.Show)
	e.POST("/show", event.Create)

}
