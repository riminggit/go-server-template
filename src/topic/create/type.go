package topicCreate

type CreateParams struct {
	Data []TopicParams `json:"data"`
}

type CreateReturn struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
}

type TopicParams struct {
	Title            string   `json:"title"`               // 题目标题
	QuestionType     int      `json:"question_type"`       // 题目类型，1为选择题，2为解析题，3既是选择又是解析
	Analysis         string   `json:"analysis"`            // 题解，内容为富文本
	SelectAnalysis   string   `json:"select_analysis"`     // 选择题答案，对象字符串
	Degree           int      `json:"degree" `             // 难度，简单：0，中等：1，难：2，极难：3
	Level            int      `json:"level"`               // 1 初级 ， 2 中级 ， 3 高级 ， 4 资深 ，5 专家 ， 6 资深专家 ， 7 研究员
	IsBaseTopic      int      `json:"is_base_topic" `      // 0 否  1是 ，是否是基础题
	IsImportantTopic int      `json:"is_important_topic" ` // 0 否  1是 ，是否是重点题
	ComeFrom         string   `json:"come_from" `          // 来源
	ClassifyId       []string `json:"classify_id"`
	CompanyId        []string `json:"company_id"`
	TagId            []string `json:"tag_id"`
	TypeId           []string `json:"type_id"`
}
