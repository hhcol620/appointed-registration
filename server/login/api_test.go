package login

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

// func Test_Loin(t *testing.T) {
// 	Login("18088630924", "288381")
// }
// var login = NewLogin("18088630924")

func Test_GetImgCode(t *testing.T) {
	getImgCode()

}

func Test_CheckCode(t *testing.T) {
	CheckCode("0653")
}

func Test_Login(t *testing.T) {
	response, err := Login("18088630924", "926955")

	if err != nil {
		fmt.Println("测试错误: ", err)
		return
	}

	aa, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	fmt.Println(string(aa))
}

func Test_solveSetCookie(t *testing.T) {
	ss := solveSetCookie("secure-key=eb596aa1-1664-4dd4-aeae-2516fb334da3; Expires=Thu, 22-Sep-2022 14:44:14 GMT; Path=/; HttpOnly")
	fmt.Println(ss)
}
