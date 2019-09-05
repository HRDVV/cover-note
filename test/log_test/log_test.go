package log_test

import (
	"errors"
	"fmt"
	"github.com/HRDVV/cover-note/model"
	"github.com/HRDVV/cover-note/utils/jwt"
	"github.com/HRDVV/cover-note/utils/log"
	"github.com/HRDVV/cover-note/utils/redis"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"testing"
)

type Demo struct {
	Name string
	Age  int
}

func TestLogger(t *testing.T) {
	log.Info("%[1]s%[1]s", "dssdsdsd")
	log.Error("343434")
	log.Warn("232323")
}

func TestRedisGet(t *testing.T) {
	log.Info(redis.Get("name"))
}

func TestRedisSet(t *testing.T) {

	redis.Set("name", "[{name: \"hrd\"}]", 0)
}

func TestStruct(t *testing.T) {
	var demo Demo
	demo.Name = "hrd"
	fmt.Println(demo.Name)
	//var r io.Reader
	err := ioutil.WriteFile("text.txt", []byte("4343"), os.ModeAppend)
	if err == nil {
	}
}

type Writer interface {
	Reader()
}
type ReaderNew interface {
	Writer
}

type TagStruct struct {
	Name string "名称"
}

func (t TagStruct) Reader() {
	log.Info("hello world")
}
func (t TagStruct) Demo() {

}

func TestTag(t *testing.T) {
	type Integer int
	var a Integer = 1
	v := reflect.ValueOf(&a)

	log.Info(reflect.TypeOf(a))
	log.Info(reflect.TypeOf(a).Kind())
	log.Info(reflect.ValueOf(a).Kind())
	v = v.Elem()
	log.Info(v.CanSet())
	v.SetInt(3)
	log.Info(strconv.Itoa(int(a)))
}

func TestJwt(t *testing.T) {
	var jwt jwt.Jwt
	var user model.User
	user.UserId = 1
	user.Username = "hrd"
	token, err := jwt.GenerateToken(user)
	if err == nil {
		log.Info(token)
	}
}

func TestParseToken(t *testing.T) {
	var jwt jwt.Jwt
	claims, err := jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJocmQiLCJleHAiOjE1Njc2MDQwODksImp0aSI6IjEiLCJpYXQiOjE1Njc1MTc2ODksImlzcyI6ImNvdmVybm90ZSIsIm5iZiI6MTU2NzUxNzY4OSwic3ViIjoibG9naW4ifQ.7KNNOwJ14gM4CRKksthP60Cg9Se3pIVHfAljl4itDm4")
	if err == nil {
		log.Info(claims.Audience)
	}
}

func DemoInterface(args ...interface{}) {
	fmt.Println(args[0])
}

func TestInterface(t *testing.T) {
	DemoInterface("3","3")
}

func TestHscan(t *testing.T) {
	log.Info(redis.HScan("list", "*"))
}

func TestHset(t *testing.T) {
	for i:=0;i<2000;i++ {
		redis.HSet("list", strconv.Itoa(i), strconv.Itoa(i))
	}
}

func TestIncr(t *testing.T) {
	redis.HDel("list", "0")
}

func TestDefer(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Info(err.(error).Error())
		}
	}()
	panic(errors.New("wewe"))
}



