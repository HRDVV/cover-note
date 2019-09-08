package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Development
	Production
}

type Development struct {
	Env string
	Redis
	Jwt
}

type Production struct {
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

func (c *Config) LoadConfig(filename string) {
	viper.SetConfigName(filename)
	viper.AddConfigPath("app_context")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变更：", e.Name)
	})
}
