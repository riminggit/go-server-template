package classifyUpdate

type UpdateParams struct {
	Data []ClassifyParams `json:"data"`
}

type UpdateReturn struct {
	Code int    `json:"code"`
	Data []string `json:"data"`
}

type ClassifyParams struct {
	ID           int    `json:"id"` // id
	ClassifyName string `json:"classify_name"`
	ImgUrl       string `json:"img_url"`
	ImgSvg       string `json:"img_svg"`
	Rank         int    `json:"rank"`
}
