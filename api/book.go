package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type BookAPI struct {
}

// SearchBook 搜索书籍
func (a BookAPI) SearchBook(c *gin.Context) {
	bookName := c.Query("name")
	bookList := bookService.SearchBookByTitle(bookName)
	if bookList == nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "not found",
		})
		return
	}
	c.JSON(200, gin.H{
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
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "not found",
		})
		return
	}
	c.JSON(200, gin.H{
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
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "please search book name first",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": book,
	})
}
