package getip

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func Test_GetIp(t *testing.T) {
	response, err := GetIp()
	if err != nil {
		log.Println("test错误")
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("响应出错: " + err.Error())
		return
	}

	fmt.Println(string(body))
}
