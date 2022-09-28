package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
* 代码描述: 响应封装
* 作者:小大白兔
* 创建时间:2022/09/28 16:11:15
 */

// 带参数返回
func ReturnWithData(c gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": httpCode,
		"data": data,
	})
}

// 4** 客户端的错误
func ErrClient(c gin.Context, httpCode int, msg string) {
	c.JSON(httpCode, gin.H{
		"code":    httpCode,
		"success": "false",
		"error":   msg,
	})
}

// 5** 服务端的错误
func ErrServer(c *gin.Context, httpCode int, msg string) {
	c.JSON(httpCode, gin.H{
		"code":    httpCode,
		"success": "false",
		"error":   msg,
	})
}

// 请求成功的响应
func ReturnSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    "",
		"success": true,
	})
}
