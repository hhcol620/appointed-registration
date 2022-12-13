package global

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var (
	LoginKey = "imed2019imed2019" // 加密方式

	RedisDb  *redis.Client
	LogSuger *zap.SugaredLogger

	Ctx = context.Background()
	// 手机号码
	Phone = ""

	// 登录获取的获取的参数
	LoginCookie = ""
)
