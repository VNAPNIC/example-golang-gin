package routers

import (
	"github.com/gin-gonic/gin"
	roleHandler "healthcare-panel/handlers/role"
	"healthcare-panel/middleware"
)

func InitRoleRouter(Router *gin.RouterGroup) {
	role := Router.Group("/role").Use(
		middleware.JWTHandler(),
	)
	{
		role.POST("", roleHandler.CreateRole)
	}
}
