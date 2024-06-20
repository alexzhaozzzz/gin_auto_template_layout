package conn

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var MerchantRDbClient *redis.Client

// GetMerchantRDb .
func GetMerchantRDb() *redis.Client {
	var merchantRDB *redis.Client
	if MerchantRDbClient == nil {
		rdb, err := newMerchantRedisCli()
		if err != nil || rdb == nil {
			panic(fmt.Sprintf("failed to get orm client, error:%s", err.Error()))
			return nil
		}

		merchantRDB = rdb
	} else {
		merchantRDB = MerchantRDbClient
	}

	return merchantRDB
}

func newMerchantRedisCli() (*redis.Client, error) {
	prefix := fmt.Sprintf("data.redis-merchant")
	addr := viper.GetString(prefix + ".addr")
	password := viper.GetString(prefix + ".password")
	database := viper.GetInt(prefix + ".database")
	maxActive := viper.GetInt(prefix + ".max_active")
	//idleTimeout := time.Duration(viper.GetInt(prefix+".idleTimeout")) * time.Second

	rdb := redis.NewClient(&redis.Options{
		Addr:       addr,
		Password:   password,
		DB:         database,
		MaxRetries: 3,
		PoolSize:   maxActive,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	MerchantRDbClient = rdb

	return rdb, nil
}
