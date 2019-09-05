package handle

import (
	"github.com/gin-gonic/gin"
)
import . "covernote-backend/model"

func GetListAll(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	list := map[string]string{
		"username": "韩瑞达",
		"userId":   "1",
	}

	ctx.JSON(200, new(Result).Succ(list))
}

//func AddList(context *gin.Context) {
//
//	context.Header("Content-Type", "application/json")
//	r := new(ResponseModel)
//
//
//}
