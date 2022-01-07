package companyUpdate

type UpdateParams struct {
	Data []DataParams `json:"data"`
}

type UpdateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}

type DataParams struct {
	ID          int    `json:"id"`
	CompanyName string `json:"company_name"`
	ImgUrl      string `json:"img_url"`
	ImgSvg      string `json:"img_svg"`
}
