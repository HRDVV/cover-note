package middleware

import (
	"covernote-backend/model"
	"covernote-backend/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AuthLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessKey string
		if ctx.Query("token") != "" {
			accessKey = ctx.Query("token")
		} else if jwt, err := ctx.Cookie("jwt"); err == nil  {
			accessKey = jwt
		} else {
			accessKey = ctx.GetHeader("Authorization")
		}
		if accessKey == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, new(model.Result).Fail(http.StatusUnauthorized, "token不存在", nil))
			return
		}
		var jwt jwt.Jwt
		var t model.Token
		claims, err := jwt.ParseToken(accessKey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, new(model.Result).Fail(http.StatusUnauthorized, "不合法的token", nil))
			ctx.Abort()
			return
		}
		t.Username = claims.Audience
		if t.QueryTokenByName().AccessKey != accessKey || claims.ExpiresAt < time.Now().Unix() {
			fmt.Println(claims.Audience)
			ctx.JSON(http.StatusUnauthorized, new(model.Result).Fail(http.StatusUnauthorized, "无效的token", nil))
			ctx.Abort()
			return
		}
		ctx.Set("accessKey", accessKey)
		ctx.Next()
	}
}
