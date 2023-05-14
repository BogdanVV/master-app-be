package routing

import (
	"github.com/bogdanvv/master-app-be/controller"
	"github.com/bogdanvv/master-app-be/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(c controller.Controller) *gin.Engine {
	r := gin.Default()

	auth := r.Group("auth")
	{
		auth.POST("/sign-up", c.Signup)
		auth.POST("/login", c.Login)
	}

	api := r.Group("api")
	api.Use(middleware.CheckAuth)
	{
		api.GET("/test", c.TestApi)

		users := api.Group("users")
		{
			users.PUT("/:userId", c.UpdateUser)
		}
	}

	return r
}
