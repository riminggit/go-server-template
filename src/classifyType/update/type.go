package classifyTypeUpdate

type UpdateParams struct {
	Data []DataParams `json:"data"`
}

type UpdateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}

type DataParams struct {
	ID         int    `json:"id"` 
	TypeName   string `json:"type_name"`
	ClassifyId int    `json:"classify_id"`
	Rank       int    `json:"rank"`
}
