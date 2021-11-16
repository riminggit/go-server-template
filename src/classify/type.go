package classify

import (
	"go-server-template/model/classify"
	"go-server-template/model/type"
)

type queryClassifyParams struct {
	Id           string `json:"id"`
	ClassifyName string `json:"classify_name"`
	Rank         string `json:"rank"`
	IsUse        string `json:"is_use"`
}
type queryClassifyReturn struct {
	Code int `json:"code"`
	Data []classifyModel.Classify
}


type queryClassifyAndTypeReturn struct {
	Code int `json:"code"`
	Data []queryClassifyAndTypeReturnData
}

type queryClassifyAndTypeReturnData struct {
	// classifyModel.Classify
	ID           int    `json:"id"` // id
	ClassifyName string `json:"classify_name"`                              // 类型名：html5、css、js等
	ImgUrl       string `json:"img_url"`                                      // 分类素材URL
	ImgSvg       string `json:"img_svg"`                                              // 素材svg
	Rank         int    `json:"rank"`                                           // 排序
	IsUse        int    `json:"is_use"` 
	TypeInfo []typeModel.Type `json:"typeInfo"`
}
