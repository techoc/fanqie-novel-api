package service

import (
	"github.com/techoc/fanqie-novel-api/models"
	"github.com/techoc/fanqie-novel-api/pkg/fanqie"
	"log"
)

type BookService struct {
}

func (a BookService) SearchBookByTitle(title string) []models.Book {
	// 1. 查询数据库
	book := models.GetBookByBookName(title)
	if book.Name == title {
		log.Println("find book in database by book name")
		return []models.Book{book}
	}
	log.Println("not find book in database by book name")
	var bookList []models.Book
	//1. 调用fanqie的接口
	bookListFanqie := fanqie.Search(title, 1, false)
	//2. 将返回的书籍入库
	// 改造成批量插入
	bookList = models.InsertBookList(bookListFanqie)
	return bookList
}

func (a BookService) GetDirectoryByBookId(bookId int64) []models.Chapter {
	//1. 调用fanqie的接口
	chapterList := fanqie.GetDirectoryByBookId(bookId)
	//2. 将返回的目录入库
	models.InsertChapterList(chapterList)
	return chapterList
}

func (a BookService) GetDetailByBookId(bookId int64) models.Book {
	return models.GetBookByBookId(bookId)
}
