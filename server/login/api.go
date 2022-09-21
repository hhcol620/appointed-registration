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
func GetImgCode() {
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
	fmt.Println(string(cc))
	out, _ := os.Create("vifyCode.png")
	defer out.Close()
	_, err = io.Copy(out, bytes.NewReader(cc))
}

// 验证验证码, 并且发送短信验证码
