package fanqie

import (
	"github.com/imroc/req/v3"
	"github.com/techoc/fanqie-novel-api/models"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Fanqie struct {
}

func Search(keywords string, page int, isAuthor bool) []models.Book {

	if isAuthor {
		req.SetQueryParams(map[string]string{
			"keyword": keywords,
			"page":    "1",
			"sort":    "author",
		})
	} else {
		return searchByTitle(keywords)
	}
	return nil
}

func searchByTitle(title string) []models.Book {
	var data Response
	resp, err := req.SetQueryParams(map[string]string{
		"q":   title,
		"aid": "1967",
	}).
		SetSuccessResult(&data).
		Get("https://novel.snssdk.com/api/novel/channel/homepage/search/search/v1/")

	if err != nil {
		panic(err)
	}

	// Status code is between 200 and 299.
	if resp.IsSuccessState() && len(data.Data.RetData) > 0 {
		size := len(data.Data.RetData)

		books := make([]models.Book, size)
		for i := 0; i < size; i++ {
			bookInfo := data.Data.RetData[i]
			bookId, _ := strconv.ParseInt(bookInfo.BookID, 10, 64)
			floatScore, _ := strconv.ParseFloat(bookInfo.Score, 32)
			book := models.Book{
				BookID:         bookId,
				Name:           bookInfo.Title,
				Author:         bookInfo.Author,
				Desc:           bookInfo.Abstract,
				CreationStatus: bookInfo.CreationStatus,
				Category:       bookInfo.Category,
				Score:          float32(floatScore),
			}
			// 进行二轮查找 获取所需信息
			bookData := searchByTitleTwo(bookInfo.Title, bookInfo.Author, bookId)
			if !bookData.isEmpty() {
				log.Printf("%s 二次查找成功，书籍Id %s\n", bookInfo.Title, bookInfo.BookID)
				// 更新书籍信息
				//book.ChaptersCount:   0
				book.FirstChapterId, _ = strconv.ParseInt(bookData.FirstChapterID, 10, 64)
				book.LastChapterId, _ = strconv.ParseInt(bookData.LastChapterID, 10, 64)
				book.LastChapterTime, _ = strconv.ParseInt(bookData.LastChapterTime, 10, 64)
				book.WordCount = bookData.WordCount
				book.ReadCount = bookData.ReadCount
				book.ThumbUrl = bookData.ThumbURL
			} else {
				log.Printf("%s 二次查找失败，书籍Id %s\n", bookInfo.Title, bookInfo.BookID)
			}
			books[i] = book
		}

		return books
	}
	return nil
}

func searchByTitleTwo(title string, author string, bookId int64) SearchBookDataList {
	var data Response
	resp, err := req.SetQueryParams(map[string]string{
		"filter":     "127,127,127,127",
		"page_count": "10",
		"page_index": "0",
		"query_type": "0",
		"query_word": title,
	}).
		SetSuccessResult(&data).
		Get("https://fanqienovel.com/api/author/search/search_book/v1")

	if err != nil {
		panic(err)
	}

	// Status code is between 200 and 299.
	if resp.IsSuccessState() && len(data.Data.SearchBookDataList) > 0 {
		bookDataList := data.Data.SearchBookDataList
		for i := 0; i < len(bookDataList); i++ {
			bookData := bookDataList[i]

			parseInt, err := strconv.ParseInt(bookData.BookID, 10, 64)
			if err != nil {
				return SearchBookDataList{}
			}
			// 校验书籍ID 和作者名
			// 书籍名可能不一致
			if parseInt == bookId {
				return bookData
			}
		}
	}
	return SearchBookDataList{}
}

func GetDirectoryByBookId(bookId int64) []models.Chapter {
	var data Response
	resp, err := req.SetQueryParams(map[string]string{
		"bookId": strconv.FormatInt(bookId, 10),
	}).
		SetSuccessResult(&data).
		Get("https://fanqienovel.com/api/reader/directory/detail")
	if err != nil {
		panic(err)
	}
	if resp.IsSuccessState() && len(data.Data.ChapterListWithVolume) > 0 {
		// 获取目录
		directories := data.Data.ChapterListWithVolume
		var chapterList []models.Chapter
		for _, directory := range directories {
			for _, chapter := range directory {
				chapterId, _ := strconv.ParseInt(chapter.ItemID, 10, 64)

				myChapter := models.Chapter{
					BookId:    bookId,
					Name:      chapter.Title,
					ChapterID: chapterId,
				}
				chapterList = append(chapterList, myChapter)
			}
		}
		// 填入上一章ID和下一章ID
		log.Println("获取目录成功，开始填充上一章ID和下一章ID")
		for i, _ := range chapterList {
			if i == 0 { // 第一章没有上一章ID
				chapterList[i].PreGroupID = 0
				chapterList[i].PreItemID = 0
				chapterList[i].NextGroupID = chapterList[i+1].ChapterID
				chapterList[i].NextItemID = chapterList[i+1].ChapterID
			} else if i == len(chapterList)-1 { // 最后一章没有下一章ID
				chapterList[i].PreGroupID = chapterList[i-1].ChapterID
				chapterList[i].PreItemID = chapterList[i-1].ChapterID
				chapterList[i].NextGroupID = 0
				chapterList[i].NextItemID = 0
			} else {
				chapterList[i].PreGroupID = chapterList[i-1].ChapterID
				chapterList[i].PreItemID = chapterList[i-1].ChapterID
				chapterList[i].NextGroupID = chapterList[i+1].ChapterID
				chapterList[i].NextItemID = chapterList[i+1].ChapterID
			}
		}
		return chapterList
	}
	return nil
}

func GetContentByChapterId(chapterId int64) models.Chapter {
	var data Response
	response, err := req.SetQueryParams(map[string]string{
		"device_platform":  "android",
		"parent_enterfrom": "novel_channel_search.tab.",
		"aid":              "2329",
		"platform_id":      "1967",
		"group_id":         strconv.FormatInt(chapterId, 10),
		"item_id":          strconv.FormatInt(chapterId, 10),
	}).
		SetSuccessResult(&data).
		Get("https://novel.snssdk.com/api/novel/book/reader/full/v1/")
	if err != nil {
		log.Printf("%d 请求正文失败\n", chapterId)
		log.Printf("%s \n", err)
		return models.Chapter{}
	}

	if response.IsSuccessState() && data.Data.Content != "" {
		bookId, _ := strconv.ParseInt(data.Data.NovelData.BookID, 10, 64)
		chapterNumber, _ := strconv.Atoi(data.Data.NovelData.Order)
		wordCount, _ := strconv.Atoi(data.Data.NovelData.WordNumber)

		var preItemID, preGroupID int64
		if data.Data.NovelData.PreItemID != "" {
			preItemID, _ = strconv.ParseInt(data.Data.NovelData.PreItemID, 10, 64)
			preGroupID, _ = strconv.ParseInt(data.Data.NovelData.PreGroupID, 10, 64)
		}
		var nextItemID, nextGroupID int64
		if data.Data.NovelData.NextItemID != "" {
			nextItemID, _ = strconv.ParseInt(data.Data.NovelData.NextItemID, 10, 64)
			nextGroupID, _ = strconv.ParseInt(data.Data.NovelData.NextGroupID, 10, 64)
		}

		// 提取正文
		re := regexp.MustCompile(`<article>([\s\S]*?)</article>`)
		submatch := re.FindStringSubmatch(data.Data.Content)
		content := submatch[1]
		// 将 <p> 标签替换为换行符
		content = strings.Replace(content, "<p>", "\n", -1)
		// 去除剩下所有的 html 标签
		re = regexp.MustCompile(`</?\w+>`)
		replaceAllString := re.ReplaceAllString(content, "")

		chapter := models.Chapter{
			ChapterID:     chapterId,
			Name:          data.Data.Title,
			BookId:        bookId,
			WordCount:     wordCount,
			Content:       replaceAllString,
			ChapterNumber: chapterNumber,
			NextGroupID:   nextGroupID,
			NextItemID:    nextItemID,
			PreGroupID:    preGroupID,
			PreItemID:     preItemID,
		}
		return chapter
	}
	return models.Chapter{}
}
