package registered

import (
	v1 "appointed-registration/helper/request"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func test() {
	request := v1.GetRequest(fmt.Sprintf("https://www.114yygh.com/web/hospital/authority?_time=%v&hosCode=120", time.Now().UnixMilli()))

	client := &http.Client{}
	// helper.SetHead(request)
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Cookie", "imed_session=y93DYH6MMN5gEfTDeJ9O7eoAbODDAj6p_5545804; imed_session=ixxz3gaf2v72HOnXtoF4Gt8BRrrI63iF_5545812; secure-key=cca45fab-7f7e-4cda-b810-3e87fc509f15; imed_session=ixxz3gaf2v72HOnXtoF4Gt8BRrrI63iF_5545812; agent_login_img_code=6a066ac944a2471aba5931f21a0080c0; cmi-user-ticket=_r8I0Vj1tuyfgR1KGp2igo029usSESN0q417zA..; imed_session_tm=1663743776210")
	request.Header.Set("Host", "www.114yygh.com")
	request.Header.Set("Referer", "https://www.114yygh.com/hospital/120/a660294efe4daaf0bcbff7d69225ce5b/200000909/source")
	request.Header.Set("Request-Source", "PC")
	request.Header.Set("sec-ch-ua", `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	request.Header.Set("sec-ch-ua-mobile", "?0")
	request.Header.Set("sec-ch-ua-platform", "Windows")
	request.Header.Set("Sec-Fetch-Dest", "empty")
	request.Header.Set("Sec-Fetch-Mode", "cors")
	request.Header.Set("Sec-Fetch-Site", "same-origin")

	response, err := client.Do(request)
	if err != nil {
		log.Println("响应失败: ", err)
		return
	}
	sss, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(sss))
}

// 筛选出需要查询的
type mou struct {
	HostCode       string `json:"hosCode"`
	FirstDeptCode  string `json:"firstDeptCode"`
	SecondDeptCode string `json:"secondDeptCode"`
	Week           int    `json:"week"`
}

// 传入医院的所有的数据进行查询
func Check(hasCode, firstDeptCode, secondDeptCode string) {
	// 通过传入数据医院,

	params := mou{
		HostCode:       hasCode,
		FirstDeptCode:  firstDeptCode,
		SecondDeptCode: secondDeptCode,
		Week:           1,
	}
	param, err := json.Marshal(params)
	if err != nil {
		fmt.Println("转换失败: ", err)
		return
	}
	// fmt.Println(string(param))
	request := v1.PostRequest("https://www.114yygh.com/web/product/list?_time=1663920385245", string(param))

	// request.Header.Add("Content-Type", "application/json")
	// helper.SetRegisteredHead(request)
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Cookie", "imed_session=OWuwDUZhOtU1Me2Mx2v0Rzb4JE57zKZm_5546400; secure-key=d0d5cc25-619b-4d08-9cd6-c7d2c97de66a; imed_session=OWuwDUZhOtU1Me2Mx2v0Rzb4JE57zKZm_5546400; agent_login_img_code=f4a744d1dbdd4255a9c08accbe618b6c; cmi-user-ticket=zwAYVjg8QMKiN0GNRZuSq08nhd5H-rjRwFjqgA..; imed_session_tm=1663920380957")
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
	// request.Header.Add("Cookie", "cmi-user-ticket=e0XeBg4Cp_d_Ya4sAWRAWOR_r6XYbm3_tEPXiA..; Domain=.114yygh.com; Expires=Mon, 09-Oct-2090 16:53:40 GMT; Path=/; HttpOnly; imed_session=uPb6UyObagj6qr6NESLYRc3hGE8uHXqw_5545891; HttpOnly; imed_session_tm=1663767573078; Domain=.114yygh.com; Path=/; HttpOnly")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("响应错误: ", err)
		return
	}
	sss, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(sss))
	fmt.Println(response.StatusCode)
}
