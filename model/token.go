package model

import (
	"covernote-backend/utils/redis"
	"encoding/json"
)

const USER_TOKEN = "USER:TOKEN"

type Token struct {
	Username    string   `json:"username"`
	AccessKey   string   `json:"accessKey"`
}

func (t *Token) SaveToken() bool {
	jsonJson, err := json.Marshal(t)
	if (err != nil) {
		panic(err)
	}
	return redis.HSet(USER_TOKEN, t.Username, string(jsonJson))
}

func (t *Token) QueryTokenByName() Token {
	var token Token
	tJson := redis.HGet(USER_INFO_KEY,  t.Username)
	json.Unmarshal([]byte(tJson), &token)
	return token
}
