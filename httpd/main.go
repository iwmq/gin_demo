package main

import (
	"database/sql"
	"fmt"
	"gin_demo/httpd/handler"
	"gin_demo/platform/newsfeed"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "jay:jay123@(localhost:3306)/gin?parseTime=true")
	title := "Hello"
	post := "world"
	if err == nil {
		// db.Exec("INSERT INTO posts (title, post) VALUES (?, ?)", title, post)
		fmt.Println(db, err, title, post)
	}

	feed := newsfeed.New()
	r := gin.Default()
	r.Static("/assets", "assets")
	r.StaticFile("/favicon.ico", "assets/favicon.ico")

	r.LoadHTMLGlob("templates/**/*")
	r.LoadHTMLFiles("templates/message.tpl.html")

	r.GET("/", handler.IndexGet())
	r.GET("/new", handler.IndexNew())
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	r.Run()
}
