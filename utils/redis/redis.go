package redisUtil

import (
	"log"
	"serverhealthcarepanel/utils/setting"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var RedisConn *redis.Client
var ctx = context.Background()

func Setup() {
	RedisConn = redis.NewClient(&redis.Options{
		Addr:         setting.RedisSetting.Host,
		Password:     setting.RedisSetting.Password, // no password set
		DB:           setting.RedisSetting.DB,       // use default DB
		MaxRetries:   setting.RedisSetting.MaxActive,
		MinIdleConns: setting.RedisSetting.MaxIdle,
		IdleTimeout:  setting.RedisSetting.IdleTimeout,
	})

	log.Printf("[info] Redis connected %s DB: %d", setting.RedisSetting.Host, setting.RedisSetting.DB)
	TestConnection()
}

func TestConnection() {
	cmd := RedisConn.Do(ctx, "PING")
	if err := cmd.Err(); err != nil {
		panic(err)
	}
	log.Println(cmd.Args()...)
}
