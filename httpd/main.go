package main

import (
	"gin_demo/controllers"
	"gin_demo/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()
	r.Static("/assets", "assets")
	r.StaticFile("/favicon.ico", "assets/favicon.ico")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", controllers.HomeGet())

	books := r.Group("/books")
	{
		books.GET("/", controllers.BookList())
		books.GET("/new", controllers.BookNew())
		books.GET("/show/:id", controllers.BookShow())
		books.GET("/api", controllers.GetAllAPI())
		books.POST("/api", controllers.CreateBookAPI())
	}

	r.Run()
}
