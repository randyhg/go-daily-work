package middleware

import (
	"errors"
	"go-daily-work/config"
	"go-daily-work/model"
	"go-daily-work/model/request"
	"go-daily-work/util"
	"net/http"

	"go-daily-work/model/response"

	"log"

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
			response.FailWithDetailed(gin.H{"reload": true}, "Unblocked or illegally visited", c)
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
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		_, rt := GetRedisJWT(claims.Email)
		if rt != token {
			response.FailWithMessage("Please login first", c)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var user model.User
		util.Master().Where("email = ?", claims.Email).First(&user)

		if user.Id == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)
		c.Next()
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

func GetRedisJWT(email string) (err error, redisJWT string) {
	redisJWT, err = util.RedisCache().Get(email).Result()
	return err, redisJWT
}

func GetUser(c *gin.Context) *model.User {
	token, _ := c.Cookie("Authorization")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		return nil
	}

	var user model.User
	if err := util.Master().Where("email = ?", claims.Email).First(&user).Error; err != nil {
		log.Fatal(err)
		return nil
	}
	return &user
}
