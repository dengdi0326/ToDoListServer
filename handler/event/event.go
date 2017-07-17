package event

import (
	"net/http"

	"github.com/labstack/echo"

	"ToDoListServer/orm"
)

func Create(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		http.Redirect(c.Response(), c.Request(), "/login", http.StatusSeeOther)
		return err
	}
	text := c.FormValue("event")
	return orm.Add(c, cookie.Value, text)
}
