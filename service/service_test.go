package service

import (
	"fmt"
	"github.com/techoc/fanqie-novel-api/conf"
	"github.com/techoc/fanqie-novel-api/models"
	"testing"
)

func TestChapterService_GetContentByChapterId(t *testing.T) {
	conf.Setup()
	models.Setup()
	chapter := ChapterService{}.GetContentByChapterId(7058190355161842180)
	fmt.Println(chapter.Content)
}
