package router

import (
	c "covernote-backend/handle"
	"github.com/gin-gonic/gin"
	)

func ListGroup(r gin.IRouter) {
	list := r.Group("/list")
	{
		list.GET("/get", c.GetListAll)
	}
}
