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
	"time"
)

type signservice struct{}

var SignService = new(signservice)

func CreateToken(email string) (string, error) {
	var user model.User
	util.Master().Where("email = ?", email).First(&user)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["username"] = user.Name
	//claims["position"] = user.Position
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

func (s *signservice) SignInService(req request.SignIn) (user model.User, err error) {
	if err = util.Master().Where("email = ? AND password = ?", req.Email, req.Password).First(&user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			milog.Error(err)
			return user, err
		} else {
			return user, err
		}
	}
	return user, nil
}

func (s *signservice) SignUpService(req request.SignUp) error {
	var user model.User
	user.Name = req.Name
	user.Password = req.Password
	user.Position = req.Position

	//if req.Position == "admin" {
	//	user.RoleId = model.AdminRoleId
	//} else if req.Position == "manager" {
	//	user.RoleId = model.ManagerRoleId
	//} else {
	//	user.RoleId = model.StaffRoleId
	//}

	if s.CheckEmail(util.Master(), "email = ?", req.Email) != nil {
		return errors.New(req.Email + " has been used")
	}
	user.Email = req.Email

	if err := util.Master().Create(&user).Error; err != nil {
		milog.Error(err)
		return err
	}

	err := middleware.AddRoleToUser(req.Position, user)
	if err != nil {
		milog.Error(err)
	}
	return nil
}

func (s *signservice) CheckEmail(db *gorm.DB, where ...interface{}) *model.User {
	var user model.User
	if err := db.Take(&user, where...).Error; err != nil {
		return nil
	}
	return &user
}
