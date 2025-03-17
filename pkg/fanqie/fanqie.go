package fanqie

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/techoc/fanqie-novel-api/models"
	"gorm.io/gorm"
	"io"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
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

// GetContentByChapterId
// 通过API获取章节内容
// 已失效
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
		log.Printf("%d request content failed\n", chapterId)
		log.Printf("%s \n", err)
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

type NovelDownloader struct {
	CODE    [][2]int
	charset [][]rune
}

func (nd *NovelDownloader) decodeContent(content string, mode int) string {
	result := ""
	for _, char := range content {
		uni := int(char)
		if nd.CODE[mode][0] <= uni && uni <= nd.CODE[mode][1] {
			bias := uni - nd.CODE[mode][0]
			if 0 <= bias && bias < len(nd.charset[mode]) && nd.charset[mode][bias] != '?' {
				result += string(nd.charset[mode][bias])
			} else {
				result += string(char)
			}
		} else {
			result += string(char)
		}
	}
	return result
}

// GetContentByChapterIdV2
// 通过解码字符获取章节内容
func GetContentByChapterIdV2(chapterId int64) models.Chapter {
	// 发送HTTP请求获取网页内容
	//url := fmt.Sprintf("https://fanqienovel.com/reader/%d", chapterId)
	url := fmt.Sprintf("http://rehaofan.jingluo.love/content?item_id=%d", chapterId)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch the URL: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch the URL: %v", resp.Status)
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	ctool, err := models.UnmarshalCtool(bytes)
	if err != nil {
		log.Fatalf("Failed to unmarshal response body: %v", err)
	}
	content := ctool.Data.Content
	// 读取响应内容
	// 解析HTML
	//doc, err := xmlquery.Parse(resp.Body)
	//if err != nil {
	//	log.Fatalf("Failed to parse the HTML: %v", err)
	//}
	// 使用XPath提取内容
	//xpath := "//div[@class='muye-reader-content noselect']//p/text()"
	//node := xmlquery.FindOne(doc, xpath)
	//content := node.Data

	//content := "\\u003cp\\u003e\uE4E3\uE510\uE4EA\uE4F3争吵仍\uE3E9持续。\\u003c/p\\u003e\\u003cp\\u003e唐散：“\uE490唐\uE3EC弱\uE508，\uE490唐\uE3EC恶劣\uE4F3\uE3EC族\uE4FE\uE4F3\uE41E！\uE478唯独\uE52A邀请唐\uE3EC，\uE487\uE55A叶\uE3EC\uE51E\uE4C3\uE41E搞针\uE49A！\uE3EB殊\uE49A待！”\\u003c/p\\u003e\\u003cp\\u003e叶\uE473：“\uE4EB资\uE403阶\uE4F3闭嘴。”\\u003c/p\\u003e\\u003cp\\u003e唐散：“\uE511\uE44A\uE436，\uE4C3算\uE521唐\uE3EC\uE415\uE4F3\uE480错，抛\uE4FF\uE483\uE51B\uE52A谈，\uE452\uE459\uE487\uE55A叶\uE3EC\uE4C3\uE4DE\uE4A8错\uE46A\uE480\uE436？”\\u003c/p\\u003e\\u003cp\\u003e叶\uE473：“\uE4EB资\uE403阶\uE4F3闭嘴。”\\u003c/p\\u003e\\u003cp\\u003e唐散：“既\uE3EE\uE487\uE55A叶\uE3EC\uE41A仗\uE444权\uE4FD\uE53E\uE431随\uE49F针\uE49A唐\uE3EC，\uE417\uE521\uE477，\uE4EF\uE48E\uE4F3\uE3EC族估计\uE548\uE46A"

	nd := NovelDownloader{
		CODE: [][2]int{{58344, 58715}, {58345, 58716}},
		charset: [][]rune{
			{'D', '在', '主', '特', '家', '军', '然', '表', '场', '4', '要', '只', 'v', '和', '?', '6', '别', '还', 'g', '现', '儿', '岁', '?', '?', '此', '象', '月', '3', '出', '战', '工', '相', 'o', '男', '直', '失', '世', 'F', '都', '平', '文', '什', 'V', 'O', '将', '真', 'T', '那', '当', '?', '会', '立', '些', 'u', '是', '十', '张', '学', '气', '大', '爱', '两', '命', '全', '后', '东', '性', '通', '被', '1', '它', '乐', '接', '而', '感', '车', '山', '公', '了', '常', '以', '何', '可', '话', '先', 'p', 'i', '叫', '轻', 'M', '士', 'w', '着', '变', '尔', '快', 'l', '个', '说', '少', '色', '里', '安', '花', '远', '7', '难', '师', '放', 't', '报', '认', '面', '道', 'S', '?', '克', '地', '度', 'I', '好', '机', 'U', '民', '写', '把', '万', '同', '水', '新', '没', '书', '电', '吃', '像', '斯', '5', '为', 'y', '白', '几', '日', '教', '看', '但', '第', '加', '候', '作', '上', '拉', '住', '有', '法', 'r', '事', '应', '位', '利', '你', '声', '身', '国', '问', '马', '女', '他', 'Y', '比', '父', 'x', 'A', 'H', 'N', 's', 'X', '边', '美', '对', '所', '金', '活', '回', '意', '到', 'z', '从', 'j', '知', '又', '内', '因', '点', 'Q', '三', '定', '8', 'R', 'b', '正', '或', '夫', '向', '德', '听', '更', '?', '得', '告', '并', '本', 'q', '过', '记', 'L', '让', '打', 'f', '人', '就', '者', '去', '原', '满', '体', '做', '经', 'K', '走', '如', '孩', 'c', 'G', '给', '使', '物', '?', '最', '笑', '部', '?', '员', '等', '受', 'k', '行', '一', '条', '果', '动', '光', '门', '头', '见', '往', '自', '解', '成', '处', '天', '能', '于', '名', '其', '发', '总', '母', '的', '死', '手', '入', '路', '进', '心', '来', 'h', '时', '力', '多', '开', '已', '许', 'd', '至', '由', '很', '界', 'n', '小', '与', 'Z', '想', '代', '么', '分', '生', '口', '再', '妈', '望', '次', '西', '风', '种', '带', 'J', '?', '实', '情', '才', '这', '?', 'E', '我', '神', '格', '长', '觉', '间', '年', '眼', '无', '不', '亲', '关', '结', '0', '友', '信', '下', '却', '重', '己', '老', '2', '音', '字', 'm', '呢', '明', '之', '前', '高', 'P', 'B', '目', '太', 'e', '9', '起', '稜', '她', '也', 'W', '用', '方', '子', '英', '每', '理', '便', '四', '数', '期', '中', 'C', '外', '样', 'a', '海', '们', '任'}, // 假设的字符集
			{'s', '?', '作', '口', '在', '他', '能', '并', 'B', '士', '4', 'U', '克', '才', '正', '们', '字', '声', '高', '全', '尔', '活', '者', '动', '其', '主', '报', '多', '望', '放', 'h', 'w', '次', '年', '?', '中', '3', '特', '于', '十', '入', '要', '男', '同', 'G', '面', '分', '方', 'K', '什', '再', '教', '本', '己', '结', '1', '等', '世', 'N', '?', '说', 'g', 'u', '期', 'Z', '外', '美', 'M', '行', '给', '9', '文', '将', '两', '许', '张', '友', '0', '英', '应', '向', '像', '此', '白', '安', '少', '何', '打', '气', '常', '定', '间', '花', '见', '孩', '它', '直', '风', '数', '使', '道', '第', '水', '已', '女', '山', '解', 'd', 'P', '的', '通', '关', '性', '叫', '儿', 'L', '妈', '问', '回', '神', '来', 'S', ' ', '四', '望', '前', '国', '些', 'O', 'v', 'l', 'A', '心', '平', '自', '无', '军', '光', '代', '是', '好', '却', 'c', '得', '种', '就', '意', '先', '立', 'z', '子', '过', 'Y', 'j', '表', ' ', '么', '所', '接', '了', '名', '金', '受', 'J', '满', '眼', '没', '部', '那', 'm', '每', '车', '度', '可', 'R', '斯', '经', '现', '门', '明', 'V', '如', '走', '命', 'y', '6', 'E', '战', '很', '上', 'f', '月', '西', '7', '长', '夫', '想', '话', '变', '海', '机', 'x', '到', 'W', '一', '成', '生', '信', '笑', '但', '父', '开', '内', '东', '马', '日', '小', '而', '后', '带', '以', '三', '几', '为', '认', 'X', '死', '员', '目', '位', '之', '学', '远', '人', '音', '呢', '我', 'q', '乐', '象', '重', '对', '个', '被', '别', 'F', '也', '书', '稜', 'D', '写', '还', '因', '家', '发', '时', 'i', '或', '住', '德', '当', 'o', 'l', '比', '觉', '然', '吃', '去', '公', 'a', '老', '亲', '情', '体', '太', 'b', '万', 'C', '电', '理', '?', '失', '力', '更', '拉', '物', '着', '原', '她', '工', '实', '色', '感', '记', '看', '出', '相', '路', '大', '你', '候', '2', '和', '?', '与', 'p', '样', '新', '只', '便', '最', '不', '进', 'T', 'r', '做', '格', '母', '总', '爱', '身', '师', '轻', '知', '往', '加', '从', '?', '天', 'e', 'H', '?', '听', '场', '由', '快', '边', '让', '把', '任', '8', '条', '头', '事', '至', '起', '点', '真', '手', '这', '难', '都', '界', '用', '法', 'n', '处', '下', '又', 'Q', '告', '地', '5', 'k', 't', '岁', '有', '会', '果', '利', '民'},
		},
	}
	mode := 0 // 不知道什么时候才用这个参数
	decodedContent := nd.decodeContent(content, mode)
	chapter := models.Chapter{
		ChapterID:     chapterId,
		Name:          "",
		BookId:        0,
		ChapterNumber: 0,
		WordCount:     0,
		Content:       decodedContent,
		NextGroupID:   0,
		NextItemID:    0,
		PreGroupID:    0,
		PreItemID:     0,
		Model:         gorm.Model{},
	}
	return chapter
}

// 假设 _downloadChapterContent 是一个已经定义好的方法
func DownloadChapterContent(chapterID int, testMode bool) (string, error) {
	// 实现下载章节内容的逻辑
	return "", nil // 这里返回一个占位符
}

func GetNewCookie() {
	// Generate new cookie
	bas := big.NewInt(1000000000000000000)
	// 设置随机数种子，以当前时间为依据，确保每次运行生成不同的随机序列
	rand.Seed(time.Now().UnixNano())

	// 生成范围起始值和结束值对应的 *big.Int 类型
	start := big.NewInt(0).Mul(bas, big.NewInt(6))
	end := big.NewInt(0).Mul(bas, big.NewInt(8))
	max := big.NewInt(0).Mul(bas, big.NewInt(9))

	// 生成一个在 [start, end] 区间内的随机整数作为起始值
	r := big.NewInt(0).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), big.NewInt(0).Sub(end, start))
	start.Add(start, r)

	// 循环打印范围内的整数
	for i := new(big.Int).Set(start); i.Cmp(max) < 0; i.Add(i, big.NewInt(1)) {
		fmt.Println(i)
	}
}
