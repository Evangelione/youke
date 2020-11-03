package infra

import (
	"context"
	"fmt"
	"time"
	"yk/internal/app/robot"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var Redis *redis.Client

func ConnRedis(config RedisConf) {
	fmt.Println("init redis...")
	// Read config
	address := config.Address

	// Connect
	Redis = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("init redis success!")
}

func SetRedisValue(key string, value interface{}, expire time.Duration) error {
	if err := Redis.Set(ctx, key, value, expire).Err(); err != nil {
		return err
	}
	return nil
}

func GetRedisValue(key string) string {
	result, err := Redis.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		robot.Logger().Error("redis：" + key + "获取失败")
		return ""
	}
	return result
}

func AddList(key string, value interface{}) error {
	if err := Redis.LPush(ctx, key, value).Err(); err != nil {
		robot.Logger().Error("redis：" + key + "添加失败")
		return err
	}
	return nil
}

func GetList(key string) ([]string, error) {
	lRange := Redis.LRange(ctx, key, 0, -1)
	if lRange.Err() != nil {
		robot.Logger().Error("redis：" + key + "获取失败")
		return nil, lRange.Err()
	}
	return lRange.Val(), nil
}

func DelList(key string) error {
	if err := Redis.Del(ctx, key).Err(); err != nil {
		robot.Logger().Error("redis：" + key + "移除失败")
		return err
	}
	return nil
}
