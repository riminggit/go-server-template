package Redis

import (
	"context"
	"encoding/json"
	"fmt"
	"go-server-template/config"
	"time"
	log "go-server-template/pkg/log"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var ctx = context.Background()

//redis连接
func InitRedis() *redis.Client {
	config := projectConfig.AppConfig
	Rdb = redis.NewClient(&redis.Options{
		Addr:         config.RedisConfig.HOST + ":" + config.RedisConfig.PORT,
		Password:     config.RedisConfig.PASSWORD,
		DialTimeout:  time.Duration(config.RedisConfig.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.RedisConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.RedisConfig.WriteTimeout) * time.Second,
		PoolSize:     config.RedisConfig.PoolSize,
		PoolTimeout:  time.Duration(config.RedisConfig.PoolTimeout) * time.Second,
		DB:           0, // use default DB
	})

	// ping一下检查是否连通
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		log.Error(err)
	}
	// PONG
	fmt.Println("redis 连接成功")

	return Rdb
}

func SetValue(key string, value interface{}, expiration time.Duration) error {

	data, jsonErr := json.Marshal(value)
	if jsonErr != nil {
		return jsonErr
	}

	// 0 意味着没有过期时间
	err := Rdb.Set(ctx, key, data, expiration*time.Second).Err()
	if err != nil {
		log.Error(err)
	}
	return err
}

func GetValue(key string) string {
	val, err := Rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		log.Error("key对应的值不存在")
	} else if err != nil {
		log.Error(err)
	}

	return val
}

func DelValue(key string) error {
	err := Rdb.Del(ctx, key).Err()
	if err != nil {
		log.Error(err)
	}
	return err
}

func PrefixDel(prefix string) {
	iter := Rdb.Scan(ctx, 0, prefix+"*", 0).Iterator()
	for iter.Next(ctx) {
		err := Rdb.Del(ctx, iter.Val()).Err()
		if err != nil {
			panic(err)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

func PrefixGetValue(prefix string) string {
	val, err := Rdb.Get(ctx, prefix+"*").Result()

	if err == redis.Nil {
		log.Error("key对应的值不存在")
	} else if err != nil {
		log.Error(err)
	}

	return val
}
