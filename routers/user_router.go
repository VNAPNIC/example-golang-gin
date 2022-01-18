package routers

import (
	"net/http"
	authHandler "serverhealthcarepanel/handlers/auth"
	userHandler "serverhealthcarepanel/handlers/user"
	"serverhealthcarepanel/middleware"

	"github.com/labstack/echo/v4"
)

func InitUserRouter(Router *echo.Group) {
	Router.POST("/login", authHandler.UserLogin)
	Router.POST("/register", userHandler.CreateUser)

	groupUser := Router.Group("/user", middleware.JWTHandler())
	groupUser.GET("", getUsers)
	groupUser.PUT("/change-password", authHandler.ChangePassword)
}

var users = []string{"Joe", "Veer", "Zion"}

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}
