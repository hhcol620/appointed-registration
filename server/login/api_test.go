package login

import (
	"fmt"
	"testing"
)

func Test_GetImgCode(t *testing.T) {
	str := getImgCode("18088630924")
	fmt.Println(str)
}

func Test_CheckCode(t *testing.T) {
	CheckCode("0653", "tqBih5MhOSWMKEJVGMkdn2kRV0VyaCzJ_5546196; secure-key=6f8a7fb9-cb6e-49b1-a865-4a37f3bb35a8; cmi-user-ticket=vauqiml33O-y3aPzwqRwZKAd-xuZPbDqvlFrWw..;; secure-key=35baa7d5-a68c-4c0b-b953-33ac7dcb478f; imed_session=qeJqzeKODVMElF9UkZrE3evrO6T9XKq2_5546721; agent_login_img_code=dcec85459dd948068e8766bad22e473f; imed_session_tm=1664016308344;")
}

// func Test_VerfiyCode(t *testing.T) {

// }

// func Test_Login(t *testing.T) {
// 	dd, _ := Login("18088630924", "673613")
// 	cc, _ := ioutil.ReadAll(dd.Body)
// 	fmt.Println(dd.StatusCode)
// 	fmt.Println(string(cc))
// }

// func Test_solveSetCookie(t *testing.T) {
// 	// ss := solveSetCookie("secure-key=eb596aa1-1664-4dd4-aeae-2516fb334da3; Expires=Thu, 22-Sep-2022 14:44:14 GMT; Path=/; HttpOnly")
// 	// fmt.Println(ss)
// }
