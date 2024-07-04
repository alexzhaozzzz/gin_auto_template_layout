package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/server"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/bootstrap"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/colorx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/util"
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

	local := util.GetLocalIP()
	fmt.Println(colorx.Green("\r\nServer run at:"))
	fmt.Printf("-  Local:   %s://localhost:%d/ \r\n", "http", httpPort)
	fmt.Printf("-  Network: %s://%s:%d/ \r\n", "http", local, httpPort)
	fmt.Println(colorx.Green("Swagger run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/swagger/index.html \r\n", httpPort)
	fmt.Printf("-  Network: %s://%s:%d/swagger/index.html \r\n", "http", local, httpPort)
	fmt.Printf("\r\n %s Enter Control + C Shutdown Server \r\n", time.Now().Format(time.DateTime))

	<-quit // 阻塞主goroutine，等待信号

	fmt.Printf("%s Shutdown Server ... \r\n", time.Now().Format(time.DateTime))

	// 给予5秒时间完成已有请求
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		fmt.Printf("%s Server Forced To Shutdown ... \r\n", time.Now().Format(time.DateTime))
		return
	}

	fmt.Printf("%s Server Exited ... \r\n", time.Now().Format(time.DateTime))

	return
}
