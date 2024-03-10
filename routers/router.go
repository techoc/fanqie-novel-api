package routers

import "github.com/techoc/fanqie-novel-api/api"

type RouterGroup struct {
}

var (
	bookApi    = api.BookAPI{}
	chapterApi = api.ChapterAPI{}
)
