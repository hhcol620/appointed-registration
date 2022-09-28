package login

import (
	"appointed-registration/global"
	"appointed-registration/helper"
	aess "appointed-registration/helper/aes"
	v1 "appointed-registration/helper/request"
	"strings"

	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang/glog"
)

// 用于登录是信息的加密
type login struct {
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
}

type Login struct {
	Mobile string
}

var (
	cookie2 = "cmi-user-ticket=vauqiml33O-y3aPzwqRwZKAd-xuZPbDqvlFrWw..;"
)

// 获取图片验证码
func (l *Login) GetImgCode() (*http.Response, error) {

	request, err := http.NewRequest("GET",
		fmt.Sprintf("https://www.114yygh.com/web/img/getImgCode?_time=%v", time.Now().UnixMilli()), nil)

	if err != nil {
		log.Println()
		return nil, err
	}

	// 设置普通头部
	helper.SetHead(request)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("响应失败: ", err)
		return nil, err
	}

	// 用于验证码的cookie
	str := "tqBih5MhOSWMKEJVGMkdn2kRV0VyaCzJ_5546196; " +
		solveSetCookie(response.Header["Set-Cookie"][0]) + cookie2 + "; " + "secure-key=35baa7d5-a68c-4c0b-b953-33ac7dcb478f; " +
		solveSetCookie(response.Header["Set-Cookie"][2]) +
		solveSetCookie(response.Header["Set-Cookie"][1]) + solveSetCookieS(response.Header["Set-Cookie"][3])

	// 将数据存储到数据库
	_, err = global.RedisDb.Set(global.Ctx, fmt.Sprintf("checkCode:%v", l.Mobile), str, time.Hour*24*7).Result()
	if err != nil {
		log.Println("getCode 储存数据失败: ", err)
		return nil, err
	}

	return response, nil
}

// checkCode, 检验验证码(解决)
func (l *Login) CheckCode(code string) error {

	request, err := http.NewRequest("GET",
		fmt.Sprintf("https://www.114yygh.com/web/checkcode?_time=%v&code=%v", time.Now().UnixMilli(), code), nil)

	if err != nil {
		log.Println("发请求失败: " + err.Error())
		return errors.New("发请求失败: " + err.Error())
	}

	// 获取需要的cookie
	cookieStr, err := global.RedisDb.Get(global.Ctx, fmt.Sprintf("checkCode:%v", l.Mobile)).Result()
	if err != nil {
		glog.Error("获取cookie失败")
		return errors.New("获取cookie失败" + err.Error())
	}

	// 请求头封装
	request.Header.Add("Cookie", cookieStr)
	request.Header.Set("Content-Type", "application/json")
	helper.SetHead(request)

	client := &http.Client{}

	response, _ := client.Do(request)

	// 验证之后的cookie
	str := solveSetCookie(response.Header["Set-Cookie"][0]) + "cmi-user-ticket=zwAYVjg8QMKiN0GNRZuSq08nhd5H-rjRwFjqgA..; " + "secure-key=1e4ad55a-020a-4ee9-b8b4-cdc7c520176d; " + solveSetCookie(response.Header["Set-Cookie"][0]) + "agent_login_img_code=e414dd97a7ed47fe8054f4e698d0d031; " + solveSetCookieS(response.Header["Set-Cookie"][1])

	// 存放
	_, err = global.RedisDb.Set(global.Ctx, fmt.Sprintf("verfiyCode:%v", l.Mobile), str, time.Hour*7*24).Result()
	if err != nil {
		log.Println("存放数据失败: ", err)
		return errors.New("存放数据失败: " + err.Error())
	}

	return nil
}

