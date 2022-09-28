package global

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	LoginKey = "imed2019imed2019"

	RedisDb *redis.Client
	Ctx     = context.Background()
	// 手机号码
	Phone = ""
)
