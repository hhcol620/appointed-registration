package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var configFile []byte

type ConfigAppointed struct {
	AppointedRegistration Config `yaml:"appointedRegistration"`
}

type Config struct {
	MysqlConfig mysqlconfig `yaml:"mysql"`
	Redisconfig redisconfig `yaml:"redis"`
}

// MySQL 数据库
type mysqlconfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	IP       string `yaml:"ip"`
	Database string `yaml:"database"`
}

// redis 数据持久层, 用于储存数据的持久性
type redisconfig struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

func GetAppointedRegistrationConfig() (e *ConfigAppointed, err error) {
	err = yaml.Unmarshal(configFile, &e)
	return e, err
}
func init() {
	var err error

	configFile, err = ioutil.ReadFile("G:/appointed-registration/config/config.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
	}

}
