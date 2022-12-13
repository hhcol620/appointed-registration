package registered

import (
	"appointed-registration/global"
	v1 "appointed-registration/helper/request"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// 筛选出需要查询的
type RegisterCode struct {
	HostCode       string `json:"hosCode"`
	FirstDeptCode  string `json:"firstDeptCode"`
	SecondDeptCode string `json:"secondDeptCode"`
	Week           int    `json:"week"`
}

/**
* 代码描述:号码查询
* 作者:小大白兔
* 创建时间:2022/10/24 16:59:42
 */
func (r *RegisterCode) Check() (*http.Response, error) {
	// 通过传入数据医院,

	params := RegisterCode{
		HostCode:       r.HostCode,
		FirstDeptCode:  r.FirstDeptCode,
		SecondDeptCode: r.SecondDeptCode,
		Week:           1,
	}

	param, err := json.Marshal(params)
	if err != nil {
		global.LogSuger.Errorf("数据解析失败: " + err.Error())
		return nil, errors.New("数据解析失败: " + err.Error())
	}

	request := v1.PostRequest(fmt.Sprintf("https://www.114yygh.com/web/product/list?_time=%v", time.Now().UnixMilli()), string(param))

	strCookie, err := global.RedisDb.Get(global.Ctx, fmt.Sprintf("register:%v", global.Phone)).Result()
	if err != nil {
		global.LogSuger.Errorf("获取cookie失败: " + err.Error())
		return nil, errors.New("获取数据失败: " + err.Error())
	}
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Cookie", strCookie)
	request.Header.Set("Host", "www.114yygh.com")
	request.Header.Set("Origin", "https://www.114yygh.com")
	request.Header.Set("Referer", "https://www.114yygh.com/hospital/146/5072ec36c845e36c95d59606f2451089/200087753/source")
	request.Header.Set("Request-Source", "PC")
	request.Header.Set("sec-ch-ua", `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	request.Header.Set("sec-ch-ua-mobile", "?0")
	request.Header.Set("sec-ch-ua-platform", "Windows")
	request.Header.Set("Sec-Fetch-Dest", "empty")
	request.Header.Set("Sec-Fetch-Mode", "cors")
	request.Header.Set("Sec-Fetch-Site", "same-origin")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {

		global.LogSuger.Errorf("响应失败: " + err.Error())
		return nil, errors.New("响应失败: " + err.Error())
	}

	return response, nil
}

func solveSetCookie(cookie string) string {
	dd := strings.Split(cookie, ";")
	return dd[0] + "; "
}

func solveSetCookieS(cookie string) string {
	dd := strings.Split(cookie, ";")
	return dd[0]
}
