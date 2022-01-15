package routers

import "github.com/labstack/echo/v4"

func InitRouter(e *echo.Echo) {
	v1 := e.Group("/v1/api")
	InitUserRouter(v1)
}