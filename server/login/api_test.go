package login

import (
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
	CheckCode("0844")
}
