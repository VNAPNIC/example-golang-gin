package routers

import (
	"github.com/gin-gonic/gin"
	authHandler "healthcare-panel/handlers/auth"
	userHandler "healthcare-panel/handlers/user"
	"healthcare-panel/middleware"
	"net/http"
)

func InitUserRouter(Router *gin.RouterGroup) {
	Router.POST("/login", authHandler.UserLogin)
	Router.POST("/register", userHandler.CreateUser)

	groupUser := Router.Group("/user").Use(
		middleware.JWTHandler(),
	)
	{
		groupUser.GET("", getUsers)
		groupUser.PUT("/logout", authHandler.UserLogout)
		groupUser.PUT("/change-password", authHandler.ChangePassword)
	}
}

var users = []string{"Joe", "Veer", "Zion"}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
