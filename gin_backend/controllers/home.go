package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
)

// HomeGet ...
func HomeGet() gin.HandlerFunc {
	is_develop := os.Getenv("DEVELOP")
	var template_name string
	if is_develop == "true" {
		template_name = "home/index_vite.tpl"
	} else {
		template_name = "home/index.tpl"
	}

	return func(c *gin.Context) {
		c.HTML(200, template_name, gin.H{
			"name": "Jay Wang",
		})
	}
}
