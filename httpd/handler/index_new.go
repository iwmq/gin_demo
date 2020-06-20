package handler

import (
	"github.com/gin-gonic/gin"
)

// IndexNew ...
func IndexNew() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "home/new.tpl", gin.H{
			"name": "Jay Wang",
		})
	}
}
