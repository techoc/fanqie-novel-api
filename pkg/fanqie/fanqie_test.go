package fanqie

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {

	//re := regexp.MustCompile(`\d+`)
	//reStrings := re.FindAllString("第 1 章 初临火影，先选个妈！", -1)
	//chapterNumber, _ := strconv.Atoi(reStrings[0])
	//fmt.Println(chapterNumber)

	//books := Search("火影：绳树都凉了，你也能救活？", 1, false)
	////searchByTitleTwo("火影：绳树都凉了，你也能救活？", "日漫", 0)
	//for _, book := range books {
	//	fmt.Println(book.BookID)
	//	fmt.Println(book.Name)
	//	fmt.Println(book.LastChapterId)
	//	chapters := GetDirectoryByBookId(book.BookID)
	//	fmt.Println("章节数", len(chapters))
	//	fmt.Println("---")
	//}

	//chapter := GetContentByChapterId(7058190355161842180)
	//fmt.Printf("%v\n", chapter)

}

func TestGetContentByChapterIdV2(t *testing.T) {
	chapter := GetContentByChapterIdV2(7258562472163049984)
	fmt.Println(chapter.Content)
}

func TestGetNewCookie(t *testing.T) {
	GetNewCookie()
}
