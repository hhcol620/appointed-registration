package router

import (
	"appointed-registration/app/login"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	// 登录相关接口
	loginRouter := Router.Group("login")
	{
		loginRouter.POST("mobile", login.GetMobile) // 手机号码的输入
		loginRouter.POST("code", login.GetCode)     // 手机验证码的输入, 然后登录
	}
}
