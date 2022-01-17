package routers

import (
	"serverhealthcarepanel/handlers/role"
	"serverhealthcarepanel/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoleRouter(Router *echo.Group) {
	role := Router.Group("/role", middleware.JWTHandler())
	role.POST("", roleHandler.CreateRole)
}
