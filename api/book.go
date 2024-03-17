package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookAPI struct {
}

// SearchBook 搜索书籍
func (a BookAPI) SearchBook(c *gin.Context) {
	bookName := c.Query("name")
	bookList := bookService.SearchBookByTitle(bookName)
	if bookList == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": bookList,
	})
}

func (a BookAPI) GetDirectoryByBookId(c *gin.Context) {
	bookIdStr := c.Query("book_id")
	bookId, _ := strconv.ParseInt(bookIdStr, 10, 64)
	chapters := bookService.GetDirectoryByBookId(bookId)
	if chapters == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": chapters,
	})
}

func (a BookAPI) GetDetailByBookId(c *gin.Context) {
	bookIdStr := c.Query("book_id")
	bookId, _ := strconv.ParseInt(bookIdStr, 10, 64)
	book := bookService.GetDetailByBookId(bookId)
	if book.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "please search book name first",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": book,
	})
}
