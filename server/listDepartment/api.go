package listdepartment

import (
	"appointed-registration/global"
	"appointed-registration/helper"
	v1 "appointed-registration/helper/request"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type HostCode struct {
	HosCode string
}

// 格式化数据
func formatArry(arr map[KeyDepartment][]ValueDepartment) (map[string][]ValueDepartment, map[string][]ValueDepartment) {

	departmentCode, departmentName := map[string][]ValueDepartment{}, map[string][]ValueDepartment{}

	for k, v := range arr {
		departmentCode[k.Code], departmentName[k.Name] = v, v

	}

	return departmentCode, departmentName
}

/**
* 代码描述:科室获取
* 作者:小大白兔
* 创建时间:2022/10/24 16:58:55
 */
func (h *HostCode) GetDepartmentFront() (*http.Response, error) {

	global.LogSuger.Info("GetDepartmentFront接口数据获取开始...")

	request := v1.GetRequest(fmt.Sprintf("https://www.114yygh.com/web/department/hos/list?_time=%v&hosCode=%v",
		time.Now().UnixMilli(), h.HosCode))

	helper.SetHead(request)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		global.LogSuger.Errorf("请求成功: " + err.Error())
		return nil, errors.New("err: " + err.Error())
	}

	global.LogSuger.Info("GetDepartmentFront接口数据获取结束...")

	return response, nil
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
