package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type Chapter struct {
	ChapterID     int64  // 章节ID
	Name          string // 章节名
	BookId        int64  // 书籍ID
	ChapterNumber int    // 第几章
	WordCount     int    // 字数
	Content       string // 内容
	NextGroupID   int64  // 下一章
	NextItemID    int64  // 下一章
	PreGroupID    int64  // 上一章
	PreItemID     int64  // 上一章
	gorm.Model
}

func InsertChapter(chapter Chapter) {
	var dbChapter Chapter
	result := db.Where("chapter_id = ?", chapter.ChapterID).First(&dbChapter)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { // 不存在
		result = db.Create(&chapter)
		if result.RowsAffected > 0 {
			log.Printf("插入章节成功,章节名 %s ,章节id %d\n", chapter.Name, chapter.ChapterID)
		} else {
			log.Printf("插入章节失败,章节名 %s ,章节id %d\n", chapter.Name, chapter.ChapterID)
		}
		return
	} else {
		log.Printf("章节已存在,章节名 %s ,章节id %d\n", chapter.Name, chapter.ChapterID)
		// 更新
		updates := db.Where("chapter_id = ?", chapter.ChapterID).Updates(&chapter)
		if updates.RowsAffected > 0 {
			log.Printf("更新章节成功,章节名 %s ,章节id %d\n", chapter.Name, chapter.ChapterID)
		}
	}
}

func InsertChapterList(chapterList []Chapter) {
	// 批量查找
	var dbChapterList []Chapter
	var chapterIdList []int64
	for _, chapter := range chapterList {
		chapterIdList = append(chapterIdList, chapter.ChapterID)
	}
	db.Where("chapter_id in ?", chapterIdList).Find(&dbChapterList)
	// 将dbChapterList转换为map
	dbChapterMap := make(map[int64]Chapter)
	for _, chapter := range dbChapterList {
		dbChapterMap[chapter.ChapterID] = chapter
	}
	for _, chapter := range chapterList {
		// 判断是否已存在
		if dbChapter, ok := dbChapterMap[chapter.ChapterID]; ok {
			// 将API获取的书籍和数据库获取的书籍合并
			chapter.ID = dbChapter.ID
		}
	}
	save := db.Save(&chapterList)
	if save.RowsAffected > 0 {
		log.Printf("insert chapter success,bookId: %d\n", chapterList[0].BookId)
	}
}

func GetChapterByChapterId(chapterId int64) (chapter Chapter) {
	db.Where("chapter_id = ?", chapterId).First(&chapter)
	return
}
