package jwt

import (
	"covernote-backend/config"
	"covernote-backend/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Jwt struct {
	claims      jwt.Claims
	expiresTime int64
}

func (j *Jwt) GenerateToken(user model.User) (string, error) {
	j.expiresTime = time.Now().Unix() + config.JWT_EXP_TIME
	j.claims = jwt.StandardClaims{
		Audience:  user.Username,     // 受众
		ExpiresAt: j.expiresTime,     // 失效时间
		Id:        user.UserId,       // 编号
		IssuedAt:  time.Now().Unix(), // 签发时间
		Issuer:    "covernote",       // 签发人
		NotBefore: time.Now().Unix(), // 生效时间
		Subject:   "login",           // 主题
	}
	jwtSecret := []byte(config.SECRET)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, j.claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func (j *Jwt) ParseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.SECRET), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
