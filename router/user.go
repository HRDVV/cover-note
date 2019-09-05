package router

import (
	c "github.com/HRDVV/cover-note/handle"
	"github.com/gin-gonic/gin"
)

func UserGroup(r gin.IRouter) {
	list := r.Group("/user")
	{
		list.POST("/login", c.Login)
		list.POST("/register", c.Register)
	}
}
