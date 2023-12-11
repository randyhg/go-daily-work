package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"go-daily-work/config"
	"go-daily-work/model"
	"go-daily-work/model/request"
	"go-daily-work/util"
	"log"
	"time"

	"gorm.io/gorm"
)

type loginservice struct{}

var LoginService = new(loginservice)

func CreateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(config.Instance.TokenKey))
	if err != nil {
		return "", err
	}
	SetRedisJWT(tokenString, email)
	return tokenString, nil
}

func (l *loginservice) Login(req request.Login) (user model.User, err error) {
	if err = util.Master().Where("email = ? AND password = ?", req.Email, req.Password).First(&user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal(err)
		} else {
			return user, err
		}
	}
	return user, nil
}

func SetRedisJWT(token string, email string) (err error) {
	timer := 60 * 60 * 24 * 7 * time.Second
	err = util.RedisCache().Set(email, token, timer).Err()
	return err
}

func DelRedisJWT(email string) error {
	return util.RedisCache().Del(email).Err()
}
