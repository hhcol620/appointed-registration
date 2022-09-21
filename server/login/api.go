package login

import (
	"appointed-registration/helper"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// 获取所有的验证码
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

	cc, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(cc))
	out, _ := os.Create("vifyCode.png")
	defer out.Close()
	_, err = io.Copy(out, bytes.NewReader(cc))
	// return response.Cookies()
	fmt.Println(response.Header["Set-Cookie"])
}

// 验证验证码, 并且发送短信验证码
func CheckCode(code string) {
	request, err := http.NewRequest("GET",
		fmt.Sprintf("https://www.114yygh.com/web/checkcode?_time=%v&code=%v", time.Now().UnixMilli(), code), nil)

	if err != nil {
		log.Println()
		return
	}

	// 设置普通头部
	helper.SetHeads(request, "agent_login_img_code=5c1f401de1444c5b9984f851f764f60a; Expires=Wed, 21-Sep-2022 06:04:38 GMT; Path=/; HttpOnly imed_session=xb45zPuipLN5BSdQOAAFOFSzmnbCRmNZ_5545794; Domain=.114yygh.com; Path=/; HttpOnly imed_session_tm=1663738478252; Domain=.114yygh.com; Path=/; HttpOnly")
	// request.Header.Add("Cookie", "agent_login_img_code=e57070abac624d198906870570d3dc13; Path=/; Expires=Wed, 21 Sep 2022 05:07:06 GMT; HttpOnly imed_session=2pA59XrK3AHCUuVZ7ldAqAwG6M7Wpgw8_5545783; Path=/; Domain=114yygh.com; HttpOnly imed_session_tm=1663735026199; Path=/; Domain=114yygh.com; HttpOnly")

	client := &http.Client{}

	response, _ := client.Do(request)
	cc, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(cc))
}

// 发送手机验证码
// func
