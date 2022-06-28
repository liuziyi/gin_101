package v1

import (
	"gin_101/pkg/app"
	"gin_101/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlePing(c *gin.Context) {
	appG := app.Gin{C: c}
	data := "sdf"
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