// 验证验证码 verfiyCode(手机号码和图形验证码)
func (l *Login) VerfiyCode(code string) error {

	// 手机号码加密
	mobileAes := aess.AesMobile(l.Mobile)

	// 替换符号
	mobileAesSwap := strings.ReplaceAll(mobileAes, "=", "%3D")

	// 将加密的手机号码的 ` '=' 转换为 '%3D' `
	request := v1.GetRequest(
		fmt.Sprintf("https://www.114yygh.com/web/common/verify-code/get?_time=%v&mobile=%v&smsKey=LOGIN&code=%v", time.Now().UnixMilli(), mobileAesSwap, code))

	// 获取所需的cookie
	cookieStr, err := global.RedisDb.Get(global.Ctx, fmt.Sprintf("verfiyCode:%v", l.Mobile)).Result()
	if err != nil {
		log.Println("获取数据失败: ", err)
		return errors.New("获取数据失败" + err.Error())
	}

	// 设置用户的头部信息
	request.Header.Add("Cookie", cookieStr)

	request.Header.Set("Content-Type", "application/json")
	helper.SetHead(request)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("响应失败: ", err)
		return errors.New("响应失败: " + err.Error())
	}

	str := solveSetCookie(response.Header["Set-Cookie"][0]) + "cmi-user-ticket=zwAYVjg8QMKiN0GNRZuSq08nhd5H-rjRwFjqgA..; " + "secure-key=1e4ad55a-020a-4ee9-b8b4-cdc7c520176d; " + solveSetCookie(response.Header["Set-Cookie"][0]) + "agent_login_img_code=e414dd97a7ed47fe8054f4e698d0d031; " + solveSetCookieS(response.Header["Set-Cookie"][1])

	// 存放
	_, err = global.RedisDb.Set(global.Ctx, fmt.Sprintf("login:%v", l.Mobile), str, time.Hour*7*24).Result()
	if err != nil {
		log.Println("存放数据失败: ", err)
		return errors.New("存放数据失败: " + err.Error())
	}
	return nil
}

// 实现登录(手机号码和短信验证码)
func (l *Login) Login(code string) error {
	// 实现加密
	mobileAes, codeAes := aess.AesECBPass(l.Mobile, code)
	loginBody := login{
		Code:   codeAes,
		Mobile: mobileAes,
	}
	param, err := json.Marshal(loginBody)
	if err != nil {
		log.Println("json 解析失败: ", err)
		return errors.New("解析失败: " + err.Error())
	}

	request := v1.PostRequest("https://www.114yygh.com/web/login?_time=1663917177873", string(param))

	// 获取所需的cookie
	cookieStr, err := global.RedisDb.Get(global.Ctx, fmt.Sprintf("login:%v", l.Mobile)).Result()
	if err != nil {
		log.Println("获取数据失败: ", err)
		return errors.New("获取数据失败" + err.Error())
	}
	request.Header.Set("Cookie", cookieStr)
	request.Header.Set("Content-Type", "application/json")
	helper.SetHead(request)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println("响应错误: ", err)
		return errors.New("响应失败: " + err.Error())
	}
	str := solveSetCookie(response.Header["Set-Cookie"][0]) + "cmi-user-ticket=zwAYVjg8QMKiN0GNRZuSq08nhd5H-rjRwFjqgA..; " + "secure-key=1e4ad55a-020a-4ee9-b8b4-cdc7c520176d; " + solveSetCookie(response.Header["Set-Cookie"][0]) + "agent_login_img_code=e414dd97a7ed47fe8054f4e698d0d031; " + solveSetCookieS(response.Header["Set-Cookie"][1])
	// 存放
	_, err = global.RedisDb.Set(global.Ctx, fmt.Sprintf("register:%v", l.Mobile), str, time.Hour*7*24).Result()
	if err != nil {
		log.Println("存放数据失败: ", err)
		return errors.New("存放数据失败: " + err.Error())
	}

	return nil
}

//  处理某一个cookie为发起的格式
func solveSetCookie(cookie string) string {
	dd := strings.Split(cookie, ";")
	return dd[0] + "; "
}

func solveSetCookieS(cookie string) string {
	dd := strings.Split(cookie, ";")
	return dd[0]
}
