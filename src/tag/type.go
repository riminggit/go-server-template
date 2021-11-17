package tag

import (
	"go-server-template/model/tag"
)

type queryTagParams struct {
	Id      string `json:"id"`
	TagName string `json:"tag_name"`
	IsUse   string `json:"is_use"`
}
type queryTagReturn struct {
	Code int `json:"code"`
	Data []tagModel.Tag
}
