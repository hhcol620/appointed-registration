package models

// 用于接口获取参数的数据绑定
type GetImgCode struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
}

// 获取验证码和手机号码
type GetCode struct {
	Code string `form:"code" json:"code" binding:"required"`
}

// 获取某个医院的科室
type GetHostpitalCode struct {
	HostpitalCode string `form:"hostpitalCode" json:"hostpitalCode" binding:"required"`
}

// swagger 接口文档
