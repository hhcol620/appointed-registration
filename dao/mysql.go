package dao

import (
	"appointed-registration/config"
	"appointed-registration/global"
	"appointed-registration/models/address"
	"appointed-registration/models/hospital"
	"log"

	"github.com/go-redis/redis/v8"
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

	// 线上
	dsn := config.AppointedRegistration.MysqlConfig.Username + ":" + config.AppointedRegistration.MysqlConfig.Password + "@tcp(" +
		config.AppointedRegistration.MysqlConfig.IP + ":" + config.AppointedRegistration.MysqlConfig.Port + ")/" + config.AppointedRegistration.MysqlConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Println("连接数据库失败")
		return
	}

	// 数据库映射
	DB.AutoMigrate(
		&address.AllAddress{},
		&address.AllGrade{},
		&hospital.Hospital{},
	)

	global.RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		Network:  "tcp",
	})
}
