package bootstrap

import "github.com/gin-gonic/gin"
import r "covernote-backend/router"

func Bootstrap() {
	router := gin.New()
	//  列表相关的路由
	r.ListGroup(router)

	router.Run(":8888")
}