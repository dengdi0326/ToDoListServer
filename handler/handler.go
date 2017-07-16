//这里做渲染模板
package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func LoginHtml(c echo.Context){
	return c.Render(http.StatusOK,"login.gtpl",nil)
}

func RegisterHtml(c echo.Context){
	return c.Render(http.StatusOK,"register.gtpl",nil)
}

//这里做显示页面模板的渲染，加上事件map 作为 render函数的参数
func ShowHtml(c echo.Context){
	return
}

