package main

import (
	"appointed-registration/dao"
	"appointed-registration/initialize"
	"appointed-registration/logger"
)

func main() {
	logger.NewInitLogger() // 日志的初始化

	dao.InitMysql()

	Router := initialize.Routers()

	Router.Run(":8081")
}
