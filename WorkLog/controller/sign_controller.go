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
	"strings"

	"github.com/gin-gonic/gin"
)

type signcontroller struct{}

var SignController = new(signcontroller)

func (s *signcontroller) SignUp(c *gin.Context) {
	var req request.SignUp
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Parameter input is wrong: "+err.Error(), c)
		return
	}

	if req.Password != req.RePassword {
		response.FailWithMessage("Passwords doesn't match", c)
		return
	}
	req.Password = util.MD5V([]byte(req.Password))
	req.Email = strings.TrimSpace(req.Email)

	if err := service.SignService.SignUpService(req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("Sign up success", c)
}

func (s *signcontroller) SignIn(c *gin.Context) {
	var req request.SignIn
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Parameter input is wrong: "+err.Error(), c)
		return
	}
	log.Info(req)

	req.Password = util.MD5V([]byte(req.Password))
	log.Info(req)
	user, err := service.SignService.SignInService(req)
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
	response.OkWithDetailed(map[string]string{"token": existToken}, "SignInService successful", c)
}

func (s *signcontroller) SignOut(c *gin.Context) {
	user := middleware.GetUser(c)
	fmt.Println(user)

	err := middleware.DelRedisJWT(user.Email)
	if err != nil {
		response.FailWithDetailed(err, "Delete token failed", c)
	}
	response.OkWithMessage("Sign out success", c)
}

func (s *signcontroller) Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in",
		"user":    user,
	})
}
