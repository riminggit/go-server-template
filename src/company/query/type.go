package companyQuery

import (
	"go-server-template/model/company"
)


type queryCompanyParams struct {
	Id           string `json:"id"`
	CompanyName string `json:"company_name"`
	IsUse        string `json:"is_use"`
}
type queryCompanyReturn struct {
	Code int `json:"code"`
	Data []companyModel.Company
}
