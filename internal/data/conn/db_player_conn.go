package conn

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var PlayerDBClient *gorm.DB

// GetPlayerDB .
func GetPlayerDB() *gorm.DB {
	var PlayerDB *gorm.DB
	if PlayerDBClient == nil {
		drv, err := newPlayerDB()
		if err != nil || drv == nil {
			panic(fmt.Sprintf("failed to get orm client, error:%s", err.Error()))
			return nil
		}

		PlayerDB = drv
	} else {
		PlayerDB = PlayerDBClient
	}

	return PlayerDB
}

func GetPlayerDBAndFun() (*gorm.DB, func(), error) {
	var PlayerDB *gorm.DB
	if PlayerDBClient == nil {
		drv, err := newPlayerDB()
		if err != nil || drv == nil {
			panic(fmt.Sprintf("failed to get orm client, error:%s", err.Error()))
			return nil, nil, err
		}

		PlayerDB = drv
	} else {
		PlayerDB = PlayerDBClient
	}

	return PlayerDB, func() {
		sqlDB, err := PlayerDB.DB()
		if err != nil || sqlDB == nil {
			return
		}
		err = sqlDB.Close()
		if err != nil {
			return
		}
	}, nil
}

func newPlayerDB() (*gorm.DB, error) {
	prefix := fmt.Sprintf("data.mysql-player")

	user := viper.GetString(prefix + ".user")
	password := viper.GetString(prefix + ".password")
	database := viper.GetString(prefix + ".database")
	host := viper.GetString(prefix + ".host")
	port := viper.GetInt(prefix + ".port")
	maxIdle := viper.GetInt(prefix + ".max_idle")
	maxOpen := viper.GetInt(prefix + ".max_open")
	maxLifetime := viper.GetInt(prefix + ".max_life_time")

	if host == "" || port <= 0 {
		return nil, errors.New("db config is err")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)
	// 初始化GORM配置
	config := gorm.Config{}

	// 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &config)

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("connect db server failed.")
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(maxIdle)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxOpen)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)

	PlayerDBClient = db

	return db, nil
}
