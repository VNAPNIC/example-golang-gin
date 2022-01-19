package redisUtil

import (
	"healthcare-panel/utils/setting"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var rdb *redis.Client
var ctx = context.Background()

func Setup() {
	rdb = redis.NewClient(&redis.Options{
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
	cmd := rdb.Do(ctx, "PING")
	if err := cmd.Err(); err != nil {
		panic(err)
	}
	log.Println(cmd.Args()...)
}

type Error struct {
	Msg string
}

func (err Error) Error() string { return err.Msg }

// Set a key/value
func Set(key string, data interface{}, time time.Duration) (bool, error) {
	err := rdb.Set(ctx, key, data, time).Err()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Exists check a key exists
func Exists(key string) (bool, error) {
	res, err := rdb.Exists(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return res == 1, nil
}

// Expiration set expiration to key
func Expiration(key string, expiration time.Duration) (bool, error) {
	res, err := rdb.Expire(ctx, key, expiration).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return res, nil
}

// Get get a key
func Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// Delete delete a key
func Delete(key string) (bool, error) {
	res, err := rdb.Del(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return res == 1, nil
}
