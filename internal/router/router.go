package router

import (
	"crud_go/internal/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoute(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "CRUD with API Gin",
			"version": "1.0",
		})
	})

	book := router.Group("/book")
	{
		book.POST("", handler.CreateBook(db))
		book.GET("/", handler.GetBook(db))
		book.PUT("/:id", handler.UpdateBook(db))
		book.DELETE("/:id", handler.DeleteBook(db))
	}

	return router
}
