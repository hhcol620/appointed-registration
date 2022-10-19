package listdepartment

import (
	"appointed-registration/app/models"
	"appointed-registration/global"
	"appointed-registration/helper/response"
	"appointed-registration/server"

	"github.com/gin-gonic/gin"
)

/**
* 代码描述:获取所有的科室
* 作者:小大白兔
* 创建时间:2022/10/07 18:51:51
 */
func GetDepartmentByCode(c *gin.Context) {
	hosCode := &models.GetHostpitalCode{}

	if err := c.ShouldBindJSON(&hosCode); err != nil {
		global.LogSuger.Errorf("获取参数失败: " + err.Error())

		response.ErrClient(c, 400, err.Error())
	}

	serve := server.NewServers("", hosCode.HostpitalCode)

	result, err := serve.GetDepartment()
	if err != nil {
		global.LogSuger.Errorf("获取科室失败: " + err.Error())

		response.ErrClient(c, 400, err.Error())
	}

	response.ReturnWithData(c, 200, result)
}
