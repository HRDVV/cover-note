package router

import (
	c "github.com/HRDVV/cover-note/handle"
	"github.com/gin-gonic/gin"
)

func ListGroup(r gin.IRouter) {
	list := r.Group("/list")
	{
		list.GET("/get", c.GetListAll)
		//list.POST("/add", c.AddList)
	}
}
