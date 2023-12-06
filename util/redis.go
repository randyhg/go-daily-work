package util

import (
	"go-daily-work/config"
	"log"
	"sync"

	"github.com/go-redis/redis"
)

var rdsCache *rdsStore
var onceRedis sync.Once

type rdsStore struct {
	client redis.UniversalClient
}

func RedisCache() redis.UniversalClient {
	return rdsCache.client
}

func RedisInit() {
	onceRedis.Do(func() {
		if len(config.Instance.RedisCache.Host) > 0 {
			rdsCache = new(rdsStore)
			rdsCache.connectDB(config.Instance.RedisCache)
		}
	})
}

func (r *rdsStore) connectDB(conf config.RedisConfig) {

	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    conf.Host,
		Password: conf.Password,
		DB:       conf.DB,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal("redis connect ping failed, err:", err)
	} else {
		log.Println("redis connect ping response:", "pong", pong)
		log.Println("Redis DB: ", conf.DB)
		r.client = client
	}
}
