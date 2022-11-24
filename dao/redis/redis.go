package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func Init() (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: "suxianjin123",                  // 密码
		DB:       viper.GetInt("redis.db"),        // 数据库
		PoolSize: viper.GetInt("redis.pool_size"), // 连接池大小
	})
	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func Close() {
	client.Close()
}
