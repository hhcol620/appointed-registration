package listdepartment

import (
	"fmt"
	"testing"
)

func Test_GetDepartment(t *testing.T) {
	a, b, err := GetDepartment("120")
	if err != nil {
		fmt.Println("test错误: ", err)
		return
	}
	fmt.Println(a, b)
}
