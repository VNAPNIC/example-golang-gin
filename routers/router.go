package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) *gin.Engine {
	v1 := g.Group("/v1/api")
	{
		InitRoleRouter(v1)
		InitUserRouter(v1)
	}

	InitSwaggerRouter(g)

	return g
}
