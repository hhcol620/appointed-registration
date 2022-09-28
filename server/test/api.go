package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 测试请求登录界面
func post() {

	params := `{
		"mobile":"18088630924",
		"password":"12345678"
		}`
	ss := []byte(params)
	fmt.Println(ss, params)
	request, err := http.NewRequest("POST", "http://127.0.0.1:8080/u/v1/user/login", bytes.NewBuffer(ss))
	if err != nil {
		fmt.Println("请求失败: ", err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("响应失败: ", err)
		return
	}
	dd, _ := ioutil.ReadAll(response.Body)
	fmt.Println("这是结果", string(dd))

}

//  imed_session=hLsrHr69SmsDJZz6Ttyo7al7j9BT3A5p_5546737;
// cmi-user-ticket=cBNe_5tnP3O8t-jnHGm6ZW1DuTf_yc4n0yXaLg..;

//  imed_session_tm=1664371544677
