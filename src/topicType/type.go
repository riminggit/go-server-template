package topicType

import "go-server-template/model/type"

type queryTypeParams struct {
	Id           string `json:"id"`
	ClassifyName string `json:"classify_name"`
	ClassifyId   string `json:"classify_id"`
	typeName     string `json:"type_name"`
	IsUse        string `json:"is_use"`
}
type queryReturn struct {
	Code int `json:"code"`
	Data []typeModel.Type
}
