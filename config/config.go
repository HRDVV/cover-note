package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Env string
	Redis
	Jwt
}

type Redis struct {
	Db       int
	Addr     string
	Password string
}

type Jwt struct {
	Secret  string
	ExpTime int
}

var GlobalConfig Config

func init() {
	viper.SetConfigName("application")
	viper.AddConfigPath("app_context")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	GlobalConfig = config
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变更：", e.Name)
	})
}



