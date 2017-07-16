package user

import (
	"github.com/labstack/echo"

	"ToDoListServer/model/user"
)

// 这里做从 request 获取传入参数，并调用model中的获取 Map 中对应数据 并比较的处理。
func Login(c echo.Context) error {
	pass:= user.Login(c.FormValue("name"))
	if c.FormValue("pass")!=pass{
		return error("not exsit the user")
	}
	return nil
}

