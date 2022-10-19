package login

import (
	"appointed-registration/app/models"
	"appointed-registration/global"
	"appointed-registration/helper/response"
	v1 "appointed-registration/server"
	"log"

	"github.com/gin-gonic/gin"
)

/**
* 代码描述: 图片验证码
* 作者:小大白兔
* 创建时间:2022/09/26 22:13:33
 */
func GetImgCode(ctx *gin.Context) {
	getMobile := &models.GetImgCode{}

	if err := ctx.ShouldBindJSON(getMobile); err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		response.ErrClient(ctx, 400, err.Error())
		return

	}

	server := v1.NewServers(getMobile.Mobile, "")

	global.Phone = getMobile.Mobile

	err := server.GetImgCode()

	if err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Println("获取失败: ", err)
		return
	}

	response.ReturnSuccess(ctx)
}

/**
* 代码描述:手机验证码发送
* 作者:小大白兔
* 创建时间:2022/09/28 20:33:09
 */
func SendMobileCode(c *gin.Context) {

	getCode := &models.GetCode{}

	if err := c.ShouldBindJSON(getCode); err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		response.ErrClient(c, 400, err.Error())
		return

	}

	server := v1.NewServers("", "")

	if err := server.GetMobileCode(getCode.Code); err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		response.ErrServer(c, 500, err.Error())
		return

	}

	// 发送成功
	response.ReturnSuccess(c)
}

/**
* 代码描述: 实现登录
* 作者:小大白兔
* 创建时间:2022/09/28 21:14:24
 */
func Login(c *gin.Context) {

	getCode := &models.GetCode{}

	if err := c.ShouldBindJSON(getCode); err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		response.ErrClient(c, 400, err.Error())
		return

	}

	server := v1.NewServers(global.Phone, "")

	if err := server.GetLogin(getCode.Code); err != nil {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		response.ErrServer(c, 500, err.Error())
		return
	}

	// 发送成功
	response.ReturnSuccess(c)
}
