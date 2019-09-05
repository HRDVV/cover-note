package bootstrap

import (
	"github.com/HRDVV/cover-note/middleware"
	"github.com/gin-gonic/gin"
)
import r "github.com/HRDVV/cover-note/router"

func Bootstrap() {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	// 用户相关的路由
	r.UserGroup(router)
	// 鉴权中间件
	router.Use(middleware.AuthLogin())
	// 列表相关的路由
	r.ListGroup(router)

	if err := router.Run(":8888"); err != nil {
		panic(err)
	}
}
