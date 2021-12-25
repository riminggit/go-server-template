package apiMap

var ApiName = map[string]string{
	API_PREFIX:               "api一级前缀",
	CLASSIFY_PREFIX:          "分类接口前缀",
	GET_QUERY_CLASSIFY:       "查询分类_单个参数",
	GET_QUERY_CLASSIFY_TYPE:  "查询分类关联二级分类_单个参数",
	POST_QUERY_CLASSIFY:      "查询分类_多个参数",
	POST_QUERY_CLASSIFY_TYPE: "查询二类关联二级分类级分类_多个参数",

	TYPE_PREFIX:     "二级分类接口前缀",
	GET_QUERY_TYPE:  "查询二级分类_单个参数",
	POST_QUERY_TYPE: "查询二级分类_单个参数",

	COMPANY_PREFIX:     "公司接口前缀",
	GET_QUERY_COMPANY:  "查询公司_单个参数",
	POST_QUERY_COMPANY: "查询公司_多个参数",

	FEEDBACK_PREFIX: "反馈接口前缀",

	TAG_PREFIX:     "标签接口前缀",
	GET_QUERY_TAG:  "查询标签_单个参数",
	POST_QUERY_TAG: "查询标签_多个参数",

	TOPIC_PREFIX:                         "题目接口前缀",
	POST_QUERY_TOPIC_SET:                 "查询套题",
	POST_QUERY_TOPIC:                     "查询题目",
	POST_QUERY_TOPIC_FROM_CLASSIFY:       "通过类型查询题目",
	POST_QUERY_TOPIC_FROM_COMPANY:        "通过公司查询题目",
	POST_QUERY_TOPIC_FROM_TAG:            "通过标签查询题目",
	POST_QUERY_TOPIC_FROM_TYPE:           "通过二级分类查询题目",
	POST_QUERY_TOPIC_RELACTION:           "查询题目及对应关系",
	POST_QUERY_TOPIC_CLASSIFY_MID:        "查询题目关系表_类型",
	POST_QUERY_TOPIC_COMPANY_MID:         "查询题目关系表_公司",
	POST_QUERY_TOPIC_TAG_MID:             "查询题目关系表_标签",
	POST_QUERY_TOPIC_TYPE_MID:            "查询题目关系表_二级分类",
	POST_QUERY_TOPIC_CLASSIFY_MID_PADING: "查询题目关系表_类型_分页",
	POST_QUERY_TOPIC_COMPANY_MID_PADING:  "查询题目关系表_公司_分页",
	POST_QUERY_TOPIC_TAG_MID_PADING:      "查询题目关系表_标签_分页",
	POST_QUERY_TOPIC_TYPE_MID_PADING:     "查询题目关系表_二级分类_分页",

	USER_COLLECT_PREFIX: "用户控制操作接口前缀",

	USER_PREFIX:                "用户接口前缀",
	USER_ACCOUNT_PREFIX:        "用户账号接口前缀",
	USER_LOGIN_PREFIX:          "用户登陆接口前缀",
	POST_USER_CHANGE_PHONE:     "用户修改手机号接口",
	POST_USER_LOGIN:            "用户登陆接口",
	POST_ADMIN_QUERY_USER_INFO: "查询用户信息",
	POST_USER_WX_LOGIN:         "用户微信登陆",
	POST_USER_WX_GET_OPEN_ID:   "获取用户微信openid",
	POST_USER_WX_DNCRYPT:       "用户微信解密",

	USER_TOPIC_PREFIX:            "用户题目接口前缀",
	POST_USER_ADD_TOPIC_QUERY:    "用户题目查询接口",
	POST_USER_ANSWER_TOPIC_QUERY: "用户答题记录查询",
}

func GetApiName(code string) string {
	msg, ok := ApiName[code]
	if ok {
		return msg
	}
	return ""
}

func GetRedisPrefixName(apiMap string) string {
	return "interview" + apiMap + ":"
}
