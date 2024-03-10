package service

import (
	"github.com/techoc/fanqie-novel-api/models"
	"github.com/techoc/fanqie-novel-api/pkg/fanqie"
	"log"
)

type ChapterService struct {
}

func (a ChapterService) GetContentByChapterId(chapterId int64) models.Chapter {
	// 1. 查询数据库中的章节内容
	chapterByChapterId := models.GetChapterByChapterId(chapterId)
	if chapterByChapterId.ChapterID == 0 && chapterByChapterId.Content == "" {
		log.Printf("not find chapter in database,chapterId: %d\n", chapterId)
		// 没有查询到，调用番茄接口获取正文
		chapter := fanqie.GetContentByChapterId(chapterId)
		if chapter.ChapterID == 0 || chapter.Content == "" {
			log.Printf("not find chapter in fanqie,chapterId: %d\n", chapterId)
			return models.Chapter{}
		}
		// 2. 将内容插入数据库
		models.InsertChapter(chapter)
		return chapter
	}
	log.Printf("find chapter in database,chapterId: %d\n", chapterId)
	return chapterByChapterId
}
