package classify

import "go-server-template/model/classify"

type queryClassifyParams struct {
	Id    string `json:"id"`
	ClassifyName  string `json:"classify_name"`
	Rank string `json:"rank"`
	IsUse string `json:"is_use"`
}
type queryClassifyReturn struct {
	Code int `json:"code"`
	Data []classifyModel.Classify
}
