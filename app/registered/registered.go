package registered

import (
	"appointed-registration/server"

	"github.com/gin-gonic/gin"
)

/**
* 代码描述: 获取票数的接口
* 作者:小大白兔
* 创建时间:2022/10/09 21:13:29
 */
func Registered(c *gin.Context) {
	err := server.GetRegister("162", "10c186f26ae7ecf8160e2dcf1f2e7312", "200053529")
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok了",
	})
}
