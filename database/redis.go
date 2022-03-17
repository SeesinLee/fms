package database

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var DBR *redis.Pool

func InitRedis()*redis.Pool{
	DBR = &redis.Pool{
		Dial: func() (redis.Conn, error) {
				re,err := redis.Dial(viper.GetString("redis.network"),viper.GetString("redis.host"))
				if err != nil {
					return nil,err
				}
				return re,err
		},
	}
	return DBR
}
