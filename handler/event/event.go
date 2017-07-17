package event

import (
	"ToDoListServer/orm"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Temlate() *Template {
	t := &Template{
		templates: template.Must(template.ParseGlob("html/*.gtpl")),
	}

	return t
}

func Show(c echo.Context) error {
	return c.Render(http.StatusOK, "show.gtpl", nil)
}

func Create(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
		http.Redirect(c.Response(), c.Request(), "/login", http.StatusSeeOther)
	}
	text := c.FormValue("event")
	return orm.Add(c, cookie.Value, text)
	//http.Redirect(c.Response(),c.Request(), "/show", http.StatusSeeOther)
	//return nil
}
