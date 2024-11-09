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
	if len(chapterByChapterId.Content) < 128 {
		log.Printf("not find chapter in database,chapterId: %d\n", chapterId)
		// 没有查询到，调用番茄接口获取正文
		//todo 改造获取正文逻辑 当获取正文时 通过NovelData查询数据库是否有该书籍
		// todo 获取章节内容时，查询数据库是否有该书籍并插入数据是否为VIP书籍
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
