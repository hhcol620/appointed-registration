package server

import (
	"appointed-registration/server/login"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Server struct {
	ApiLogin login.Login
}

func NewServers(apiLogin login.Login) *Server {
	return &Server{
		ApiLogin: apiLogin,
	}
}

/**
* 代码描述: 获取图片验证码
* 作者:小大白兔
* 创建时间:2022/09/26 22:12:47
 */
func (s *Server) GetImgCode() error {
	response, err := s.ApiLogin.GetImgCode()
	if err != nil {
		log.Println("响应失败: ", err)
		return errors.New("响应失败: " + err.Error())
	}

	cc, _ := ioutil.ReadAll(response.Body)

	out, _ := os.Create(fmt.Sprintf("%v.jpg", s.ApiLogin.Mobile))
	defer out.Close()

	_, err = io.Copy(out, bytes.NewReader(cc))
	if err != nil {
		log.Println("copy失败: ", err)
		return errors.New("copys失败: " + err.Error())
	}

	return nil
}

/**
* 代码描述: 实现发验证码
* 作者:小大白兔
* 创建时间:2022/09/28 20:07:24
 */
func (s *Server) GetMobileCode(code string) error {

	err := s.ApiLogin.CheckCode(code)
	if err != nil {
		log.Println("验证图片验证码失败: ", err)
		return err
	}

	err = s.ApiLogin.VerfiyCode(code)
	if err != nil {
		log.Println("验证图片验证码失败: ", err)
		return err
	}

	// 返回没有错误信息
	return nil
}

/**
* 代码描述: 实现登录
* 作者:小大白兔
* 创建时间:2022/09/28 21:05:47
 */
func (s *Server) GetLogin(code string) error {

	// 实现登录
	if err := s.ApiLogin.Login(code); err != nil {

		log.Println("登录失败")
		return err

	}

	return nil
}
