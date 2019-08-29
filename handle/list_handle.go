package handle

import "github.com/gin-gonic/gin"
import ."covernote-backend/model"

func GetListAll(context *gin.Context) {
	context.Header("Content-Type", "application/json")
		r := new(ResponseModel)
		list := map[string]string{
		"username": "韩瑞达",
		"userId": "1",
	}

	context.JSON(200, r.Succ(list))
}
