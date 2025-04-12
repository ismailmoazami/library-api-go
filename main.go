package main

import (
	"library/controllers"

	"library/database"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	database.InitDB()

	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.AddBook)
	router.GET("/books/:id", controllers.GetBookByID)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run("localhost:9090")
}
