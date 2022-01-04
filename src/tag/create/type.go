package tagCreate

type CreateParams struct {
	Data []DataParams `json:"data"`
}

type CreateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}

type DataParams struct {
	TagName string `json:"tag_name"`
}
