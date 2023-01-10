package utils

import (
	"context"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	bigcacheStore "github.com/eko/gocache/store/bigcache/v4"
	redisStore "github.com/eko/gocache/store/redis/v4"
	"github.com/go-redis/redis/v8"
	"github.com/goccy/go-json"
	"keyi/config"
	"time"
)

var Cache *cache.Cache[[]byte]

func init() {
	fmt.Println("init cache...")
	if config.Config.RedisURL != "" {
		var s = redisStore.NewRedis(redis.NewClient(&redis.Options{
			Addr: config.Config.RedisURL,
		}))
		Cache = cache.New[[]byte](s)
	} else {
		var client, _ = bigcache.New(
			context.Background(),
			bigcache.DefaultConfig(5*time.Minute),
		)
		var s = bigcacheStore.NewBigcache(client)
		Cache = cache.New[[]byte](s)
	}
}

const maxDuration time.Duration = 1<<63 - 1

func SetCache(key string, value any, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	if expiration == 0 {
		expiration = maxDuration
	}
	return Cache.Set(context.Background(), key, data, store.WithExpiration(expiration))
}

func GetCache(key string, value any) error {
	data, err := Cache.Get(context.Background(), key)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, value)
}

func DeleteCache(key string) error {
	err := Cache.Delete(context.Background(), key)
	if err == nil {
		return nil
	}
	if err.Error() == "Entry not found" {
		return nil
	}
	return err
}
