package handle

import (
	"github.com/HRDVV/cover-note/utils/log"
	"github.com/gin-gonic/gin"
)
import . "github.com/HRDVV/cover-note/model"

func GetListAll(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	list := map[string]string {
		"username": "韩瑞达",
		"userId":   "1",
}
	log.Error("3456765432345678")
	ctx.JSON(200, new(Result).Succ(list))
}

//func AddList(context *gin.Context) {
//
//	context.Header("Content-Type", "application/json")
//	r := new(ResponseModel)
//
//
//}
