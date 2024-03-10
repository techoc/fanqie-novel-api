package api

import "github.com/techoc/fanqie-novel-api/service"

type ApiGroup struct {
	BookAPI
}

var (
	bookService    = service.BookService{}
	chapterService = service.ChapterService{}
)
