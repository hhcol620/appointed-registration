package dao

import (
	"appointed-registration/config"
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
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

	//  redis连接
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Println("连接数据库失败")
		return
	}
	fmt.Println("连接成功")
}
