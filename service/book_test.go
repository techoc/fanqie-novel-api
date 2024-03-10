package service

import (
	"github.com/techoc/fanqie-novel-api/conf"
	"github.com/techoc/fanqie-novel-api/models"
	"testing"
)

func TestBookService_SearchBookByTitle(t *testing.T) {
	conf.Setup()
	models.Setup()
	bookService := &BookService{}
	books := bookService.SearchBookByTitle("火影我绳树绝不活在回忆里")
	for _, book := range books {
		bookService.GetDirectoryByBookId(book.BookID)
	}
}
