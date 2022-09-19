package login

import (
	aess "appointed-registration/helper/aes"
	"appointed-registration/helper/request"
	"fmt"
	"net/http"
	"time"
)

// type Login struct {
// 	Code   string
// 	Mobile string
// }

// func NewLogin(mobile, code string) *Login {
// 	return &Login{
// 		Mobile: mobile,
// 		Code:   code,
// 	}
// }

// 获取一个验证码时的设置的cookie
func GetImgCode() *http.Response {
	param := request.Params{
		Headers: request.Headers{
			"Accept":          "application/json, text/plain, */*",
			"Accept-Encoding": "gzip, deflate, br",
			"Accept-Language": "zh-CN,zh;q=0.9",
			"Connection":      "keep-alive",
			"Content-Type":    "application/json;charset=UTF-8",
			"Cookie":          "imed_session=GZhzLJ7vt4mX4srOz6cIQFQt2x3ymDlP_5543831; cmi-user-ticket=scbWimlAaYCni9KsJPUF5CdxmhuKDIQqsk4qmw..; secure-key=d9d068ec-badb-4e42-ac64-f636c4c3aaf6; agent_login_img_code=e34b7852b8de47debbcee494d6b7d805; imed_session=K3sf43ayR1ccj4BfaKOfPO59vRjIkDH2_5543876; imed_session_tm=1663163037334",
			"Host":            "www.114yygh.com",
			"Referer":         "https://www.114yygh.com/",
			"Request-Source":  "PC",
			// "sec-ch-ua": "Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105",
			"sec-ch-ua-mobile":   "?0",
			"sec-ch-ua-platform": "Windows",
			"Sec-Fetch-Dest":     "empty",
			"Sec-Fetch-Mode":     "cors",
			"Sec-Fetch-Site":     "same-origin",
			"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
		},
	}
	response, err := request.Get("https://www.114yygh.com/web/img/getImgCode?_time=1663165130980", param)
	if err != nil {
		return nil
	}
	defer response.Response.Body.Close()
	return response.Response
}

// 发送验证码, 通过传递数据, 拿到cookie, 实现登录
func VerfiyCode() *http.Response {
	param := request.Params{
		Headers: request.Headers{
			"Accept":          "application/json, text/plain, */*",
			"Accept-Encoding": "gzip, deflate, br",
			"Accept-Language": "zh-CN,zh;q=0.9",
			"Connection":      "keep-alive",
			"Content-Type":    "application/json;charset=UTF-8",
			"Cookie":          "imed_session=GZhzLJ7vt4mX4srOz6cIQFQt2x3ymDlP_5543831; cmi-user-ticket=scbWimlAaYCni9KsJPUF5CdxmhuKDIQqsk4qmw..; secure-key=d9d068ec-badb-4e42-ac64-f636c4c3aaf6; agent_login_img_code=e34b7852b8de47debbcee494d6b7d805; imed_session=K3sf43ayR1ccj4BfaKOfPO59vRjIkDH2_5543876; imed_session_tm=1663163037334",
			"Host":            "www.114yygh.com",
			"Referer":         "https://www.114yygh.com/",
			"Request-Source":  "PC",
			// "sec-ch-ua": "Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105",
			"sec-ch-ua-mobile":   "?0",
			"sec-ch-ua-platform": "Windows",
			"Sec-Fetch-Dest":     "empty",
			"Sec-Fetch-Mode":     "cors",
			"Sec-Fetch-Site":     "same-origin",
			"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
		},
	}
	response, err := request.Get("https://www.114yygh.com/web/checkcode?_time=1663561706726&code=0531", param)
	if err != nil {
		return nil
	}
	defer response.Response.Body.Close()
	return response.Response
}

