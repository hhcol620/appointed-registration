package models

// 用于接口获取参数的数据绑定
type GetImgCode struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
}

// 获取验证码和手机号码
type GetCode struct {
	Code string
}

// swagger 接口文档
