package routers

import (
	authHalder "serverhealthcarepanel/handlers/auth"
	userHalder "serverhealthcarepanel/handlers/user"

	"github.com/labstack/echo/v4"
)

func InitUserRouter(Router *echo.Group) {
	Router.POST("/login", authHalder.UserLogin)
	user := Router.Group("/user") // , middleware.JWTHandler()
	user.POST("", userHalder.CreateUser)
}
