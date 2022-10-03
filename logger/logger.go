package logger

import (
	"appointed-registration/global"
	"log"

	"go.uber.org/zap"
)

/**
* 代码描述: 日志的配置文件
* 作者:小大白兔
* 创建时间:2022/10/03 21:34:29
 */

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./logger.log",
		"stderr",
	}
	return cfg.Build()
}

/**
* 代码描述: 日志新建
* 作者:小大白兔
* 创建时间:2022/10/03 21:35:48
 */

func NewInitLogger() {
	logger, err := NewLogger()
	if err != nil {
		log.Println("日志初始化失败: ")
		return
	}

	global.LogSuger = logger.Sugar()
}
