package routers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"serverhealthcarepanel/handlers"
	"serverhealthcarepanel/middleware"
)

func InitUserRouter(Router *echo.Group) {
	Router.POST("/login", handlers.UserLogin)
	user := Router.Group("/user", middleware.JWTHandler())
	user.GET("", getUsers)
}

var (
	users = []string{"Joe", "Veer", "Zion"}
)

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}
