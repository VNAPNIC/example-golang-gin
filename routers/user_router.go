package routers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	authHalder "serverhealthcarepanel/handlers/auth"
	userHalder "serverhealthcarepanel/handlers/user"
)

func InitUserRouter(Router *echo.Group) {
	Router.POST("/login", authHalder.UserLogin)
	user := Router.Group("/user") // , middleware.JWTHandler()
	user.POST("", userHalder.CreateUser)
	user.GET("", getUsers)
}

var (
	users = []string{"Joe", "Veer", "Zion"}
)

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}
