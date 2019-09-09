package main

import (
	app "github.com/HRDVV/cover-note/bootstrap"
	. "github.com/HRDVV/cover-note/config"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(GlobalConfig.Env)
}

func main() {
	app.Bootstrap()
}
