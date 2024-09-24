package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/zituocn/gow/lib/config"
	"net"
	"strings"
	"time"
)

var (
	ws = "ws://127.0.0.1:9222"
)

type BaiduService struct {
}

type Result struct {
	Url         string `json:"url"`
	Record      bool   `json:"record"`
	Title       string `json:"title"`
	Datetime    string `json:"datetime"`
	Description string `json:"description"`
	Tips        string `json:"tips"`
}

func (m *BaiduService) CheckURL(keyword string) (ret *Result, err error) {
	if keyword == "" {
		return
	}
	addr := "https://www.baidu.com"
	timeCtx, cancel := context.WithTimeout(getChromeCtx(), time.Second*15)
	defer cancel()

	var body string
	var record bool
	var title, description, recordTime string

	ret = &Result{
		Url:    keyword,
		Record: false,
	}
	waitTime := config.DefaultInt("wait", 1000)
	err = chromedp.Run(timeCtx,
		chromedp.Navigate(addr),
		chromedp.SetValue(`#kw`, keyword, chromedp.ByID),
		chromedp.Click(`su`, chromedp.ByID),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.Sleep(time.Duration(waitTime)*time.Millisecond),
		chromedp.OuterHTML(`body`, &body, chromedp.ByQuery),
	)
	if err != nil {
		return
	}

	if strings.Contains(body, "百度安全验证") {
		err = fmt.Errorf("出现 百度安全验证")
		return
	}

	if strings.Contains(body, "https://wappass.baidu.com/static/captcha") {
		err = fmt.Errorf("出现百度安全验证")
		return
	}

	if strings.Contains(body, "未找到相关结果") || strings.Contains(body, "没有找到该URL") {
		return
	}

	if strings.Contains(body, keyword) {
		record = true
	}

	reader := bytes.NewReader([]byte(body))
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	ss := doc.Find(".c-container").First()
	ss.Find(".c-title").Each(func(i int, s *goquery.Selection) {
		title = strings.TrimSpace(s.Text())
	})
	ss.Find(".c-gap-top-small").Find("span").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			recordTime = strings.TrimSpace(s.Text())
		}
		if i == 1 {
			description = strings.TrimSpace(s.Text())
		}
	})

	if recordTime == "" || description == "" {

		ss.Find(".c-span9").Find("span").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				recordTime = strings.TrimSpace(s.Text())
			}
			if i == 1 {
				description = strings.TrimSpace(s.Text())
			}
		})
	}

	ret.Record = record
	ret.Title = title
	ret.Description = description
	ret.Datetime = recordTime
	if ret.Record {
		ret.Tips = "已收录"
	}

	return
}

func getChromeCtx() context.Context {
	var chromeCtx context.Context
	allowOpts := chromedp.DefaultExecAllocatorOptions[:]
	allowOpts = append(allowOpts,
		chromedp.Flag("headless", true),                                 //关掉浏览器窗口
		chromedp.Flag("enable-automation", false),                       //是否显示自动化测试标识
		chromedp.Flag("disable-blink-features", "AutomationControlled"), //禁止掉chrome标识
		chromedp.Flag("blink-settings", "imageEnable=true"),             //不渲染图片
		chromedp.Flag("ignore-certificate-errors", true),                //忽略错误
		chromedp.Flag("disable-web-security", true),                     //禁用网络安全标志
		chromedp.Flag("disable-gpu", true),
		chromedp.DisableGPU,
	)

	if checkChromePort() {
		c, _ := chromedp.NewRemoteAllocator(context.Background(), ws)
		chromeCtx, _ = chromedp.NewContext(c)
	} else {
		c, _ := chromedp.NewExecAllocator(context.Background(), allowOpts...)
		chromeCtx, _ = chromedp.NewContext(c)
	}

	return chromeCtx
}

func checkChromePort() bool {
	addr := net.JoinHostPort("", "9222")
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
