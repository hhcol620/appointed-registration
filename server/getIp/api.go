package getip

import (
	v1 "appointed-registration/helper/request"
	"errors"
	"log"
	"net/http"
)

/**
* 代码描述: 获取ip
* 作者:小大白兔
* 创建时间:2022/10/02 10:24:15
 */
func GetIp() (*http.Response, error) {
	request := v1.GetRequest("http://api.proxy.ipidea.io/getProxyIp?num=100&return_type=json&lb=1&sb=0&flow=1&regions=&protocol=http")

	request.Header.Set("Content-Type", "multipart/form-data")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("响应出错: " + err.Error())
		return nil, errors.New("响应出错: " + err.Error())
	}

	return response, nil
}

// http代理
func httpProxy()
