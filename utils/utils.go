package utils

import (
	"crypto/sha1"
	"github.com/HRDVV/cover-note/model"
	"github.com/gin-gonic/gin"
	"hash"
	"net/http"
)

func CryptoSha1Field(originStr string, handle func(hash hash.Hash)) {
	h := sha1.New()
	h.Write([]byte(originStr))
	handle(h)
	h.Reset()
}

func StatusInternalServerError(ctx *gin.Context) {
	if err := recover(); err != nil {
		ctx.JSON(http.StatusInternalServerError, new(model.Result).Fail(http.StatusInternalServerError, "系统内部错误", nil))
	}
}