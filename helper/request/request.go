package request

import (
	"bytes"
	"log"
	"net/http"
)

// 封装get请求, 并且返回request
func GetRequest(url string) *http.Request {

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("get请求失败: ", err)
		return nil
	}

	return request

}

// 封装post 请求
func PostRequest(url, param string) *http.Request {

	params := []byte(param)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	if err != nil {
		log.Println("get请求失败: ", err)
		return nil
	}

	return request
}
