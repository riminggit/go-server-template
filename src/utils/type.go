package utils

type AddParams struct {
	Color []string `json:"color"`
}

type CreateReturn struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type queryColorReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}

type ColorData struct {
	ID    int    `json:"id"`    // id
	Color string `json:"color"` // 颜色
}
