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
	req.Header.Set("Cookie", "med_session=rBJT3t7HbGuGOkCpgGJgYKCKYaMCgSe5_5545210; imed_session=rBJT3t7HbGuGOkCpgGJgYKCKYaMCgSe5_5545210; cmi-user-ticket=eH5tKKzbpePMlsGGTzYn61_CuvD-fJC28BjhGQ..; secure-key=4d4696a3-00fe-4e05-8fed-1d99838bf890; imed_session=UjsnaxifXHVj4qVNsrF3WL3xeQZIpEQB_5545243; agent_login_img_code=6b35f68e38004d3083047c0c383ee0f4; imed_session_tm=1663573193953")
	// req.Header.Set("Cookie", cookie)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
}
