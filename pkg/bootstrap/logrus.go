package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

func InitLog() *logrus.Logger {
	// 初始化Logger
	logger := logrus.New()

	lv := viper.GetInt32("log.log_level")
	logger.SetLevel(logrus.Level(lv))

	// 定义日志文件的基础路径，不包含日期部分
	path := viper.GetString("log.log_dir")
	// 创建并添加DailyRotateHook
	dailyRotateHook := &DailyRotateHook{
		Logger:      logger,
		LogFilePath: path,
	}
	logger.AddHook(dailyRotateHook)

	return logger
}

// DailyRotateHook 自定义的日志按天分割Hook
type DailyRotateHook struct {
	*logrus.Logger
	LogFilePath string
}

// Fire 实现logrus.Hook接口的Fire方法
func (hook *DailyRotateHook) Fire(entry *logrus.Entry) error {
	// 根据日期构建日志文件路径
	currentTime := time.Now()
	logFileName := fmt.Sprintf("%s-%d-%02d-%02d.log", hook.LogFilePath, currentTime.Year(), currentTime.Month(), currentTime.Day())

	// 确保日志目录存在
	if err := os.MkdirAll(filepath.Dir(logFileName), 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	// 打开或创建日志文件
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// 写入日志
	data, err := entry.Bytes()
	if err != nil {
		return fmt.Errorf("get entry Bytes err: %v", err)
	}

	if _, err := entry.WriterLevel(logrus.DebugLevel).Write(data); err != nil {
		return fmt.Errorf("failed to write log entry: %v", err)
	}

	return nil
}

// Levels 实现logrus.Hook接口的Levels方法
func (hook *DailyRotateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
