package app

import (
	"github.com/gin-gonic/gin"
	"go-daily-work/WorkLog/route"
	"go-daily-work/config"
)

func StartServer() {
	c := gin.Default()
	route.RegisterRoutes(c)
	c.Run(config.Instance.Port)
}
