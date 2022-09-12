package alladdress

import (
	"appointed-regidtration/helper"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type address struct {
	code         string
	levelText    string
	maintain     bool
	name         string
	openTimeText string
	picture      string
}

// 获取医院相关数据
func GetAddress() (map[string]address, error) {

	result, arr := map[string]interface{}{}, map[string]address{}

	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://www.114yygh.com/web/hospital/list?_time=1662970211457&keywords=&levelId=0&areaId=0&pageNo=1&pageSize=20", nil)
	if err != nil {
		log.Println("err:", err)
		return nil, errors.New("err: " + err.Error())
	}

	helper.SetHead(request)

	response, err := client.Do(request)
	if err != nil {
		log.Println("err:", err)
		return nil, errors.New("err: " + err.Error())
	}
	cc, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(cc, &result)
	if err != nil {
		log.Println(err)
		return nil, errors.New("err: " + err.Error())
	}

	// 遍历所有的数据
	for _, v := range result["data"].(map[string]interface{})["list"].([]interface{}) {

		res := address{
			code:         (v.(map[string]interface{}))["code"].(string),
			levelText:    (v.(map[string]interface{}))["levelText"].(string),
			maintain:     (v.(map[string]interface{}))["maintain"].(bool),
			name:         (v.(map[string]interface{}))["name"].(string),
			openTimeText: (v.(map[string]interface{}))["openTimeText"].(string),
			picture:      (v.(map[string]interface{}))["picture"].(string),
		}

		arr[(v.(map[string]interface{}))["name"].(string)] = res
	}

	return arr, nil
}
