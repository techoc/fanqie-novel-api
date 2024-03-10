package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type ChapterAPI struct {
}

// GetContentByBookId 获取章节内容
func (a ChapterAPI) GetContentByBookId(c *gin.Context) {
	chapterIdStr := c.Query("chapter_id")
	chapterId, _ := strconv.ParseInt(chapterIdStr, 10, 64)
	chapter := chapterService.GetContentByChapterId(chapterId)
	if chapter.Content == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error",
			"data": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": chapter,
	})
}
