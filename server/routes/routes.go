package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sujin/todo-app/controllers"
	_ "github.com/sujin/todo-app/docs"
	"github.com/sujin/todo-app/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine) {
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3600", "http://localhost:9000","http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "Jesus Coming Soon!"})
    })

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    auth := r.Group("/auth")
    {
        auth.POST("/signup", controllers.Signup)
        auth.POST("/login", controllers.Login)
    }

    todoRoutes := r.Group("/todos")
    todoRoutes.Use(middleware.AuthMiddleware())
    {
        todoRoutes.GET("", controllers.GetTodos)
        todoRoutes.GET("/:id", controllers.GetTodoByID)
        todoRoutes.POST("", controllers.CreateTodo)
        todoRoutes.PUT("/:id", controllers.UpdateTodo)
        todoRoutes.DELETE("/:id", controllers.DeleteTodo)
    }
}

