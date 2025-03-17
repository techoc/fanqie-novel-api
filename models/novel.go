package models

import "encoding/json"

func UnmarshalCtool(data []byte) (Ctool, error) {
	var r Ctool
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Ctool) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Ctool struct {
	Code    int64  `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type Data struct {
	AuditStatus   string        `json:"audit_status"`
	ChapterType   string        `json:"chapter_type"`
	Content       string        `json:"content"`
	GroupID       string        `json:"group_id"`
	ItemID        string        `json:"item_id"`
	NovelData     NovelData     `json:"novel_data"`
	PayStatus     DataPayStatus `json:"pay_status"`
	PlaySchemaURL string        `json:"play_schema_url"`
	SubGenre      string        `json:"sub_genre"`
	Title         string        `json:"title"`
}

type NovelData struct {
	AdFreeShow         string             `json:"ad_free_show"`
	AdShowNum          string             `json:"ad_show_num"`
	Author             string             `json:"author"`
	AuthorSchemaURL    string             `json:"author_schema_url"`
	AuthorThumbURL     string             `json:"author_thumb_url"`
	AutoPay            string             `json:"auto_pay"`
	BenefitTime        string             `json:"benefit_time"`
	BookID             string             `json:"book_id"`
	BookName           string             `json:"book_name"`
	Category           string             `json:"category"`
	CategoryID         int64              `json:"category_id"`
	ChapterTitle       string             `json:"chapter_title"`
	ChapterType        string             `json:"chapter_type"`
	ColumnSchemaURL    string             `json:"column_schema_url"`
	CopyrightInfo      string             `json:"copyright_info"`
	CreateTime         string             `json:"create_time"`
	CreationStatus     string             `json:"creation_status"`
	FeedRecommendText  string             `json:"feed_recommend_text"`
	Genre              string             `json:"genre"`
	GroupID            string             `json:"group_id"`
	InBookshelf        string             `json:"in_bookshelf"`
	IsAdBook           string             `json:"is_ad_book"`
	IsPraiseBook       string             `json:"is_praise_book"`
	Isbn               string             `json:"isbn"`
	ItemID             string             `json:"item_id"`
	ItemStatus         string             `json:"item_status"`
	LengthType         string             `json:"length_type"`
	LiteAdShowNum      string             `json:"lite_ad_show_num"`
	NeedPay            string             `json:"need_pay"`
	NextGroupID        string             `json:"next_group_id"`
	NextItemID         string             `json:"next_item_id"`
	NovelFreeStatus    string             `json:"novel_free_status"`
	Order              string             `json:"order"`
	OriginChapterTitle string             `json:"origin_chapter_title"`
	OriginalAuthorName string             `json:"original_author_name"`
	PaidBook           string             `json:"paid_book"`
	PayStatus          NovelDataPayStatus `json:"pay_status"`
	Platform           string             `json:"platform"`
	PreGroupID         string             `json:"pre_group_id"`
	PreItemID          string             `json:"pre_item_id"`
	RelatedAudioBookID string             `json:"related_audio_book_id"`
	Source             string             `json:"source"`
	SubGenre           string             `json:"sub_genre"`
	ThumbURL           string             `json:"thumb_url"`
	Title              string             `json:"title"`
	VipBook            string             `json:"vip_book"`
	WordNumber         string             `json:"word_number"`
}

type NovelDataPayStatus struct {
	AutoPayStatus string `json:"auto_pay_status"`
	Msg           string `json:"msg"`
	Status        string `json:"status"`
}

type DataPayStatus struct {
	Status string `json:"status"`
}
