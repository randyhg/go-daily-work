package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-daily-work/WorkLog/controller"
	"go-daily-work/middleware"
	"net/http"
)

func RegisterRoutes(c *gin.Engine) {
	public := c.Group("api")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	config.AllowCredentials = true

	public.Use(cors.New(config))
	// Handling preflight requests
	public.OPTIONS("/*any", func(c *gin.Context) {
		c.Status(200)
	})
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

		workLog := private.Group("workLog")
		{
			workLog.GET("/list", controller.WorkLogController.GetWorkLog)
			workLog.POST("/add", controller.WorkLogController.AddWorkLog)
			workLog.POST("/update", controller.WorkLogController.EditWorkLog)
			workLog.POST("/delete", controller.WorkLogController.DeleteWorkLog)
		}

		category := private.Group("category")
		{
			category.GET("/list", controller.CategoryController.GetCategory)
			category.POST("/add", controller.CategoryController.AddCategory)
			category.POST("/update", controller.CategoryController.EditCategory)
			category.POST("/delete", controller.CategoryController.DeleteCategory)
		}

		project := private.Group("project")
		{
			project.GET("/list", controller.ProjectController.GetProject)
			project.POST("/add", controller.ProjectController.AddProject)
			project.POST("/update", controller.ProjectController.EditProject)
			project.POST("/delete", controller.ProjectController.DeleteProject)
		}
	}
}
