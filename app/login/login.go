package login

import (
	"appointed-registration/app/models"

	"github.com/gin-gonic/gin"
)

/**
* 代码描述: 图片验证码
* 作者:小大白兔
* 创建时间:2022/09/26 22:13:33
 */
func GetImgCode(c *gin.Context) {
	getMobile := &models.GetImgCode{}

	if err := c.ShouldBindJSON(getMobile); err != nil {

	}
}

// 登录 (手机号码的获取)
func GetMobile(c *gin.Context) {

}

// 登录 (手机验证码的获取)
func GetCode(c *gin.Context) {

}
