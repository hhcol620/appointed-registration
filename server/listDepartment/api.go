package listdepartment

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 获取一个医院的所有科室
func GetDepartment(hosCode string) error {
	client := &http.Client{}

	request, err :=
		http.NewRequest("GET",
			fmt.Sprintf("https://www.114yygh.com/web/department/hos/list?_time=%v&hosCode=%v", time.Now().UnixMilli(), hosCode), nil)
	if err != nil {
		log.Println("err:", err)
		return nil
	}
	req, err := client.Do(request)

}
