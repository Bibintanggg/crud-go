package handler

import (
	"crud_go/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book []model.Book

		// bind get data
		result := db.Find(&book)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Gagal mengambil data"})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    book,
			"count":   len(book),
		})
	}
}

func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book model.Book

		// decode
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})

			return
		}

		// Melakukan create data dengna model book
		if err := db.Create(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Gagal menyimmpan data"})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Book berhasil dibuat !",
			"data":    book,
		})
	}
}

func UpdateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var book model.Book

		// find book
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Book tidak ditemukan !"})

			return
		}

		// bind update data
		var updateData model.Book
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON Body" + err.Error()})

			return
		}

		// update data
		result := db.Model(&book).Updates(updateData)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Gagal mengudate book"})

			return
		}

		// mencari book sesuai id
		db.First(&book, id)

		c.JSON(http.StatusOK, gin.H{
			"succes":  true,
			"message": "Book berhasil diperbarui",
			"data":    book,
		})
	}
}

func DeleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var book model.Book

		result := db.Delete(&book, id)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Gagal menghapus book",
			})

			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Book not found",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Book berhasil dihapus",
			"success": true,
		})

	}
}
