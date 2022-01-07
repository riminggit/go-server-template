package topicDelete

type DeleteParams struct {
	IDList  []string `json:"id_list"`
}

type DeleteReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}
