package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ChapterAPI struct {
}

// GetContentByBookId 获取章节内容
func (a ChapterAPI) GetContentByBookId(c *gin.Context) {
	chapterIdStr := c.Query("chapter_id")
	chapterId, _ := strconv.ParseInt(chapterIdStr, 10, 64)
	chapter := chapterService.GetContentByChapterId(chapterId)
	if chapter.ChapterID == 0 || len(chapter.Content) < 128 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "get content error",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": chapter,
	})
}
