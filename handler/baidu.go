package handler

import (
	"fmt"
	"github.com/zituocn/gow"
	"github.com/zituocn/rich-api/service"
	"strings"
)

func BaiduCheck(c *gow.Context) {
	keyword := c.GetString("url")
	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		c.DataJSON(1, "请传入URL参数")
		return
	}
	bs := new(service.BaiduService)
	ret, err := bs.CheckURL(keyword)
	if err != nil {
		c.DataJSON(1, fmt.Sprintf("查询错误:%v", err))
		return
	}

	c.DataJSON(ret)
}
