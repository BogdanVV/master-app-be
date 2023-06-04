package routing

import (
	"github.com/bogdanvv/master-app-be/controller"
	"github.com/bogdanvv/master-app-be/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes(c controller.Controller) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware)

	auth := r.Group("auth")
	{
		auth.POST("/sign-up", c.Signup)
		auth.POST("/login", c.Login)
		auth.POST("/refresh-token", middleware.CheckAuth, c.RefreshToken)
	}

	api := r.Group("api")
	api.Use(middleware.CheckAuth)
	{
		api.GET("/test", c.TestApi)

		users := api.Group("users")
		{
			users.PUT("/:userId", c.UpdateUser)
		}

		todos := api.Group("todos")
		{
			todos.POST("/", c.CreateTodo)
			todos.GET("/", c.GetAllTodos)
			todos.GET("/:id", c.GetTodoById)
			todos.DELETE("/:id", c.DeleteTodoById)
			todos.PUT("/:id", c.UpdateTodoById)
		}
	}

	return r
}
