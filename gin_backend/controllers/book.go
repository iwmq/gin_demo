package controllers

import (
	"gin_demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBookInput ...
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdateBookInput ...
type UpdateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// GetAllAPI ...
func GetAllAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.Book
		models.DB.Find(&books)
		c.JSON(http.StatusOK, gin.H{
			"data": books,
		})
	}
}

// CreateBookAPI ...
func CreateBookAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate input
		var input CreateBookInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Create book
		book := models.Book{Title: input.Title, Author: input.Author}
		models.DB.Create(&book)
		c.JSON(http.StatusOK, gin.H{
			"data": book,
		})
	}
}

// DetailBookAPI ...
func DetailBookAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book

		id := c.Param("id")

		if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
			// Error
		}

		c.JSON(http.StatusOK, gin.H{
			"data": book,
		})
	}
}

// UpdateBookAPI ...
func UpdateBookAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Get book if exists
		var book models.Book
		if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Record not found!",
			})
			return
		}

		// Validate input
		var input UpdateBookInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Update book
		models.DB.Model(&book).Update(input)
		c.JSON(http.StatusOK, gin.H{
			"data": book,
		})
	}
}

// DeleteBookAPI ...
func DeleteBookAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Get book if exists
		var book models.Book
		if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Record not found!",
			})
			return
		}

		// Delete book
		models.DB.Delete(&book)
		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

// BookList ...
func BookList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "books/index.tpl", gin.H{
			"name": "Jay Wang",
		})
	}
}

// BookNew ...
func BookNew() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "books/new.tpl", gin.H{
			"name": "Jay Wang",
		})
	}
}


// BookShow ...
func BookShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book

		id := c.Param("id")

		if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
			// Error
		}

		c.HTML(200, "books/show.tpl", gin.H{
			"name": "Jay Wang",
			"data": book,
		})
	}
}


// BookEdit ...
func BookEdit() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book

		id := c.Param("id")

		if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
			// Error
		}

		c.HTML(200, "books/edit.tpl", gin.H{
			"name": "Jay Wang",
			"data": book,
		})
	}
}
