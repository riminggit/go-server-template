package feedback

type CreateParams struct {
	TopicId        int    `json:"topic_id"`
	FeedbackType   int    `json:"feedback_type"`
	FacebackAnswer string `json:"faceback_answer"`
	Content        string `json:"content"`
	Grade          int    `json:"grade"`
}

type CommonReturn struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

type DeleteParams struct {
	ID       []string `json:"id"`
	UserId   []string `json:"user_id"`
	TopicId  []int    `json:"topic_id"`
	CreateAt []string `json:"create_at"` // 创建时间
	DeleteAt []string `json:"delete_at"`
	UpdateAt []string `json:"update_at"`
	RealDel  string   `json:"real_del"` // 1 真实删除  0 假删除
}

type UserDeleteParams struct {
	ID string `json:"id"`
}

type AnswerFeedbackParams struct {
	ID             int    `json:"id"`
	FacebackAnswer string `json:"faceback_answer"`
	IsRead         int    `json:"is_read"`
}

type UserUpdateParams struct {
	ID             int    `json:"id"`
	Content        string `json:"content"`
	Grade          int    `json:"grade"`
}
