package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"go-daily-work/config"
	milog "go-daily-work/log"
	"go-daily-work/middleware"
	"go-daily-work/model"
	"go-daily-work/model/request"
	"go-daily-work/util"
	"gorm.io/gorm"
	"log"
	"time"
)

type loginservice struct{}

var LoginService = new(loginservice)

func CreateToken(email string) (string, error) {
	var user model.User
	util.Master().Where("email = ?", email).First(&user)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["username"] = user.Name
	claims["position"] = user.Position
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(config.Instance.TokenKey))
	if err != nil {
		return "", err
	}
	err = middleware.SetRedisJWT(tokenString, email)
	if err != nil {
		milog.Error(err)
	}
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
