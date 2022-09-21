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

func SetHeads(req *http.Request, cookie string) {
	req.Header.Set("Request-Source", "PC")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Encodingt", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "www.114yygh.com")
	req.Header.Add("Referer", "https://www.114yygh.com/")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", cookie)
	// req.Header.Set("Cookie", cookie)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
}

func SetRegisteredHead(req *http.Request) {
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Request-Source", "PC")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "www.114yygh.com")
	req.Header.Add("Referer", "https://www.114yygh.com/hospital/120/a660294efe4daaf0bcbff7d69225ce5b/200000909/source")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "imed_session=y93DYH6MMN5gEfTDeJ9O7eoAbODDAj6p_5545804; imed_session=y93DYH6MMN5gEfTDeJ9O7eoAbODDAj6p_5545804; secure-key=2c84c72d-6973-4ce4-8f75-8b46dc723c28; imed_session=y93DYH6MMN5gEfTDeJ9O7eoAbODDAj6p_5545804; cmi-user-ticket=WiOMolmfpHDSXFAKB4uU_SpN6q4xIkIuQOrQGg..; agent_login_img_code=7b6485b2b2a241c9b243b43ffd38de21; imed_session_tm=1663741419468")
	// req.Header.Set("Cookie", cookie)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
}
