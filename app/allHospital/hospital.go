package allhospital

import (
	"appointed-registration/server"

	"github.com/gin-gonic/gin"
)

func GetAllHostpitals(c *gin.Context) {
	err := server.GetAllHostpital()
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "结束",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "结束",
	})
}
