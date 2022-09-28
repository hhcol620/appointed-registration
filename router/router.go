package router

import (
	"appointed-registration/app/login"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	// 登录相关接口
	loginRouter := Router.Group("login")
	{
		loginRouter.POST("imgCode", login.GetImgCode)      // 获取图片验证码
		loginRouter.POST("sendCode", login.SendMobileCode) // 手机短信验证码的发送
		loginRouter.POST("login", login.Login)             // 登录
	}
}
