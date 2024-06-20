package conn

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/spf13/viper"
	"net"
	"time"
)

var ClickhouseDB *clickhouse.Conn

// GetClickhouseDB .
func GetClickhouseDB() *clickhouse.Conn {
	var chDB *clickhouse.Conn
	if ClickhouseDB == nil {
		drv, err := newClickhouseDB()
		if err != nil || drv == nil {
			panic(fmt.Sprintf("failed to get orm client, error:%s", err.Error()))
			return nil
		}

		chDB = drv
	} else {
		chDB = ClickhouseDB
	}

	return chDB
}

func newClickhouseDB() (*clickhouse.Conn, error) {
	//host: 127.0.0.1
	//port: 8123
	//database: log
	//user: default
	//password: LgGqjwF6I43462T
	//timeout_connect: 2
	//timeout_query: 4
	prefix := fmt.Sprintf("data.clickhouse-log")

	host := viper.GetString(prefix + ".host")
	port := viper.GetInt(prefix + ".port")
	database := viper.GetString(prefix + ".database")
	user := viper.GetString(prefix + ".user")
	password := viper.GetString(prefix + ".password")
	timeout_connect := viper.GetInt(prefix + ".timeout_connect")
	//timeout_query := viper.GetInt(prefix + ".timeout_query")

	path := fmt.Sprintf("%s:%d", host, port)

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{path},
		Auth: clickhouse.Auth{
			Database: database,
			Username: user,
			Password: password,
		},
		DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
			dialCount := 0
			dialCount++
			var d net.Dialer
			return d.DialContext(ctx, "tcp", addr)
		},
		Debug: true,
		Debugf: func(format string, v ...any) {
			fmt.Printf(format+"\n", v...)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:          time.Second * time.Duration(timeout_connect),
		MaxOpenConns:         5,
		MaxIdleConns:         10,
		ConnMaxLifetime:      time.Duration(10) * time.Minute,
		ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
		ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "my-app", Version: "0.1"},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	ClickhouseDB = &conn

	return &conn, nil
}
