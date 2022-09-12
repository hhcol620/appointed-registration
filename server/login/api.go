package login

import (
	"errors"
	"net/http"
)

type login map[string]interface{}

// 实现用户的登录, 而后获取cookie
func Login() {

}

func (l login) getErWei() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://img.114yygh.com/staticres/web/code_login_wechat.png", nil)
	if err != nil {
		return errors.New("")
	}
	// 添加头部信息
}
