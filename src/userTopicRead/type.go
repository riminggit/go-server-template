package userTopicRead

type CreateParams struct {
	TopicId int `json:"topic_id"` // 对应题目id
}

type QuerSimpleParams struct {
	UserId  int `json:"user_id"`  // 用户id
	TopicId int `json:"topic_id"` // 对应题目id
}

type UpdateSimpleParams struct {
	UserId  int `json:"user_id"`  // 用户id
	TopicId int `json:"topic_id"` // 对应题目id
	ReadNum int `json:"read_num"` // 对应题目id
}

type DeleteParams struct {
	UserId  int `json:"user_id"`  // 用户id
	TopicId int `json:"topic_id"` // 对应题目id
}

type CreateReturn struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
