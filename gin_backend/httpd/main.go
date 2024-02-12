package main

import (
	"gin_demo/controllers"
	"gin_demo/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env
	// NOTE: .env must be put in the directory where we invoke the executable
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

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
		books.GET("/edit/:id", controllers.BookEdit())
		books.GET("/api", controllers.GetAllAPI())
		books.POST("/api", controllers.CreateBookAPI())
		books.PATCH("/api/:id", controllers.UpdateBookAPI())
		books.DELETE("/api/:id", controllers.DeleteBookAPI())
	}

	r.Run()
}
