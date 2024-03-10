package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type Book struct {
	BookID          int64   // 书籍ID
	Name            string  // 书名
	Author          string  // 作者
	Desc            string  // 书籍简介
	Category        string  // 分类
	CreationStatus  string  // 状态
	ChaptersCount   int     // 章节数
	FirstChapterId  int64   // 第一章ID
	LastChapterId   int64   // 最后一章ID
	LastChapterTime int64   // 最后更新时间（时间戳格式）
	Score           float32 // 评分
	WordCount       int64   // 书籍总字数
	ReadCount       int64   // 阅读量
	ThumbUrl        string  // 封面图
	gorm.Model
}

func InsertBook(book Book) Book {
	var dbBook Book
	// 通过bookId 查询是否已经入库了
	result := db.Where("book_id = ?", book.BookID).First(&dbBook)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { // 不存在
		result = db.Create(&book)
		if result.RowsAffected > 0 {
			log.Printf("insert book success, bookName: %s,bookId: %d\n", book.Name, book.BookID)
		} else {
			log.Printf("insert book failed, bookName: %s,bookId: %d\n", book.Name, book.BookID)
		}
		return book
	}
	//更新书籍
	//如果上次更新时间与当前时间相差大于24小时，且未完结，则更新书籍
	now := time.Now()
	duration := now.Sub(dbBook.UpdatedAt)

	if dbBook.CreationStatus == "1" && duration.Hours() > 24 {
		log.Printf("%s %d 未完结书籍更新时间大于24小时，更新书籍信息\n", dbBook.Name, dbBook.BookID)

		dbBook.Name = book.Name
		dbBook.Author = book.Author                   // 作者
		dbBook.Desc = book.Desc                       // 书籍简介
		dbBook.Category = book.Category               // 分类
		dbBook.CreationStatus = book.CreationStatus   // 状态
		dbBook.ChaptersCount = book.ChaptersCount     // 章节数
		dbBook.FirstChapterId = book.FirstChapterId   // 第一章ID
		dbBook.LastChapterId = book.LastChapterId     // 最后一章ID
		dbBook.LastChapterTime = book.LastChapterTime // 最后更新时间（时间戳格式）
		dbBook.Score = book.Score                     // 评分
		dbBook.WordCount = book.WordCount             // 书籍总字数
		dbBook.ReadCount = book.ReadCount             // 阅读量
		dbBook.ThumbUrl = book.ThumbUrl

		db.Save(&book)
	} else {
		log.Printf("%s %d 上次更新时间与当前时间相差小于24小时或书籍已完结，不更新书籍\n", dbBook.Name, dbBook.BookID)
	}
	return dbBook
}

func GetBookByBookId(bookId int64) Book {
	var book Book
	result := db.Where("book_id = ?", bookId).First(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Printf("not find book in database,bookId: %d\n", bookId)
		return book
	}
	return book
}

func GetBookByBookName(bookName string) (book Book) {
	db.Where("name = ?", bookName).First(&book)
	return
}
