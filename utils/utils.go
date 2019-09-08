package utils

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"github.com/HRDVV/cover-note/model"
	"github.com/HRDVV/cover-note/utils/log"
	"github.com/gin-gonic/gin"
	"hash"
	"net/http"
	"os"
)

func CryptoSha1Field(originStr string, handle func(hash hash.Hash)) {
	h := sha1.New()
	h.Write([]byte(originStr))
	handle(h)
	h.Reset()
}

func StatusInternalServerError(ctx *gin.Context) {
	if err := recover(); err != nil {
		log.Error(err.(error).Error())
		ctx.JSON(http.StatusInternalServerError, new(model.Result).Fail(http.StatusInternalServerError, "系统内部错误", nil))
	}
}

func EnvFlag() string {
	var env *string = flag.String("env", "dev", "env params: enum(dev, prod)")
	if !flag.Parsed() {
		flag.Parse()
	}
	if *env != "dev" && *env != "prod" {
		fmt.Println("enum(dev, prod)")
		os.Exit(1)
	}
	return *env
}
