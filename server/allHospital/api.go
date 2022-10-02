package allhospital

import (
	"appointed-registration/helper"
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
func GetAddress(levelId int, areaId string, pageNo int) (string, error) {

	// result := map[string]interface{}{}

	client := &http.Client{}

	request, err := http.NewRequest("GET",
		fmt.Sprintf("https://www.114yygh.com/web/hospital/list?_time=%v&keywords=&levelId=%v&areaId=%v&pageNo=%v&pageSize=20",
			time.Now().UnixMilli(), levelId, areaId, pageNo), nil)

	if err != nil {
		log.Println("err:", err)
		return "", errors.New("err: " + err.Error())
	}

	helper.SetHead(request)

	response, err := client.Do(request)
	if err != nil {
		log.Println("err:", err)
		return "", errors.New("err: " + err.Error())
	}

	cc, _ := ioutil.ReadAll(response.Body)
	return string(cc), nil
	// err = json.Unmarshal(cc, &result)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, errors.New("err: " + err.Error())
	// }
	// return result, nil

}
