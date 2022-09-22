package dao

import (
	"appointed-registration/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql() {
	config, err := config.GetAppointedRegistrationConfig()
	if err != nil {
		log.Println("解析错误: ", err)
		return
	}
	fmt.Println(config)
	// 线上
	dsn := config.AppointedRegistration.MysqlConfig.Username + ":" + config.AppointedRegistration.MysqlConfig.Password + "@tcp(" +
		config.AppointedRegistration.MysqlConfig.IP + ":" + config.AppointedRegistration.MysqlConfig.Port + ")/" + config.AppointedRegistration.MysqlConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	// 线下
	// dsn := "root:123456@tcp(127.0.0.1:3306)/gin_wall?charset=utf8mb4&parseTime=True&loc=Local"
	//全局模式

	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Println("连接数据库失败")
		return
	}
	fmt.Println("连接成功")
}
