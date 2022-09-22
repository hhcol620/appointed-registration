package listdepartment

import (
	"appointed-registration/helper"
	v1 "appointed-registration/helper/request"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 获取一个医院的所有科室

func GetDepartment(hosCode string) (map[string][]ValueDepartment, map[string][]ValueDepartment, error) {

	// request, err := http.NewRequest("GET",
	// 	fmt.Sprintf("https://www.114yygh.com/web/department/hos/list?_time=%v&hosCode=%v",
	// 		time.Now().UnixMilli(), hosCode), nil)

	// if err != nil {
	// 	log.Println("err:", err)
	// 	return nil, nil, errors.New("err: " + err.Error())
	// }
	request := v1.GetRequest(fmt.Sprintf("https://www.114yygh.com/web/department/hos/list?_time=%v&hosCode=%v",
		time.Now().UnixMilli(), hosCode))

	helper.SetHead(request)
	// defer request.Body.Close()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("err:", err)
		return nil, nil, errors.New("err: " + err.Error())
	}

	cc, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("解析错误: ", err)
		return nil, nil, errors.New("解析出错: err" + err.Error())
	}

	result := map[string]interface{}{}

	err = json.Unmarshal(cc, &result)

	arr := map[KeyDepartment][]ValueDepartment{}

	for _, v := range result["data"].(map[string]interface{})["list"].([]interface{}) {

		// 遍历所有的数据子列表
		departmentList := []ValueDepartment{}
		for _, departments := range v.(map[string]interface{})["subList"].([]interface{}) {

			department := ValueDepartment{
				Code:      departments.(map[string]interface{})["code"].(string),
				Dept1Code: departments.(map[string]interface{})["dept1Code"].(string),
				HotDept:   departments.(map[string]interface{})["hotDept"].(bool),
				Name:      departments.(map[string]interface{})["name"].(string),
			}

			departmentList = append(departmentList, department)
		}

		arr[KeyDepartment{Code: v.(map[string]interface{})["code"].(string), Name: v.(map[string]interface{})["name"].(string)}] = departmentList
	}

	// 解析数据返回
	departmentCode, departmentName := formatArry(arr)
	return departmentCode, departmentName, nil
}

// 格式化数据
func formatArry(arr map[KeyDepartment][]ValueDepartment) (map[string][]ValueDepartment, map[string][]ValueDepartment) {
	departmentCode, departmentName := map[string][]ValueDepartment{}, map[string][]ValueDepartment{}
	for k, v := range arr {
		departmentCode[k.Code], departmentName[k.Name] = v, v
	}
	return departmentCode, departmentName
}