// 实现登录
func Login(mobile, code string) *http.Response {
	mobileAes, codeAes := aess.AesECBPass(mobile, code)
	param := request.Params{
		Headers: request.Headers{
			"Accept":          "application/json, text/plain, */*",
			"Accept-Encoding": "gzip, deflate, br",
			"Accept-Language": "zh-CN,zh;q=0.9",
			"Connection":      "keep-alive",
			"Content-Type":    "application/json;charset=UTF-8",
			"Cookie":          "imed_session=KrEkp4rvF9YibXnnLDTLOk6eyg41Kky2_5545193; cmi-user-ticket=oZ4QFC9wi73xFfe7EmhXg031-bhQvRypZ6IT-Q..; secure-key=920bb44e-92b4-4a01-ac34-1cf997f4abe5; imed_session=KrEkp4rvF9YibXnnLDTLOk6eyg41Kky2_5545193; agent_login_img_code=db8f866ec19740e0a5394ed170f3b4f7; imed_session_tm=1663558448900",
			"Host":            "www.114yygh.com",
			"Referer":         "https://www.114yygh.com/",
			"Request-Source":  "PC",
			// "sec-ch-ua": "Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105",
			"sec-ch-ua-mobile":   "?0",
			"sec-ch-ua-platform": "Windows",
			"Sec-Fetch-Dest":     "empty",
			"Sec-Fetch-Mode":     "cors",
			"Sec-Fetch-Site":     "same-origin",
			"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
		},
		Data: map[string]interface{}{
			"mobile": mobileAes,
			"code":   codeAes,
		},
	}
	request, err := request.Post(fmt.Sprintf("https://www.114yygh.com/web/login?_time=%v", time.Now().UnixMilli()), param)
	if err != nil {
		fmt.Println("err: ", err)
		return nil
	}
	return request.Response
}

// func Login(mobile, code string) {

// 	// aesMobile, aesCode := aes.AesECBPass(mobile, code)
// 	param := request.Params{
// 		Headers: request.Headers{
// 			"Accept":          "application/json, text/plain, */*",
// 			"Accept-Encoding": "gzip, deflate, br",
// 			"Accept-Language": "zh-CN,zh;q=0.9",
// 			"Connection":      "keep-alive",
// 			"Content-Type":    "application/json;charset=UTF-8",
// 			"Cookie":          "imed_session=GZhzLJ7vt4mX4srOz6cIQFQt2x3ymDlP_5543831; cmi-user-ticket=scbWimlAaYCni9KsJPUF5CdxmhuKDIQqsk4qmw..; secure-key=d9d068ec-badb-4e42-ac64-f636c4c3aaf6; agent_login_img_code=e34b7852b8de47debbcee494d6b7d805; imed_session=K3sf43ayR1ccj4BfaKOfPO59vRjIkDH2_5543876; imed_session_tm=1663163037334",
// 			"Host":            "www.114yygh.com",
// 			"Referer":         "https://www.114yygh.com/",
// 			"Request-Source":  "PC",
// 			// "sec-ch-ua": "Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105",
// 			"sec-ch-ua-mobile":   "?0",
// 			"sec-ch-ua-platform": "Windows",
// 			"Sec-Fetch-Dest":     "empty",
// 			"Sec-Fetch-Mode":     "cors",
// 			"Sec-Fetch-Site":     "same-origin",
// 			"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
// 		},
// 		Cookies: request.Cookies{
// 			{Name: "imed_session", Value: "ljLHb0HKL0Lh4F7tqsen4G79PO89xrKI_5543885"},
// 			{Name: "agent_login_img_code", Value: "072320623b3e4980acf33f6531a9bfeb"},
// 		},
// 	}
// 	response, err := request.Get("https://www.114yygh.com/web/common/verify-code/get?_time=1663165695513&mobile=XkgVEAv6Hu4tHRtYwsnrWQ%3D%3D&smsKey=LOGIN&code=3430", param)
// 	if err != nil {
// 		return
// 	}
// 	cc, _ := ioutil.ReadAll(response.Response.Body)
// 	fmt.Println(string(cc))

// }
