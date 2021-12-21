package tagQuery

import (
	"go-server-template/model/tag"
)

type QueryTagParams struct {
	Id      string `json:"id"`
	TagName string `json:"tag_name"`
	IsUse   string `json:"is_use"`
}
type QueryTagReturn struct {
	Code int `json:"code"`
	Data []tagModel.Tag
}

type QueryTagMultipleParams struct {
	Id      []string `json:"id"`
	TagName []string `json:"tag_name"`
	IsUse   string `json:"is_use"`
}