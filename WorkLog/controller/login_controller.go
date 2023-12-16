package controller

import (
	"fmt"
	"go-daily-work/WorkLog/service"
	"go-daily-work/log"
	"go-daily-work/middleware"
	"go-daily-work/model/request"
	"go-daily-work/model/response"
	"go-daily-work/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type logincontroller struct{}

var LoginController = new(logincontroller)

func (l *logincontroller) Login(c *gin.Context) {
	var req request.Login
	_ = c.ShouldBindJSON(&req)
	log.Info(req)

	req.Password = util.MD5V([]byte(req.Password))
	user, err := service.LoginService.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	existToken, _ := middleware.GetRedisJWT(user.Email)
	if existToken != "" {
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", existToken, 3600*24*30, "", "", false, true)
	} else {
		token, err := service.CreateToken(user.Email)
		existToken = token
		if err != nil {
			log.Fatal(err)
			return
		}
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)
	}
	response.OkWithDetailed(map[string]string{"token": existToken}, "Login successful", c)
}

func (l *logincontroller) Logout(c *gin.Context) {
	user := middleware.GetUser(c)
	fmt.Println(user)

	err := middleware.DelRedisJWT(user.Email)
	if err != nil {
		response.FailWithDetailed(err, "Delete token failed", c)
	}
	response.OkWithMessage("Sign out success", c)
}

func (l *logincontroller) Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in",
		"user":    user,
	})
}
