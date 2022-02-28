package apiMap

const (
	API_PREFIX = "api"

	// 一级分类
	CLASSIFY_PREFIX          = "/classify"
	GET_QUERY_CLASSIFY       = "/query-classify"
	GET_QUERY_CLASSIFY_TYPE  = "/query-classify-type"
	POST_QUERY_CLASSIFY      = "/query-classify-multiple"
	POST_QUERY_CLASSIFY_TYPE = "/query-classify-type-multiple"
	POST_CREATE_CLASSIFY     = "/create-classify-multiple"
	POST_UPDATE_CLASSIFY     = "/update-classify-multiple"
	POST_DELETE_CLASSIFY     = "/delete-classify-multiple"

	// 二级分类
	TYPE_PREFIX      = "/type"
	GET_QUERY_TYPE   = "/query-type"
	POST_QUERY_TYPE  = "/query-type-multiple"
	POST_CREATE_TYPE = "/create-type-multiple"
	POST_UPDATE_TYPE = "/update-type-multiple"
	POST_DELETE_TYPE = "/delete-type-multiple"

	// 公司信息
	COMPANY_PREFIX      = "/company"
	GET_QUERY_COMPANY   = "/query-company"
	POST_QUERY_COMPANY  = "/query-company-multiple"
	POST_CREATE_COMPANY = "/create-company-multiple"
	POST_UPDATE_COMPANY = "/update-company-multiple"
	POST_DELETE_COMPANY = "/delete-company-multiple"

	// 反馈
	FEEDBACK_PREFIX           = "/feedback"
	POST_QUERY_FEEDBACK       = "/query-feedback"
	POST_QUERY_FEEDBACK_USER  = "/query-feedback-user"
	POST_CREATE_FEEDBACK_USER = "/create-feedback-user"
	POST_DELETE_FEEDBACK_USER = "/delete-feedback-user"
	POST_DELETE_FEEDBACK      = "/delete-feedback"
	POST_UPDATE_FEEDBACK      = "/update-feedback"
	POST_UPDATE_FEEDBACK_USER = "/update-feedback-user"

	// 标签
	TAG_PREFIX      = "/tag"
	GET_QUERY_TAG   = "/query-tag"
	POST_QUERY_TAG  = "/query-tag-multiple"
	POST_CREATE_TAG = "/create-tag-multiple"
	POST_UPDATE_TAG = "/update-tag-multiple"
	POST_DELETE_TAG = "/delete-tag-multiple"

	// 题目
	TOPIC_PREFIX                                = "/topic"
	POST_QUERY_TOPIC_FROM_CLASSIFY              = "/query-topic-from-classify"
	POST_QUERY_TOPIC_FROM_CLASSIFY_RELATION     = "/query-topic-from-classify-relation"
	POST_QUERY_TOPIC_FROM_CLASSIFY_GET_TOPIC_ID = "/query-topic-from-classify-get-topic-id"
	POST_QUERY_TOPIC_FROM_COMPANY               = "/query-topic-from-company"
	POST_QUERY_TOPIC_FROM_COMPANY_RELATION      = "/query-topic-from-company-relation"
	POST_QUERY_TOPIC_FROM_COMPANY_GET_TOPIC_ID  = "/query-topic-from-company-get-topic-id"
	POST_QUERY_TOPIC_FROM_TAG                   = "/query-topic-from-tag"
	POST_QUERY_TOPIC_FROM_TAG_RELATION          = "/query-topic-from-tag-relation"
	POST_QUERY_TOPIC_FROM_TAG_GET_TOPIC_ID      = "/query-topic-from-tag-get-topic-id"
	POST_QUERY_TOPIC_FROM_TYPE                  = "/query-topic-from-type"
	POST_QUERY_TOPIC_FROM_TYPE_RELATION         = "/query-topic-from-type-relation"
	POST_QUERY_TOPIC_FROM_TYPE_GET_TOPIC_ID     = "/query-topic-from-type-get-topic-id"
	POST_QUERY_TOPIC_RELACTION                  = "/query-topic-relaction"
	POST_QUERY_TOPIC_CLASSIFY_MID               = "/query-topic/classify/mid"
	POST_QUERY_TOPIC_COMPANY_MID                = "/query-topic/company/mid"
	POST_QUERY_TOPIC_TAG_MID                    = "/query-topic/tag/mid"
	POST_QUERY_TOPIC_TYPE_MID                   = "/query-topic/type/mid"
	POST_QUERY_TOPIC_CLASSIFY_MID_PADING        = "/query-topic/classify/mid/pading"
	POST_QUERY_TOPIC_COMPANY_MID_PADING         = "/query-topic/company/mid/pading"
	POST_QUERY_TOPIC_TAG_MID_PADING             = "/query-topic/tag/mid/pading"
	POST_QUERY_TOPIC_TYPE_MID_PADING            = "/query-topic/type/mid/pading"
	POST_QUERY_TOPIC                            = "/query-topic"
	POST_QUERY_TOPIC_RANDOM                     = "/query-topic-random"
	POST_QUERY_TOPIC_SIMPLE                     = "/query-topic-simple"
	POST_CREATE_TOPIC                           = "/create-topic"
	POST_DELETE_TOPIC                           = "/delete-topic"
	POST_UPDATE_TOPIC                           = "/update-topic"
	POST_QUERY_TOPIC_SET                        = "/query-topic-set"
	POST_QUERY_TOPIC_SET_SIMPLE                 = "/query-topic-set-simple"
	POST_QUERY_TOPIC_SET_SIMPLE_SINGLE          = "/query-topic-set-simple-single"
	POST_CREATE_TOPIC_SET                       = "/create-topic-set"
	POST_DELETE_TOPIC_SET                       = "/delete-topic-set"
	POST_UPDATE_TOPIC_SET                       = "/update-topic-set"
	POST_QUERY_TOPIC_SET_RANDOM                 = "/query-topic-set-random"
	TOPIC_ALL_COUNT                             = "/topic-all-count"
	TOPIC_ID_MIN                                = "/topic-id-min"
	TOPIC_SET_ALL_COUNT                         = "/topic-set-all-count"
	TOPIC_SET_ID_MIN                            = "/topic-set-id-min"
	GET_A_DAILY_TOPIC                           = "/a-daily-topic"

	// 用户控制
	USER_COLLECT_PREFIX = "/userCollect"

	// 用户相关
	USER_PREFIX                          = "/user"
	USER_ACCOUNT_PREFIX                  = "/user/account-manage"
	USER_LOGIN_PREFIX                    = "/user-login"
	POST_USER_CHANGE_PHONE               = "/change-phone"
	POST_USER_LOGIN                      = "/login"
	POST_ADMIN_QUERY_USER_INFO           = "/query-user-info"
	POST_QUERY_USER_EXPERIENCE           = "/query-user-experience"           // 管理员查询
	POST_QUERY_USER_EXPERIENCE_FROM_USER = "/query-user-experience-from-user" // 用户查询

	// 微信相关
	POST_USER_WX_LOGIN       = "/wxLogin"
	POST_USER_WX_GET_OPEN_ID = "/wxapp-get-openid"
	POST_USER_WX_DNCRYPT     = "/wxapp-dncrypt"

	// 用户题目
	USER_TOPIC_PREFIX             = "/user-topic"
	POST_USER_ADD_TOPIC_QUERY     = "/query-user-topic"
	POST_USER_ANSWER_TOPIC_QUERY  = "/query-user-answer-topic"
	POST_USER_ANSWER_TOPIC_CREATE = "/create-user-answer-topic"
	POST_USER_ANSWER_TOPIC_UPDATE = "/update-user-answer-topic"

	// 查询题目数量
	QUERY_TOPIC_COUNT_BY_DEGREE = "/query-topic-count-by-degree"
	QUERY_TOPIC_COUNT_BY_LEVEL  = "/query-topic-count-by-level"
)
