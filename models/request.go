package models

type Request struct {
	UrlId  uint `db:"url_id" json:"url_id"`
	Result int  `db:"result" json:"result"`
}
