package allhospital

import (
	"appointed-registration/global"
	"appointed-registration/helper"
	"errors"
	"fmt"
	"io/ioutil"
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

	global.LogSuger.Info("GetAddress 接口请求开始...")

	client := &http.Client{}

	request, err := http.NewRequest("GET",
		fmt.Sprintf("https://www.114yygh.com/web/hospital/list?_time=%v&keywords=&levelId=%v&areaId=%v&pageNo=%v&pageSize=20",
			time.Now().UnixMilli(), levelId, areaId, pageNo), nil)

	if err != nil {
		global.LogSuger.Errorf("请求失败: " + err.Error())
		return "", errors.New("请求失败: " + err.Error())
	}

	helper.SetHead(request)

	response, err := client.Do(request)
	if err != nil {
		global.LogSuger.Info("响应失败: " + err.Error())
		return "", errors.New("响应失败: " + err.Error())
	}

	cc, _ := ioutil.ReadAll(response.Body)

	global.LogSuger.Info("GetAddress 接口请求结束...")

	return string(cc), nil

}
