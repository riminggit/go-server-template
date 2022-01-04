package classifyTypeDelete

type DeleteParams struct {
	IDList   []string `json:"id_list"`
	NameList []string `json:"name_list"`
	RealDel  string   `json:"real_del"` // 1 真实删除  0 假删除
}

type DeleteReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}
