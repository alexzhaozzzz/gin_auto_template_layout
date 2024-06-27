package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/server"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/bootstrap"
)

func Run() {
	// 定义一个命令行Flag来指定配置文件路径
	path := pflag.StringP("config", "c", "", "Path to the config file")
	pflag.Parse()

	err := bootstrap.LoadConfig(*path)
	if err != nil {
		panic("load config error: " + err.Error())
	}

	bootstrap.InitLog()

	r := server.InitServer()

	httpHost := viper.GetString("server.http.host")
	httpPort := viper.GetInt("server.http.port")
	if httpHost == "" || httpPort <= 0 {
		panic("get http config err")
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", httpHost, httpPort),
		Handler: r,
	}

	// 用于监听中断信号的通道
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// 开始监听端口
		if err = srv.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			panic(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	fmt.Println("Server is running...")

	<-quit // 阻塞主goroutine，等待信号

	sTime := time.Now().Format(time.DateTime)
	fmt.Println("Shutting down server:", sTime)

	// 给予5秒时间完成已有请求
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown:", err)
	}

	eTime := time.Now().Format(time.DateTime)
	fmt.Println("Server exited:", eTime)

	return
}
