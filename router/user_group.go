package router

import (
	c "covernote-backend/handle"
	"github.com/gin-gonic/gin"
)

func UserGroup(r gin.IRouter) {
	list := r.Group("/user")
	{
		list.POST("", c.Login)
	}
}
