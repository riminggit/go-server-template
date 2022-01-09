package feedback

import (
	"go-server-template/model/user"
)

type QueryParams struct {
	Id           []string `json:"id"`
	UserId       []string `json:"user_id"`
	TopicId      int      `json:"topic_id"`
	FeedbackType int      `json:"feedback_type"`
	Grade        int      `json:"grade"`
	CreateAt     []string `json:"create_at"` // 创建时间
	DeleteAt     []string `json:"delete_at"`
	UpdateAt     []string `json:"update_at"`
	PageNum      int      `json:"pageNum"`
	PageSize     int      `json:"pageSize"`
	IsUse        string   `json:"is_use"`
}

type UserQueryParams struct {
	Id           []string `json:"id"`
	TopicId      int      `json:"topic_id"`
	FeedbackType int      `json:"feedback_type"`
	Grade        int      `json:"grade"`
	CreateAt     []string `json:"create_at"` // 创建时间
	DeleteAt     []string `json:"delete_at"`
	UpdateAt     []string `json:"update_at"`
	IsUse        string   `json:"is_use"`
}

type CreateParams struct {
	TopicId        int    `json:"topic_id"`
	FeedbackType   int    `json:"feedback_type"`
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
	ID []string `json:"id"`
}

type AnswerFeedbackParams struct {
	ID             int    `json:"id"`
	FacebackAnswer string `json:"faceback_answer"`
	IsRead         int    `json:"is_read"`
}

type UserUpdateParams struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Grade   int    `json:"grade"`
}

type queryPaddingReturn struct {
	Code int                    `json:"code"`
	Data queryPaddingReturnData `json:"data"`
}

type queryPaddingReturnData struct {
	Data           []userModel.UserFeedback `json:"data"`
	PagingArgument PageArgument             `json:"pageArgument"`
}

type PageArgument struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

type CommonQueryReturn struct {
	Code int                      `json:"code"`
	Data []userModel.UserFeedback `json:"data"`
}
