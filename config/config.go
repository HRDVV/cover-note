package config

const (
	USE_REDIS_DB   = 1
	REDIS_ADDR     = "172.18.0.4:6379"
	REDIS_PASSWORD = ""
	SECRET         = "HRDVV"
	JWT_EXP_TIME   = 24 * 3600
)

var GoEnv = [3]string{"debug", "test", "release"}