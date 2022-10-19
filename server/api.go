package server

import (
	"appointed-registration/dao"
	"appointed-registration/global"
	"appointed-registration/helper/helper"
	"appointed-registration/models/hospital"
	registered "appointed-registration/server/Registered"
	allhospital "appointed-registration/server/allHospital"
	listdepartment "appointed-registration/server/listDepartment"
	"appointed-registration/server/login"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type Server struct {
	ApiLogin login.Login
	HosCode  listdepartment.HostCode
}

func NewServers(mobile, hosCode string) *Server {
	return &Server{
		ApiLogin: login.Login{
			Mobile: mobile,
		},

		HosCode: listdepartment.HostCode{
			HosCode: hosCode,
		},
	}
}

/**
* 代码描述: 获取图片验证码
* 作者:小大白兔
* 创建时间:2022/09/26 22:12:47
 */
func (s *Server) GetImgCode() error {
	global.LogSuger.Info("GetImgCode 开始执行...")

	response, err := s.ApiLogin.GetImgCode()
	if err != nil {

		global.LogSuger.Errorf("响应失败: " + err.Error())
		return errors.New("响应失败: " + err.Error())

	}

	cc, _ := ioutil.ReadAll(response.Body)

	name := path.Join("mobile", fmt.Sprintf("%v.jpg", s.ApiLogin.Mobile))

	out, _ := os.Create(name)
	defer out.Close()

	_, err = io.Copy(out, bytes.NewReader(cc))
	if err != nil {
		global.LogSuger.Errorf("copys失败: " + err.Error())
		return errors.New("copys失败: " + err.Error())
	}

	global.LogSuger.Info("GetImgCode 执行结束...")

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
		global.LogSuger.Errorf("获取数据失败: " + err.Error())
		return err
	}

	err = s.ApiLogin.VerfiyCode(code)
	if err != nil {
		global.LogSuger.Errorf("获取手机验证码失败")
		return err
	}

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
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		log.Println("登录失败")
		return err

	}

	return nil
}

/**
* 代码描述: 所有医院的获取并且存入数据库
* 作者:小大白兔
* 创建时间:2022/10/01 22:44:31
 */
//  map[code:237 distance:<nil> levelText:二级医院 maintain:false name:北京市昌平区沙河医院 openTimeText:10:15 picture://img.114yygh.com/image/image-003/23177077420475949.png]
func GetAllHostpital() error {

	for i := 1; ; i++ {

		dataStr, err := allhospital.GetAddress(0, "0", i)
		if err != nil {
			return errors.New("错误: " + err.Error())
		}

		// 查询出所有的结果数据
		if strings.Contains(dataStr, "[]") {
			break
		}

		result, err := helper.FormatString(dataStr)
		if err != nil {
			global.LogSuger.Errorf("转换失败: " + err.Error())
			return errors.New("转换失败: " + err.Error())
		}

		arr := []*hospital.Hospital{}
		// fmt.Println(v.(map[string]interface{})["code"])
		// 将获取的到数据存储
		for _, v := range result {
			a := &hospital.Hospital{
				Code:      v.(map[string]interface{})["code"].(string),
				LevelText: v.(map[string]interface{})["levelText"].(string),
				Maintain:  v.(map[string]interface{})["maintain"].(bool),
				Name:      v.(map[string]interface{})["name"].(string),
			}
			arr = append(arr, a)
		}
		dao.DB.Create(&arr)
		time.Sleep(3 * time.Second)
	}

	return nil
}

/**
* 代码描述: 获取医院的科室信息
* 作者:小大白兔
* 创建时间:2022/10/02 17:35:53
 */
func (s *Server) GetDepartment() (map[string]interface{}, error) {

	response, err := s.HosCode.GetDepartmentFront()
	if err != nil {
		log.Println("响应失败: " + err.Error())
		return nil, errors.New("响应失败: " + err.Error())
	}

	// 将数据转换, 并且发送数据位前端可以使用的数据
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("数据读取失败: " + err.Error())
		return nil, errors.New("数据读取失败: " + err.Error())
	}

	result := map[string]interface{}{}

	if err = json.Unmarshal(body, &result); err != nil {
		log.Println("数据解析失败: " + err.Error())
		return nil, errors.New("数据解析失败: " + err.Error())
	}

	return result, nil
}

/**
* 代码描述:获取科室的号码数
* 作者:小大白兔
* 创建时间:2022/10/09 16:14:11
 */

func GetRegister(hostCode, firstDeptCode, secondDeptCode string) error {
	res := registered.RegisterCode{
		HostCode:       hostCode,
		FirstDeptCode:  firstDeptCode,
		SecondDeptCode: secondDeptCode,
	}

	// 获取响应
	response, err := res.Check()
	if err != nil {
		global.LogSuger.Errorf("获取响应失败")
		return errors.New("获取响应失败")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		global.LogSuger.Errorf("数据解析失败: " + err.Error())
		return errors.New("数据解析失败: " + err.Error())
	}

	fmt.Println(string(body))
	return nil
}
