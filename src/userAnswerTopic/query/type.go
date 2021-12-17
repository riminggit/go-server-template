package userAnswerTopicQuery

import (
	"go-server-template/model/user"
)

type PageArgument struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

type queryUserAnswerTopicParams struct {
	Id               string   `json:"id"`
	TopicIdList      string   `json:"topic_id_list"` // 题目标题
	TopicSetId       string   `json:"topic_set_id"`
	TopicSetName     string   `json:"topic_set_name"`
	AnswerStart      []string `json:"answer_start"`
	AnswerEnd        []string `json:"answer_end"`
	AnswerNum        string   `json:"answer_num"`
	AnswerCorrectNum string   `json:"answer_correct_num" `
	IsAchieve        string   `json:"is_achieve"`
	TopicDifficulty  string   `json:"topic_difficulty"`
	TopicLevel       string   `json:"topic_level"`
	CreateAt         []string `json:"create_at"`
	DeleteAt         []string `json:"delete_at"`
	UpdateAt         []string `json:"update_at"`
	IsUse            string   `json:"is_use"`
	PageNum          int      `json:"pageNum"`
	PageSize         int      `json:"pageSize"`
}
type queryUserAnswerTopicReturn struct {
	Code int                       `json:"code"`
	Data UserAnswerTopicReturnData `json:"data"`
}

type UserAnswerTopicReturnData struct {
	Data           []userModel.UserAnswerTopicRecord `json:"data"`
	PagingArgument PageArgument                      `json:"pageArgument"`
}
