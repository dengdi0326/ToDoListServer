//这里做渲染模板
package handler

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
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

func LoginHtml(c echo.Context) error {
	return c.Render(http.StatusOK, "login.gtpl", nil)
}

func RegisterHtml(c echo.Context) error {
	return c.Render(http.StatusOK, "register.gtpl", nil)
}

func Show(c echo.Context) error {
	return c.Render(http.StatusOK, "show.gtpl", nil)
}