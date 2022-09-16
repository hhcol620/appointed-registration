package request

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Headers map[string]string

type Cookies []*http.Cookie

type Data map[string]interface{}

// params
type Params struct {
	Data    Data
	Headers Headers
	Cookies Cookies
}

type responses struct {
	Response *http.Response
	Body     string
}

func request(method string, url string, params Params) (r responses, e error) {
	r, e = do(method, url, params)
	if e != nil {
		return
	}

	return
}

func do(method string, url string, params Params) (r responses, e error) {
	request, e := http.NewRequest(method, url, nil)
	// header
	addHeaders(request, params.Headers)
	// cookie
	addCookies(request, params.Cookies)
	// 参数
	addData(request, params.Data)
	r.Response, e = http.DefaultClient.Do(request)
<<<<<<< HEAD
	// defer r.Response.Body.Close()
=======
	defer r.Response.Body.Close()
>>>>>>> 99cc4cd... 封装请求方法
	if e != nil {
		return
	}
	unCoding(&r)
	return
}

func addHeaders(request *http.Request, headers Headers) {
	for k, v := range headers {
		request.Header.Add(k, v)
	}
}

// 添加cookie
func addCookies(request *http.Request, cookies Cookies) {
	for _, v := range cookies {
		request.AddCookie(v)
	}
}

// 添加参数
func addData(request *http.Request, data Data) {
	query := request.URL.Query()
	for k, v := range data {
		query.Add(k, fmt.Sprint(v))
	}
	request.URL.RawQuery = query.Encode()
}

func unCoding(r *responses) {
	if r.Response.StatusCode == 200 {
		switch r.Response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(r.Response.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)
				if err != nil && err != io.EOF {
					panic(err)
				}
				if n == 0 {
					break
				}
				r.Body += string(buf)
			}
		default:
			bodyByte, _ := ioutil.ReadAll(r.Response.Body)
			r.Body = string(bodyByte)
		}
	} else {
		bodyByte, _ := ioutil.ReadAll(r.Response.Body)
		r.Body = string(bodyByte)
	}
}

// request
func Request(method string, url string, params Params) (responses, error) {
	return request(method, url, params)
}

// get
func Get(url string, params Params) (responses, error) {
	return request("GET", url, params)
}

// post
func Post(url string, params Params) (responses, error) {
	return request("POST", url, params)
}

// put
func Put(url string, params Params) (responses, error) {
	return request("PUT", url, params)
}

// delete
func Delete(url string, params Params) (responses, error) {
	return request("DELETE", url, params)
}

// func Get(url string) *http.Response {
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		log.Println("get请求失败: ", err)
// 		return nil
// 	}

// 	helper.SetHead(req)

// 	client := &http.Client{}

// 	response, err := client.Do(req)
// 	if err != nil {
// 		log.Println("响应失败: ", err)
// 		return nil
// 	}
// 	defer response.Body.Close()

// 	return response
// }

// func Post(url string, body io.Reader) *http.Response {

// }
