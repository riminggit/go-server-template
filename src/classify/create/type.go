package classifyCreate

type CreateParams struct {
	Data []ClassifyParams `json:"data"`
}

type CreateReturn struct {
	Code int    `json:"code"`
	Data []string `json:"data"`
}

type ClassifyParams struct {
	ClassifyName string `json:"classify_name"`
	ImgUrl       string `json:"img_url"`
	ImgSvg       string `json:"img_svg"`
	Rank         int `json:"rank"`
}
