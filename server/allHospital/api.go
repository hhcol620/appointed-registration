package alladdress

import (
	"appointed-registration/helper"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
func GetAddress(levelId int, areaId string, pageNo int) (map[string]address, error) {

	result, arr := map[string]interface{}{}, map[string]address{}

	client := &http.Client{}

	request, err := http.NewRequest("GET",
		fmt.Sprintf("https://www.114yygh.com/web/hospital/list?_time=%v&keywords=&levelId=%v&areaId=%v&pageNo=%v&pageSize=20",
			time.Now().UnixMilli(), levelId, areaId, pageNo), nil)

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

		if v == nil {
			return nil, errors.New("数据不能为空")
		}

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
