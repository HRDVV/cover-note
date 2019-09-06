package model

import (
	"encoding/json"
	"github.com/HRDVV/cover-note/utils/log"
	"github.com/HRDVV/cover-note/utils/redis"
)

const USER_INFO_KEY = "USER:INFO_TABLE"

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User)UserRegister() bool {
	userJson, err := json.Marshal(u)
	if err == nil {
		return redis.HSet(USER_INFO_KEY, u.Username, string(userJson))
	} else {
		return false
	}
}

func (u *User)QueryAllUser() map[string]string {
	return redis.HScan(USER_INFO_KEY, "*")
}

func (u *User)QueryUserByName() User{
	var user User
	userJson := redis.HGet(USER_INFO_KEY, u.Username)
	log.Error(userJson)
	if userJson != "" {
		err := json.Unmarshal([]byte(userJson), &user)
		if err != nil {
			panic(err)
		}
	}
	return user
}
