package login

import (
	"appointed-registration/helper"
	aess "appointed-registration/helper/aes"
	v1 "appointed-registration/helper/request"
	"strings"

	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type login struct {
	Code   string `json:"code"`
	Mobile string `json:"mobile"`
}

var (
	str     = ""
	cookie2 = "cmi-user-ticket=vauqiml33O-y3aPzwqRwZKAd-xuZPbDqvlFrWw..;"
)

// 获取图片验证码
func getImgCode() {
	request, err := http.NewRequest("GET",
		fmt.Sprintf("https://www.114yygh.com/web/img/getImgCode?_time=%v", time.Now().UnixMilli()), nil)

	if err != nil {
		log.Println()
		return
	}

	// 设置普通头部
	helper.SetHead(request)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("响应失败: ", err)
		return
	}

	cc, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(cc))
	out, _ := os.Create("vifyCode.png")
	defer out.Close()
	_, err = io.Copy(out, bytes.NewReader(cc))

	str = "tqBih5MhOSWMKEJVGMkdn2kRV0VyaCzJ_5546196; " +
		solveSetCookie(response.Header["Set-Cookie"][0]) + cookie2 + "; " + "secure-key=35baa7d5-a68c-4c0b-b953-33ac7dcb478f; " +
		solveSetCookie(response.Header["Set-Cookie"][2]) +
		solveSetCookie(response.Header["Set-Cookie"][1]) + solveSetCookie(response.Header["Set-Cookie"][3])

}

// checkCode, 并且发送短信验证码
func CheckCode(code string) {
	request, err := http.NewRequest("GET",
		fmt.Sprintf("https://www.114yygh.com/web/checkcode?_time=%v&code=%v", time.Now().UnixMilli(), code), nil)

	if err != nil {
		log.Println()
		return
	}

	// fmt.Println("这是组装的cookie: ", str)
	request.Header.Add("Cookie", "tqBih5MhOSWMKEJVGMkdn2kRV0VyaCzJ_5546196; secure-key=38104538-0515-46dd-a97f-ed30fe712f3e; cmi-user-ticket=vauqiml33O-y3aPzwqRwZKAd-xuZPbDqvlFrWw..;; secure-key=35baa7d5-a68c-4c0b-b953-33ac7dcb478f; imed_session=yI6V89niXO1awLO5usL9seh3M39YD00a_5546202; agent_login_img_code=bf944c661f124d4cb3eeeee2e5292959; imed_session_tm=1663860845399;")

	request.Header.Set("Content-Type", "application/json")
	helper.SetHead(request)

	client := &http.Client{}

	response, _ := client.Do(request)
	cc, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(cc))

}

// 验证验证码 verfiyCode(手机号码和图形验证码)
func VerfiyCode(mobile, code string) {

	// 手机号码加密
	mobileAes := aess.AesMobile(mobile)

	// 替换符号
	mobileAesSwap := strings.ReplaceAll(mobileAes, "=", "%3D")

	// 将加密的手机号码的 ` '=' 转换为 '%3D' `
	resquest := v1.GetRequest(
		fmt.Sprintf("https://www.114yygh.com/web/common/verify-code/get?_time=%v&mobile=%v&smsKey=LOGIN&code=%v", time.Now().UnixMilli(), mobileAesSwap, code))

	// 设置用户的头部信息
}

// 实现登录(手机号码和短信验证码)
func Login(mobile, code string) (*http.Response, error) {
	// 实现加密
	mobileAes, codeAes := aess.AesECBPass(mobile, code)
	loginBody := login{
		Code:   codeAes,
		Mobile: mobileAes,
	}
	param, err := json.Marshal(loginBody)
	if err != nil {
		log.Println("json 解析失败: ", err)
		return nil, errors.New("解析失败: " + err.Error())
	}

	request := v1.PostRequest(fmt.Sprintf("https://www.114yygh.com/web/login?_time=%v", time.Now().UnixMilli()), string(param))

	request.Header.Set("Cookie", "imed_session=tqBih5MhOSWMKEJVGMkdn2kRV0VyaCzJ_5546196; cmi-user-ticket=vauqiml33O-y3aPzwqRwZKAd-xuZPbDqvlFrWw..; secure-key=07641dae-76ee-4bf9-83b0-74c59e149afa; imed_session=tqBih5MhOSWMKEJVGMkdn2kRV0VyaCzJ_5546196; agent_login_img_code=6a09687c8cbe426c9611decc3c4344c7; imed_session_tm=1663858985974")
	request.Header.Set("Content-Type", "application/json")
	helper.SetHead(request)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("响应错误: ", err)
		return nil, errors.New("响应失败: " + err.Error())
	}
	return response, nil
}

//  处理某一个cookie为发起的格式
func solveSetCookie(cookie string) string {
	dd := strings.Split(cookie, ";")
	return dd[0] + "; "
}
