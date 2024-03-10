package models

import "gorm.io/gorm"

type Author struct {
	// 作者id
	AuthorId uint `json:"author_id"`
	// 作者名
	Name string `json:"name"`
	// 作者简介
	Desc string `json:"desc"`
	// 总创作天数
	TotalCreativeDays int `json:"total_creative_days"`
	// 总作品数
	TotalBooksCount int `json:"total_books_count"`
	// 粉丝数
	FansCount int `json:"fans_count"`
	// 总字数
	TotalNumberOfWords int `json:"total_number_of_words"`
	gorm.Model
}
