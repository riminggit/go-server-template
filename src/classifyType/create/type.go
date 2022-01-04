package classifyTypeCreate

type CreateParams struct {
	Data []DataParams `json:"data"`
}

type CreateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}

type DataParams struct {
	TypeName   string `json:"type_name"`
	ClassifyId int `json:"classify_id"`
	Rank       int    `json:"rank"`
}
