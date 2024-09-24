package router

import (
	"github.com/zituocn/gow"
	"github.com/zituocn/rich-api/handler"
)

func APIRouter(r *gow.Engine) {
	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		auth.Use(Auth())
		{
			auth.GET("/baidu/check", handler.BaiduCheck)
		}
	}
}
