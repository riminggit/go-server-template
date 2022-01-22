package topicSetQuery

import (
	"go-server-template/model/topic"
	"time"
)

type QueryTopicSetParams struct {
	Id                 []string `json:"id"`
	Name               string   `json:"name"`
	TopicSetIdList     string   `json:"topic_set_id_list"`
	TopicSetDifficulty string   `json:"topic_set_difficulty"`
	TopicSetLevel      string   `json:"topic_set_level"`
	TopicType          string   `json:"topic_type"` // 练习：1  考试 ：2
	CreateAt           []string `json:"create_at"`
	DeleteAt           []string `json:"delete_at"`
	UpdateAt           []string `json:"update_at"`
	PageNum            int      `json:"pageNum"`
	PageSize           int      `json:"pageSize"`
	NeedTopicDetail    string   `json:"need_topic_detail"`
	IsUse              string   `json:"is_use"`
}
type queryTopicSetReturn struct {
	Code int                     `json:"code"`
	Data queryTopicSetReturnData `json:"data"`
}

type queryTopicSetReturnData struct {
	Data           []TopicInfoReturnData `json:"data"`
	PagingArgument PageArgument          `json:"pageArgument"`
}

type PageArgument struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

type queryTopicSetSimpleReturn struct {
	Code int                   `json:"code"`
	Data []TopicInfoReturnData `json:"data"`
}

type TopicInfoReturnData struct {
	ID                 int                `json:"id" `                  // id
	Name               string             `json:"name" `                // 套题名称
	TopicSetIdList     string             `json:"topic_set_id_list"`    // 题目id列表
	TopicSetDifficulty int                `json:"topic_set_difficulty"` // 题目难度，由所有的题目难度指数相加后除题目数
	TopicSetLevel      int                `json:"topic_set_level"`      // 题目等级，由所有的题目等级指数相加后除题目数
	TopicType          string             `json:"topic_type" `          // 练习：1  考试 ：2
	Remark             string             `json:"remark"`               // 备注
	CreateAt           time.Time          `json:"create_at"`            // 创建时间
	TopicList          []topicModel.Topic `json:"topic_list"`
}

type QueryTopicSetRandomParams struct {
	IsQueryByUserLevel string `json:"is_query_by_user_level"` // 0 否  1 是
}

type QueryTopicSetRandomReturn struct {
	Code int                 `json:"code"`
	Data TopicInfoReturnData `json:"data"`
}
