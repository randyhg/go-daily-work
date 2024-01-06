package middleware

import (
	"errors"
	"go-daily-work/config"
	milog "go-daily-work/log"
	"go-daily-work/model"
	"go-daily-work/model/request"
	"go-daily-work/util"
	"net/http"
	"time"

	"go-daily-work/model/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(config.Instance.TokenKey),
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.FailWithMessage("Please login first", c)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		j := NewJWT()

		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "Authorization expires", c)
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			response.FailWithMessage("Please login first", c)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		rt, _ := GetRedisJWT(claims.Email)
		if rt != token {
			response.FailWithMessage("Please login first", c)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var user model.User
		util.Master().Where("email = ?", claims.Email).First(&user)

		if user.Id == 0 {
			response.FailWithMessage("Unregistered user", c)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)
		c.Next()
	}
}

func Permission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := GetUser(c)
		can, err := util.Permify().UserHasPermission(uint(user.Id), permission)
		if err != nil {
			milog.Error(err)
			response.FailWithMessage("You don't have access to this route", c)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if can {
			c.Next()
			return
		} else {
			response.FailWithMessage("You don't have access to this route", c)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

func SetRedisJWT(token string, email string) (err error) {
	timer := 1 * time.Hour
	err = util.RedisCache().Set(email, token, timer).Err()
	if err != nil {
		return err
	}
	return nil
}

func DelRedisJWT(email string) error {
	return util.RedisCache().Del(email).Err()
}

func GetRedisJWT(email string) (redisJWT string, err error) {
	redisJWT, err = util.RedisCache().Get(email).Result()
	return redisJWT, err
}

func GetUser(c *gin.Context) *model.User {
	//userInterface, ok := c.Get("user")
	//user := userInterface.(model.User)
	//if !ok {
	//	return nil
	//}
	//return &user

	if userInterface, ok := c.Get("user"); userInterface != nil {
		user := userInterface.(model.User)
		if !ok {
			return nil
		}
		return &user
	} else {
		token, err := c.Cookie("Authorization")
		if err != nil {
			milog.Error(err)
			return nil
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			milog.Error(err)
			return nil
		}

		var user model.User
		if err := util.Master().Where("email = ?", claims.Email).First(&user).Error; err != nil {
			milog.Error(err)
			return nil
		}
		return &user
	}
}

func AddRoleToUser(position string, user model.User) error {
	if err := util.Permify().AddRolesToUser(uint(user.Id), position); err != nil {
		return err
	}
	return nil
}
