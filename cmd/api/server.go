package api

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/server"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/bootstrap"
)

func Run() {
	// 定义一个命令行Flag来指定配置文件路径
	path := pflag.StringP("config", "c", "", "Path to the config file")
	pflag.Parse()

	err := bootstrap.LoadConfig(path)
	if err != nil {
		panic("load config error: " + err.Error())
	}

	bootstrap.InitLog()

	r := server.InitServer()

	httpHost := viper.GetString("server.http.host")
	httpPort := viper.GetInt("server.http.port")

	if httpHost == "" || httpPort <= 0 {
		panic("http path error: ")
	}
	err = r.Run(fmt.Sprintf("%s:%d", httpHost, httpPort))
	if err != nil {
		panic("http start error: " + err.Error())
	}

	return
}
