# RICH-API

baidu收录检测API

### 一、特性

* 使用chromedp检测某地址是否被baidu收录;
* HTTP API化了，可以部署在任意操作系统下;

### 二、安装方式

拉取代码

```shell
git clone github.com/zituocn/rich-api
```

docker 方式跑 chromedp

```shell
docker pull chromedp/headless-shell
```

监听在本机的9222端口

```shell
docker run -itd --restart=always --name chromedp -p 9222:9222  chromedp/headless-shell
```


```
cd rich-adpi
go run .
```

### 三、配置文件：
确保 conf/app.conf的配置文件内容


```ini
app_name = rich-api

# 系统运行的模式，分别是：prod和dev
#   dev 表示开始模式
#   prod 表示生产模式，请确保正式环境运行在此模式下
run_mode = prod

# web服务运行的端口号
#   也可以写成：127.0.0.1:5566
#   :5566 等形式
http_addr = 5566
auto_render = false
session_on = false

# 是否打开gzip
gzip_on = true

# 是否忽略路由中的大小写，如果不忽略，会区分URL中的大小写
ignore_case = true

# Go语言html模板的符号，不能修改
template_left = <%
template_right = %>

# 日志最长存放时间(天)
log_storage_max_day = 30

wait = 1000


[auth]
# API请求时的鉴权token
token = 6eff526e68eabf54a28e5d136d4eba9c
```



### 三、第三方库

* gow
* chromedp
* logx