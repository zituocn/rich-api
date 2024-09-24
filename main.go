package main

import (
	"github.com/zituocn/gow"
	"github.com/zituocn/logx"
	"github.com/zituocn/rich-api/conn"
	"github.com/zituocn/rich-api/router"
)

func init() {
	conn.InitLog()
}

func main() {
	r := gow.Default()
	r.SetAppConfig(gow.GetAppConfig())
	router.APIRouter(r)
	err := r.Run()
	if err != nil {
		logx.Panic(err)
	}
}
