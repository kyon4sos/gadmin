package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"neko/src/main/model"
	"time"
)

type Cache struct {
}

var ctx = context.Background()
var rdb *redis.Client
var Test = "test"

func init() {
	rdb = RedisClient()
}

func SetPreFixKey(key string, value model.Base) bool {
	key = value.GetCachePrefix() + ":" + key
	Set(key,value,600)
	return true
}

func Set(key string, value interface{}, exp int) bool {
	marshal, err2 := json.Marshal(value)
	if err2 != nil {
		return false
	}
	_, err := RedisClient().Set(ctx, key, marshal,
		time.Duration(exp)*time.Second).Result()
	if err != nil {
		return false
	}
	return true
}
func Get(key string) (*interface{}, error) {
	var target interface{}
	res, err := RedisClient().Get(ctx, key).Result()
	if len(res) == 0 {
		return nil, nil
	}
	err = json.Unmarshal([]byte(res), &target)
	if err != nil {
		return nil, err
	}
	return &target, nil
}
func RedisClient() *redis.Client {
	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6380",
			Password: "123456",
			DB:       0,
		})
	}
	return rdb
}
