package models

// 用于接口获取参数的数据绑定
type GetImgCode struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
}

// 获取输入的图片验证码
type GetMobile struct {
	Mobile string
	Code   string
}

// swagger 接口文档
