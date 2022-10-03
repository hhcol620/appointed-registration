package router

import (
	"appointed-registration/app/allhospital"
	"appointed-registration/app/listdepartment"
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

	// 所有的医院
	allHospitalss := Router.Group("hospital")
	{
		allHospitalss.GET("", allhospital.GetAllHostpitals)
	}

	// 一个医院的科室
	allDepartment := Router.Group("department")
	{
		allDepartment.POST("", listdepartment.ListDepartment)
	}
}
