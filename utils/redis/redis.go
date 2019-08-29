package redis

import (
	"covernote-backend/config"
	"github.com/go-redis/redis"
	"time"
)


var client = redis.NewClient(&redis.Options{
	Addr: config.REDIS_ADDR,
	Password: config.REDIS_PASSWORD,
	DB: config.USE_REDIS_DB,
})

func Get(key string) string{
	val, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func Set(key string, value string, epr time.Duration) {
	err := client.Set(key,  value, epr).Err()
	if err != nil {
		panic(err)
	}
}




