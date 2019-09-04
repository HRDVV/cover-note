package router

import (
	c "covernote-backend/handle"
	"covernote-backend/middleware"
	"github.com/gin-gonic/gin"
)

func ListGroup(r gin.IRouter) {
	list := r.Group("/list")
	list.Use(middleware.AuthLogin())
	{
		list.GET("/get", c.GetListAll)
		//list.POST("/add", c.AddList)
	}
}
