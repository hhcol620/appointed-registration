package helper

import "net/http"

// 构造请求的头部
func SetHead(req *http.Request) {
	req.Header.Set("Request-Source", "PC")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Encodingt", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "www.114yygh.com")
	req.Header.Add("Referer", "https://www.114yygh.com/")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "imed_session=0B19QDlODoTooHBRi2GJWbLhCeJPZoyc_5543232; imed_session=8KRV3NGmTsB069EJYTOenqVZSrlGBkzW_5543228; secure-key=4d19ee26-38e6-45de-8aa0-e8497308c769; cmi-user-ticket=9nYFi5NiEogFZOChEYvxpNG4MTQUyUY_T22UuQ..; agent_login_img_code=9852e090938545ae900c20c0c585760d; imed_session=0B19QDlODoTooHBRi2GJWbLhCeJPZoyc_5543232; imed_session_tm=1662970188078")
	// req.Header.Set("Cookie", cookie)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
}
