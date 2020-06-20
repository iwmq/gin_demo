package handler

import (
	"github.com/gin-gonic/gin"
)

// IndexGet ...
func IndexGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "home/index.tpl", gin.H{
			"name": "Jay Wang",
		})
	}
}
