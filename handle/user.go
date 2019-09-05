package handle

import (
	. "covernote-backend/model"
	"covernote-backend/utils"
	"covernote-backend/utils/jwt"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"hash"
	"net/http"
)

func Login(ctx *gin.Context) {
	defer utils.StatusInternalServerError(ctx)
	var user User
	var t Token
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, new(Result).Fail(http.StatusBadRequest, "参数错误", nil))
		return
	}
	redisUser := user.QueryUserByName()
	if redisUser.Username == "" {
		ctx.JSON(http.StatusOK, new(Result).Fail(http.StatusOK, "该用户不存在", nil))
	} else {
		utils.CryptoSha1Field(user.Password, func(h hash.Hash) {
			user.Password = hex.EncodeToString(h.Sum(nil))
		})
		if user.Password != redisUser.Password {
			ctx.JSON(http.StatusOK, new(Result).Fail(http.StatusOK, "密码错误", nil))
		} else {
			var jwt jwt.Jwt
			if accessKey, err := jwt.GenerateToken(user); err == nil {
				t.AccessKey = accessKey
				t.Username = user.Username
				t.SaveToken()
				ctx.JSON(http.StatusOK, new(Result).Succ(true))
			}
		}
	}

}

func Register(ctx *gin.Context) {
	defer utils.StatusInternalServerError(ctx)
	var user User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, new(Result).Fail(http.StatusBadRequest, "参数错误", nil))
		return
	}
	if user.QueryUserByName().Username != "" {
		ctx.JSON(http.StatusOK, new(Result).Fail(http.StatusOK, "该用户已存在", nil))
		return
	}
	utils.CryptoSha1Field(user.Password, func(h hash.Hash) {
		user.Password = hex.EncodeToString(h.Sum(nil))
	})
	if ok := user.UserRegister(); ok {
		ctx.JSON(http.StatusOK, new(Result).Succ(true))
	} else {
		panic(errors.New("register error"))
	}
}


