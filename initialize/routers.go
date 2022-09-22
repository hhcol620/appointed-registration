package initialize

import (
	"appointed-registration/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("u/v1")
	router.InitUserRouter(ApiGroup)
	return Router
}
