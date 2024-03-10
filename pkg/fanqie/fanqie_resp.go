package fanqie

type Response struct {
	Code    int64  `json:"code"`
	Data    Data   `json:"data"`
	LogID   string `json:"log_id"`
	Message string `json:"message"`
	Now     int64  `json:"now"`
}

type Data struct {
	AllItemIDS            []string                  `json:"allItemIds"`
	AuditStatus           string                    `json:"audit_status"`
	ChapterType           string                    `json:"chapter_type"`
	ChapterListWithVolume [][]ChapterListWithVolume `json:"chapterListWithVolume"`
	Content               string                    `json:"content"`
	ExactlyMatch          bool                      `json:"exactly_match"`
	GroupID               string                    `json:"group_id"`
	HasMore               bool                      `json:"has_more"`
	ItemID                string                    `json:"item_id"`
	NewRechargeUser       string                    `json:"new_recharge_user"`
	NovelData             NovelData                 `json:"novel_data"`
	Offset                int64                     `json:"offset"`
	PayStatus             DataPayStatus             `json:"pay_status"`
	RetData               []RetDatum                `json:"ret_data"`
	SearchAuthorDataList  []string                  `json:"search_author_data_list"`
	SearchBookDataList    []SearchBookDataList      `json:"search_book_data_list"`
	SearchID              string                    `json:"search_id"`
	Strategy              Strategy                  `json:"strategy"`
	Title                 string                    `json:"title"`
	TotalCount            int64                     `json:"total_count"`
	VolumeNameList        []string                  `json:"volumeNameList"`
}

type ChapterListWithVolume struct {
	FirstPassTime     string `json:"firstPassTime"`
	IsChapterLock     bool   `json:"isChapterLock"`
	IsPaidPublication bool   `json:"isPaidPublication"`
	IsPaidStory       bool   `json:"isPaidStory"`
	ItemID            string `json:"itemId"`
	NeedPay           int64  `json:"needPay"`
	Title             string `json:"title"`
	VolumeName        string `json:"volume_name"`
}

type NovelData struct {
	Abstract                 string             `json:"abstract"`
	AdFreeShow               string             `json:"ad_free_show"`
	AdShowNum                string             `json:"ad_show_num"`
	Author                   string             `json:"author"`
	AuthorSchemaURL          string             `json:"author_schema_url"`
	AuthorThumbURL           string             `json:"author_thumb_url"`
	AutoPay                  string             `json:"auto_pay"`
	BasePrice                string             `json:"base_price"`
	BenefitTime              string             `json:"benefit_time"`
	BookID                   string             `json:"book_id"`
	BookName                 string             `json:"book_name"`
	BookshelfURL             string             `json:"bookshelf_url"`
	Category                 string             `json:"category"`
	CategoryID               int64              `json:"category_id"`
	ChapterTitle             string             `json:"chapter_title"`
	ChapterType              string             `json:"chapter_type"`
	ColumnSchemaURL          string             `json:"column_schema_url"`
	CopyrightInfo            string             `json:"copyright_info"`
	CreateTime               string             `json:"create_time"`
	CreationStatus           string             `json:"creation_status"`
	CustomTotalPrice         string             `json:"custom_total_price"`
	DiscountCustomTotalPrice string             `json:"discount_custom_total_price"`
	FeedRecommendText        string             `json:"feed_recommend_text"`
	Genre                    string             `json:"genre"`
	GroupID                  string             `json:"group_id"`
	InBookshelf              string             `json:"in_bookshelf"`
	IsAdBook                 string             `json:"is_ad_book"`
	IsPraiseBook             string             `json:"is_praise_book"`
	Isbn                     string             `json:"isbn"`
	ItemID                   string             `json:"item_id"`
	ItemStatus               string             `json:"item_status"`
	LengthType               string             `json:"length_type"`
	LiteAdShowNum            string             `json:"lite_ad_show_num"`
	NeedPay                  string             `json:"need_pay"`
	NextGroupID              string             `json:"next_group_id"`
	NextItemID               string             `json:"next_item_id"`
	NovelFreeStatus          string             `json:"novel_free_status"`
	Order                    string             `json:"order"`
	OriginChapterTitle       string             `json:"origin_chapter_title"`
	OriginalAuthorName       string             `json:"original_author_name"`
	PaidBook                 string             `json:"paid_book"`
	PayStatus                NovelDataPayStatus `json:"pay_status"`
	Platform                 string             `json:"platform"`
	PreGroupID               string             `json:"pre_group_id"`
	PreItemID                string             `json:"pre_item_id"`
	RelatedAudioBookID       string             `json:"related_audio_book_id"`
	SaleStatus               string             `json:"sale_status"`
	SaleType                 string             `json:"sale_type"`
	SerialCount              string             `json:"serial_count"`
	Source                   string             `json:"source"`
	Status                   string             `json:"status"`
	SubGenre                 any                `json:"sub_genre"`
	ThumbURL                 string             `json:"thumb_url"`
	Title                    string             `json:"title"`
	VipBook                  string             `json:"vip_book"`
	WordNumber               string             `json:"word_number"`
}

type NovelDataPayStatus struct {
	AutoPayStatus string `json:"auto_pay_status"`
	Msg           string `json:"msg"`
	Status        string `json:"status"`
}

type DataPayStatus struct {
	Status string `json:"status"`
}

type RetDatum struct {
	Abstract          string `json:"abstract"`
	AddBookshelfCount string `json:"add_bookshelf_count"`
	AudioThumbURI     string `json:"audio_thumb_uri"`
	Author            string `json:"author"`
	BookID            string `json:"book_id"`
	Category          string `json:"category"`
	CreationStatus    string `json:"creation_status"`
	Genre             string `json:"genre"`
	ItemSchemaURL     string `json:"item_schema_url"`
	PageURL           string `json:"page_url"`
	Score             string `json:"score"`
	ThumbURL          string `json:"thumb_url"`
	Title             string `json:"title"`
	ToutiaoClickRate  string `json:"toutiao_click_rate"`
}

type SearchBookDataList struct {
	Author           string `json:"author"`
	BookAbstract     string `json:"book_abstract"`
	BookID           string `json:"book_id"`
	BookName         string `json:"book_name"`
	Category         string `json:"category"`
	CreationStatus   int64  `json:"creation_status"`
	FirstChapterID   string `json:"first_chapter_id"`
	LastChapterID    string `json:"last_chapter_id"`
	LastChapterTime  string `json:"last_chapter_time"`
	LastChapterTitle string `json:"last_chapter_title"`
	ReadCount        int64  `json:"read_count"`
	ThumbURI         string `json:"thumb_uri"`
	ThumbURL         string `json:"thumb_url"`
	WordCount        int64  `json:"word_count"`
}

type Strategy struct {
	NewReaderTextSize int64 `json:"new_reader_text_size"`
}

func (a SearchBookDataList) isEmpty() bool {
	if a.BookName == "" || a.BookID == "" {
		return true
	}
	return false
}
