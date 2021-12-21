package companyQuery

import (
	"go-server-template/model/company"
)


type QueryCompanyParams struct {
	Id           string `json:"id"`
	CompanyName string `json:"company_name"`
	IsUse        string `json:"is_use"`
}
type QueryCompanyReturn struct {
	Code int `json:"code"`
	Data []companyModel.Company
}


type QueryCompanyMultipleParams struct {
	Id           []string `json:"id"`
	CompanyName  []string `json:"company_name"`
	IsUse        string `json:"is_use"`
}