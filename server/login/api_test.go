package login

import (
	"fmt"
	"testing"
)

var l = &Login{
	Mobile: "18088630924",
}

func Test_GetImgCode(t *testing.T) {
	// var l = &Login{
	// 	Mobile: "18088630924",
	// }
	str, err := l.GetImgCode()
	if err != nil {
		return
	}
	fmt.Println(str)
}

// func Test_CheckCode(t *testing.T) {
// 	str, err := l.CheckCode("1426", "tqBih5MhOSWMKEJVGMkdn2kRV0VyaCzJ_5546196;secure-key=69238e52-ae81-406e-b609-ebe105afc6be;cmi-user-ticket=vauqiml33O-y3aPzwqRwZKAd-xuZPbDqvlFrWw..;secure-key=35baa7d5-a68c-4c0b-b953-33ac7dcb478f;imed_session=smhtFRv3IhLV79R1x5Ogm3ogPD8Id7ch_5548190;agent_login_img_code=b62c7bb5b47749f7a6045a8992beeeb1;imed_session_tm=1664457221998")
// 	fmt.Println(str)
// 	fmt.Println(err)
// }
// func Test_VerfiyCode(t *testing.T) {
// 	l.VerfiyCode("7689", fmt.Sprintf("imed_session=EMPmZ4Z8OYBq6gugmVgzkEwVd6JOAa7D_5548184; "+
// 		"cmi-user-ticket=zwAYVjg8QMKiN0GNRZuSq08nhd5H-rjRwFjqgA..; "+
// 		"secure-key=35baa7d5-a68c-4c0b-b953-33ac7dcb478f; "+
// 		"imed_session=c5AjZZj69xjenHQVLsoqkhEx7950lJ7b_5548188;  "+
// 		"agent_login_img_code=e414dd97a7ed47fe8054f4e698d0d031; "+
// 		"imed_session_tm=1664456802525"))
// }

func Test_Login(t *testing.T) {
	// l.VerfiyCode("757770", "imed_session=v2lKcALkLWoEW48JmISag1vgrZ6uw9Qa_5548201; secure-key=22352861-24d3-4cfe-897c-f6cda9e051f1; imed_session=v2lKcALkLWoEW48JmISag1vgrZ6uw9Qa_5548201; cmi-user-ticket=T-8DDGC8XUFEnCU_Q0DjTRm-LAaVjDmCiLlcWA..; agent_login_img_code=51b8a26e41844636ad45aa9150769bf9; imed_session_tm=1664460448051")
}
