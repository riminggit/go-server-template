package companyCreate

type CreateParams struct {
	Data []DataParams `json:"data"`
}

type CreateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}

type DataParams struct {
	CompanyName string `json:"company_name"`
	ImgUrl      string `json:"img_url"`
	ImgSvg      string `json:"img_svg"`
}
