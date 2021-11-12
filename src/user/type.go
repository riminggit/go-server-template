package user

type QueryUserParams struct {
	Email    string   `json:"email"`
	NickName string   `json:"nick_name"`
	Phone    string   `json:"phone"`
	ComeFrom string   `json:"come_from"`
	Gender   string   `json:"gender"`
	City     []string `json:"city"`
	Province []string `json:"province"`
	Country  []string `json:"country"`
	Language []string `json:"language"`
	IsAdmin  string   `json:"is_admin"`
	Birthday []string `json:"birthday"`
	CreateAt []string `json:"create_at"`
	DeleteAt []string `json:"delete_at"`
	UpdateAt []string `json:"update_at"`
	IsUse    string   `json:"is_use"`
	PageNum  int      `json:"pageNum"`
	PageSize int      `json:"pageSize"`
}

type QueryUserData struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	NickName  string `json:"nick_name"`
	Phone     string `json:"phone"`
	AvatarUrl string `json:"avatar_url"`
	ComeFrom  string `json:"come_from"`
	Gender    string `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Language  string `json:"language"`
	IsAdmin   string `json:"is_admin"`
	Birthday  string `json:"birthday"`
	CreateAt  string `json:"create_at"`
	UpdateAt  string `json:"update_at"`
}

type PageArgument struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}
type QueryUserReturnData struct {
	Code int            `json:"code"`
	Data UserReturnData `json:"data"`
}
type UserReturnData struct {
	Data           []QueryUserData `json:"data"`
	PagingArgument PageArgument    `json:"pageArgument"`
}
