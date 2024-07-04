package bootstrap

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

func InitLog() {
	level := viper.GetString("log.log_level")
	dir := viper.GetString("log.log_dir")
	prefix := viper.GetString("log.log_prefix")
	maxSize := viper.GetInt("log.max_size")

	// 设置日志格式。
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	lv, err := logrus.ParseLevel(level)
	if err != nil {
		panic("logrus parse level err")
	}
	logrus.SetLevel(lv)
	logrus.SetReportCaller(true) // 打印文件、行号和主调函数。

	now := time.Now()
	// 格式化当前时间为"2006-01-02 15:04:05"的格式
	formattedTime := now.Format(time.DateOnly)
	path := fmt.Sprintf("%s/%s-%s.log", dir, prefix, formattedTime)
	// 实现日志滚动。
	logger := &lumberjack.Logger{
		Filename:   path,    // 日志输出文件路径。
		MaxSize:    maxSize, // 日志文件最大 size(MB)，缺省 100MB。
		MaxBackups: 10,      // 最大过期日志保留的个数。
		MaxAge:     30,      // 保留过期文件的最大时间间隔，单位是天。
		LocalTime:  true,    // 是否使用本地时间来命名备份的日志。
	}

	logrus.SetOutput(logger)

	return
}
