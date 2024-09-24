package conn

import (
	"github.com/zituocn/gow"
	"github.com/zituocn/gow/lib/config"
	"github.com/zituocn/logx"
	"io"
	"os"
)

// InitLog init log
func InitLog() {
	runMode := config.DefaultString("run_mode", "dev")
	maxDay := config.DefaultInt("log_storage_max_day", 7)
	if runMode == gow.ProdMode {
		logx.SetWriter(io.MultiWriter(
			os.Stdout,
			logx.NewFileWriter(logx.FileOptions{
				StorageType: logx.StorageTypeDay,
				MaxDay:      maxDay,
				Dir:         "./logs",
				Prefix:      "web",
			}))).SetColor(false)
	} else {
		logx.SetColor(true)
	}
	logx.Info("-------------------------------------------")
	logx.Info("开始启动 RichAPI 系统 ...")
	logx.Info("-------------------------------------------")
}
