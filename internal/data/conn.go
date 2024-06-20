package data

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var MerchantDbClient *gorm.DB

// GetMerchantDB .
func GetMerchantDB() *gorm.DB {
	var merchantDB *gorm.DB
	if MerchantDbClient == nil {
		drv, err := getNewMerchantDb()
		if err != nil || drv == nil {
			panic(fmt.Sprintf("failed to get orm client, error:%s", err.Error()))
			return nil
		}

		merchantDB = drv
	} else {
		merchantDB = MerchantDbClient
	}

	return merchantDB
}

func GetMerchantDBAndFun() (*gorm.DB, func(), error) {
	var merchantDB *gorm.DB
	if MerchantDbClient == nil {
		drv, err := getNewMerchantDb()
		if err != nil || drv == nil {
			panic(fmt.Sprintf("failed to get orm client, error:%s", err.Error()))
			return nil, nil, err
		}

		merchantDB = drv
	} else {
		merchantDB = MerchantDbClient
	}

	return merchantDB, func() {
		sqlDB, err := merchantDB.DB()
		if err != nil || sqlDB == nil {
			return
		}
		err = sqlDB.Close()
		if err != nil {
			return
		}
	}, nil
}

func getNewMerchantDb() (*gorm.DB, error) {
	prefix := fmt.Sprintf("data.database-merchant")

	dialect := viper.GetString(prefix + ".dialect")
	user := viper.GetString(prefix + ".user")
	password := viper.GetString(prefix + ".password")
	database := viper.GetString(prefix + ".database")
	host := viper.GetString(prefix + ".host")
	port := viper.GetInt(prefix + ".port")
	maxIdle := viper.GetInt(prefix + ".maxIdle")
	maxOpen := viper.GetInt(prefix + ".maxOpen")
	maxLifetime := viper.GetInt(prefix + ".maxLifetime")

	if host == "" || port <= 0 {
		return nil, errors.New("db config is err")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)

	gdb := &gorm.DB{}
	switch dialect {
	case "mysql":
		// 初始化GORM配置
		config := gorm.Config{}

		// 打开数据库连接
		db, err := gorm.Open(mysql.Open(dsn), &config)

		if err != nil {
			return nil, err
		}

		gdb = db
	default:
		return nil, errors.New("dialect not support")
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		panic("connect db server failed.")
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(maxIdle)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxOpen)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)

	MerchantDbClient = gdb

	return gdb, nil
}
