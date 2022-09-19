package login

import (
	"fmt"
	"io/ioutil"
	"testing"
)

// func Test_Loin(t *testing.T) {
// 	Login("18088630924", "288381")
// }
// var login = NewLogin("18088630924")

func Test_GetImgCode(t *testing.T) {
	res := GetImgCode()
	dd, _ := ioutil.ReadAll(res.Body)
	fmt.Println(dd)
}
func Test_VerfiyCode(t *testing.T) {
	res := VerfiyCode()
	fmt.Println(res.StatusCode)
	re, _ := ioutil.ReadAll(res.Body)
	fmt.Println(re)
}
func Test_Login(t *testing.T) {
	res := Login("18088630924", "020096")
	re, _ := ioutil.ReadAll(res.Body)
	fmt.Println(re)
}
