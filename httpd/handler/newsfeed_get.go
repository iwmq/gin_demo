package handler

import (
	"gin_demo/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewsfeedGet ...
func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, gin.H{
			"result": results,
		})
	}
}
