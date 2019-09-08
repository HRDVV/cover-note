package main

import (
	app "github.com/HRDVV/cover-note/bootstrap"
	. "github.com/HRDVV/cover-note/config"
	"github.com/HRDVV/cover-note/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	var c Config
	env := utils.EnvFlag()
	c.LoadConfig(env)
	gin.SetMode(viper.GetString("env"))
}

func main() {
	app.Bootstrap()
}
