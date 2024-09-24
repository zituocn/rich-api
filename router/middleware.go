package router

import (
	"github.com/zituocn/gow"
	"github.com/zituocn/gow/lib/config"
)

var (
	token = config.DefaultString("auth::token", "6eff526e68eabf54a28e5d136d4eba9c")
)

func Auth() gow.HandlerFunc {
	return func(c *gow.Context) {
		tk := c.GetHeader("token")
		if tk == "" {
			tk = c.GetString("token")
		}
		if tk == "" {
			c.DataJSON(1, "缺少token")
			c.StopRun()
		}
		if tk != token {
			c.DataJSON(1, "鉴权失败")
			c.StopRun()
		}

		c.Next()
	}
}
