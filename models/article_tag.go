package models

type ArticleTag struct {
	ID        int `json:"id"`
	ArticleId int `json:"article_id"`
	TagId     int `json:"tag_id"`
}
