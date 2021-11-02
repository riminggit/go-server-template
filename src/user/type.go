package user

type ChangePhoneParams struct {
	Phone int `json:"phone" valid:"Mobile"`
}
type ChangePhoneReturnData struct {
	Code int `json:"code"`
}
