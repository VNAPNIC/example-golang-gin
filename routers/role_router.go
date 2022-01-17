package routers

import (
	roleHandler "serverhealthcarepanel/handlers/role"

	"github.com/labstack/echo/v4"
)

func InitRoleRouter(Router *echo.Group) {
	role := Router.Group("/role") // , middleware.JWTHandler()
	role.POST("", roleHandler.CreateRole)
}
