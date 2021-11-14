package classify

type queryClassifyParams struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type queryClassifyReturn struct {
	Token    string         `json:"token"`
	Code     int            `json:"code"`
}
