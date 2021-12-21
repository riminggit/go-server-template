package classifyTypeQuery

import "go-server-template/model/type"

type QueryTypeParams struct {
	Id           string `json:"id"`
	ClassifyName string `json:"classify_name"`
	ClassifyId   string `json:"classify_id"`
	TypeName     string `json:"type_name"`
	IsUse        string `json:"is_use"`
}
type QueryReturn struct {
	Code int `json:"code"`
	Data []typeModel.Type
}


type QueryTypeMultipleParams struct {
	Id           []string `json:"id"`
	ClassifyName []string `json:"classify_name"`
	ClassifyId   []string `json:"classify_id"`
	TypeName     []string `json:"type_name"`
	IsUse        string `json:"is_use"`
}