package log_test

import (
	"covernote-backend/utils/log"
	"covernote-backend/utils/redis"
	"testing"
)

func TestLogger(t *testing.T) {
	log.Info("%[1]s%[1]s","dssdsdsd")
	log.Error("343434")
	log.Warn("232323")
}

func TestRedisGet(t *testing.T) {
	log.Info(redis.Get("name"))
}

func TestRedisSet(t *testing.T) {
	redis.Set("name", "[{name: \"hrd\"}]", 0)
}

