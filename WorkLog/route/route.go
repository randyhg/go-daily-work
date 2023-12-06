package route

import (
	"go-daily-work/WorkLog/controller"
	"go-daily-work/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(c *gin.Engine) {
	public := c.Group("api")
	{
		public.POST("/login", controller.LoginController.Login)
		public.GET("/logout", controller.LoginController.Logout)
		public.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		})
	}

	private := c.Group("api")
	private.Use(middleware.JWTAuth())
	{
		private.GET("/validate", middleware.JWTAuth(), controller.LoginController.Validate)
	}
}
