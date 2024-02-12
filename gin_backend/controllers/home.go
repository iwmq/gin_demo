package controllers

import "github.com/gin-gonic/gin"

// HomeGet ...
func HomeGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "home/index_vite.tpl", gin.H{
			"name": "Jay Wang",
		})
	}
}
