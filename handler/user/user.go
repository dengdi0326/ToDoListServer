package user

import (
	"net/http"

	"github.com/labstack/echo"

	"ToDoListServer/model/user"
	"ToDoListServer/orm"
)

// 这里做从 request 获取传入参数，并调用model中的获取 Map 中对应数据 并比较的处理。
func Login(c echo.Context) error {
	pass := user.Login(c.FormValue("name"))

	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = c.FormValue("name")
	c.SetCookie(cookie)

	if pass == "" {
		http.Redirect(c.Response(), c.Request(), "/register", http.StatusSeeOther)
	} else {
		if pass == c.FormValue("pass") {
			http.Redirect(c.Response(), c.Request(), "/show", http.StatusSeeOther)
		}
		http.Redirect(c.Response(), c.Request(), "/", http.StatusSeeOther)
	}
	return nil
}

//这里做从 request 获取传入参数，并调用orm 中的 创建用户  并添加用户
func Register(c echo.Context) error {
	orm.DbmapServer.Creater(c.FormValue("name"), c.FormValue("pass"))
	http.Redirect(c.Response(), c.Request(), "/show", http.StatusSeeOther)
	return nil
}
