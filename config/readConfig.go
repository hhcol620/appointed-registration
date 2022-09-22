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
}

type mysqlconfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	IP       string `yaml:"ip"`
	Database string `yaml:"database"`
}

func GetAppointedRegistrationConfig() (e *ConfigAppointed, err error) {
	err = yaml.Unmarshal(configFile, &e)
	return e, err
}
func init() {
	var err error
	configFile, err = ioutil.ReadFile("../config/config.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
	}

}
