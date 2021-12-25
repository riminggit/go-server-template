package topicQuery

import (
	"go-server-template/model/classify"
	"go-server-template/model/company"
	"go-server-template/model/tag"
	"go-server-template/model/topic"
	"go-server-template/model/type"
)

type PageArgument struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

type queryTopicParams struct {
	Id               []string `json:"id"`
	Title            string   `json:"title"` // 题目标题
	QuestionType     []string `json:"question_type"`
	Degree           []string `json:"degree"`              // 难度，简单：0，中等：1，难：2，极难：3
	Level            []string `json:"level"`               // 1 初级 ， 2 中级 ， 3 高级 ， 4 资深 ，5 专家 ， 6 资深专家 ， 7 研究员
	IsBaseTopic      string   `json:"is_base_topic"`       // 0 否  1是 ，是否是基础题
	IsImportantTopic string   `json:"is_important_topic" ` // 0 否  1是 ，是否是重点题
	ComeFrom         string   `json:"come_from"`           // 来源
	ClassifyId       string   `json:"classify_id"`         // 类型id
	CompanyId        string   `json:"company_id"`          // 公司名
	TagId            string   `json:"tag_id"`              // 标签id
	TypeId           string   `json:"type_id"`             // 分类id
	TopicSetId       string   `json:"topic_set_id"`        // 套题id
	CreateAt         []string `json:"create_at"`
	DeleteAt         []string `json:"delete_at"`
	UpdateAt         []string `json:"update_at"`
	Relation         []string `json:"relation"` // 是否查询对应关系，参数有classify、tag、type、company
	IsUse            string   `json:"is_use"`
	PageNum          int      `json:"pageNum"`
	PageSize         int      `json:"pageSize"`
}

type queryTopicReturn struct {
	Code int             `json:"code"`
	Data TopicReturnData `json:"data"`
}

type TopicReturnData struct {
	Data           []topicModel.Topic `json:"data"`
	PagingArgument PageArgument       `json:"pageArgument"`
}

type queryTopicReturnRelation struct {
	Code int                     `json:"code"`
	Data TopicReturnDataRelation `json:"data"`
}

type TopicReturnDataRelation struct {
	Data           []TopicData  `json:"data"`
	PagingArgument PageArgument `json:"pageArgument"`
}

type TopicData struct {
	topicModel.Topic
	ClassifyInfo []classifyModel.Classify `json:"classifyInfo"`
	CompanyInfo  []companyModel.Company   `json:"companyInfo"`
	TagInfo      []tagModel.Tag           `json:"tagInfo"`
	TypeInfo     []typeModel.Type         `json:"typeInfo"`
}

type queryTopicFromClassifyParams struct {
	ClassifyId []string `json:"classify_id"`
	CreateAt   []string `json:"create_at"`
	DeleteAt   []string `json:"delete_at"`
	UpdateAt   []string `json:"update_at"`
	Relation   []string `json:"relation"` // 是否查询对应关系，参数有classify、tag、type、company
	IsUse      string   `json:"is_use"`
}

type queryTopicFromCompanyParams struct {
	CompanyId []string `json:"company_id"`
	CreateAt  []string `json:"create_at"`
	DeleteAt  []string `json:"delete_at"`
	UpdateAt  []string `json:"update_at"`
	Relation  []string `json:"relation"` // 是否查询对应关系，参数有classify、tag、type、company
	IsUse     string   `json:"is_use"`
}

type queryTopicFromTagParams struct {
	TagId    []string   `json:"tag_id"` // 标签id
	CreateAt []string `json:"create_at"`
	DeleteAt []string `json:"delete_at"`
	UpdateAt []string `json:"update_at"`
	Relation []string `json:"relation"` // 是否查询对应关系，参数有classify、tag、type、company
	IsUse    string   `json:"is_use"`
}

type queryTopicFromTypeParams struct {
	TypeId   []string   `json:"type_id"` // 分类id
	CreateAt []string `json:"create_at"`
	DeleteAt []string `json:"delete_at"`
	UpdateAt []string `json:"update_at"`
	Relation []string `json:"relation"` // 是否查询对应关系，参数有classify、tag、type、company
	IsUse    string   `json:"is_use"`
}

type QueryTopicIdReturn struct {
	Code        int      `json:"code"`
	TopicIdList []string `json:"topic_id_list"` // 题目id数组
}

type queryTopicNoPadingReturn struct {
	Code int                `json:"code"`
	Data []topicModel.Topic `json:"data"`
}

type queryTopicNoPadingRelationReturn struct {
	Code int         `json:"code"`
	Data []TopicData `json:"data"`
}
