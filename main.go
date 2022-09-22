package main

import (
	"appointed-registration/dao"
	"appointed-registration/initialize"
)

func main() {
	dao.InitMysql()
	Router := initialize.Routers()
	Router.Run(":8080")
}
