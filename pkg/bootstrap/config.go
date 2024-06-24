package bootstrap

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var PermConfigList []PermissionConfig

type PermissionConfig struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

// LoadConfig ...
func LoadConfig(flagPath *string) error {
	// 根据Flag的值设置Viper的配置文件路径
	if *flagPath != "" {
		viper.SetConfigFile(*flagPath)
	} else {
		// 如果没有提供Flag，则使用默认的配置文件路径
		viper.SetConfigName("config")    // name of config file (without extension)
		viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath("./configs") // path to look for the config file in
	}

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	LoadPermConfigFromJson()

	return nil
}

func LoadPermConfigFromJson() {
	// 读取JSON文件
	jsonFile, err := os.ReadFile("./configs/perm_config.json")
	if err != nil {
		panic(fmt.Errorf("fatal Open Permission Config File: %w", err))
	}

	fmt.Println(string(jsonFile))

	// 解析JSON数据到结构体
	config := make([]PermissionConfig, 0)
	if err := json.Unmarshal(jsonFile, &config); err != nil {
		panic(fmt.Errorf("fatal Unmarshal Permission Config File: %w", err))
	}

	PermConfigList = config

	return
}
