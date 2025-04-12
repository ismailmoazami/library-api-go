package controllers

import (
	"library/database"
	"library/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book

	database.DB.Find(&books)

	c.JSON(http.StatusOK, books)

}

func AddBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&book)

	c.JSON(http.StatusOK, book)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := database.DB.Find(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found!"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := database.DB.Find(&book, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := database.DB.Find(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found!"})
		return
	}

	database.DB.Delete(&book, id)
	c.JSON(http.StatusOK, book)
}
