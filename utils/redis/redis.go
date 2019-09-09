package redis

import (
	"github.com/HRDVV/cover-note/config"
	"github.com/go-redis/redis"
	"time"
)

const (
	HSCAN_THRESHOLD int64 = 1000
	CURSOR_END            = 0
)

var client = redis.NewClient(&redis.Options{
	Addr:     config.GlobalConfig.Redis.Addr,
	Password: config.GlobalConfig.Redis.Password,
	DB:       config.GlobalConfig.Redis.Db,
})

func Get(key string) string {
	isExists := client.Exists(key).Val()
	if isExists == 1 {
		val, err := client.Get(key).Result()
		if err != nil {
			panic(err)
		}
		return val
	} else {
		return ""
	}
}

func Set(key string, value interface{}, epr time.Duration) bool {
	isSucc, err := client.Set(key, value, epr).Result()
	if err != nil {
		panic(err)
	}
	if isSucc == "ok" {
		return true
	} else {
		return false
	}
}

func Incr(key string) int64 {
	isExists, err := client.Exists(key).Result()
	if err != nil {
		panic(err)
	}
	if isExists == 1 {
		return client.Incr(key).Val()
	} else {
		Set(key, 0, 0)
		return client.Incr(key).Val()
	}
}

func HGet(key string, field string) string {
	isExists := client.HExists(key, field).Val()
	if isExists {
		result, err := client.HGet(key, field).Result()
		if err != nil {
			panic(err)
		}
		return result
	} else {
		return ""
	}
}

func HSet(key string, field string, value interface{}) bool {
	isSucc, err := client.HSet(key, field, value).Result()
	if err != nil {
		panic(err)
	}
	return isSucc
}

func HScan(key string, match string) map[string]string {
	var keys []string
	var cursor uint64 = 0
	var err error
	var result = make(map[string]string)
	for {
		keys, cursor, err = client.HScan(key, cursor, match, HSCAN_THRESHOLD).Result()
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(keys); i += 2 {
			result[keys[i]] = keys[i+1]
		}
		if cursor != CURSOR_END {
			cursor = cursor
		} else {
			return result
		}
	}
}

func HDel(key string, fields ...string) bool {
	isSucc, err := client.HDel(key, fields...).Result()
	if err != nil {
		panic(err)
	}
	if isSucc == 1 {
		return true
	} else {
		return false
	}
}
