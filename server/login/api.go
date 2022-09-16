package login

import (
	"appointed-registration/helper/request"
	"fmt"
)

// 获取一个验证码时的设置的cookie
func getImgCode() {
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
		return
	}
	defer response.Response.Body.Close()

	fmt.Println(response.Response.Cookies())
}

func verfiyCode(mobile, code string) {

}
func Login(code string) {

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
