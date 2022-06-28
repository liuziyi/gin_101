package routers

import (
	"gin_101/pkg/middleware/jwt"
	"gin_101/routers/api"
	"gin_101/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.POST("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/ping", v1.HandlePing)
	}
	return r
}
