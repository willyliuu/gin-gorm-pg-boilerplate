package main

import (
	"gin-gorm-pg/controllers"
	"gin-gorm-pg/middlewares"
	"gin-gorm-pg/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	book := r.Group("/books")

	book.GET("/", controllers.FindBooks)
	book.GET("/:id", controllers.FindBook)
	book.POST("/", controllers.CreateBook)
	book.PUT("/:id", controllers.UpdateBook)
	book.DELETE("/:id", controllers.DeleteBook)

	user := r.Group("/users")

	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.Login)
	user.Use(middlewares.JwtAuthMiddleware())
	user.GET("/", controllers.CurrentUser)

	r.Run() // listen and serve on 0.0.0.0:8080
}
