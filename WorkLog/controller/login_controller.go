package controller

import (
	"fmt"
	"go-daily-work/WorkLog/service"
	"go-daily-work/middleware"
	"go-daily-work/model/request"
	"go-daily-work/model/response"
	"go-daily-work/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type logincontroller struct{}

var LoginController = new(logincontroller)

func (l *logincontroller) Login(c *gin.Context) {
	var req request.Login
	_ = c.ShouldBindJSON(&req)

	req.Password = util.MD5V([]byte(req.Password))
	user, err := service.LoginService.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token, err := service.CreateToken(user.Email)
	if err != nil {
		log.Fatal(err)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)
	response.OkWithDetailed(map[string]string{"token": token}, "Login successful", c)
}

func (l *logincontroller) Logout(c *gin.Context) {
	user := middleware.GetUser(c)
	fmt.Println(user)

	err := service.DelRedisJWT(user.Email)
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
