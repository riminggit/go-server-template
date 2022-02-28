package userCreate

type UserRegisterParams struct {
	ID        int    `json:"id"` // id
	ComeFrom  string `json:"come_from"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Language  string `json:"language"`
	Email     string `json:"email"`
	Password  string `json:"password" valid:"Required;"`
	Phone     string `json:"phone"`
	UpdateAt  int64  `json:"update_at"`
	CreateAt  int64  `json:"create_at"`
}
