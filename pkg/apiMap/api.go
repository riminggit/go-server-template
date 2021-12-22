package apiMap

const (
	API_PREFIX = "api"

	CLASSIFY_PREFIX          = "/classify"
	GET_QUERY_CLASSIFY       = "/query-classify"
	GET_QUERY_CLASSIFY_TYPE  = "/query-classify-type"
	POST_QUERY_CLASSIFY      = "/query-classify-multiple"
	POST_QUERY_CLASSIFY_TYPE = "/query-classify-type-multiple"

	COMPANY_PREFIX     = "/company"
	GET_QUERY_COMPANY  = "/query-company"
	POST_QUERY_COMPANY = "/query-company-multiple"

	FEEDBACK_PREFIX = "/feedback"

	TAG_PREFIX = "/tag"
	GET_QUERY_TAG  = "/query-tag"
	POST_QUERY_TAG = "/query-tag-multiple"

	TOPIC_PREFIX = "/topic"
	POST_QUERY_TOPIC_SET = "/query-topic-set"
	POST_QUERY_TOPIC = "/query-topic"
	POST_QUERY_TOPIC_RELACTION = "/query-topic-relaction"
	POST_QUERY_TOPIC_CLASSIFY_MID = "/query-topic/classify/mid"
	POST_QUERY_TOPIC_COMPANY_MID = "/query-topic/company/mid"
	POST_QUERY_TOPIC_TAG_MID = "/query-topic/tag/mid"
	POST_QUERY_TOPIC_TYPE_MID = "/query-topic/type/mid"
	

	TYPE_PREFIX = "/type"
	GET_QUERY_TYPE  = "/query-type"
	POST_QUERY_TYPE = "/query-type-multiple"

	USER_COLLECT_PREFIX = "/userCollect"

	USER_PREFIX = "/user"
	USER_ACCOUNT_PREFIX = "/user/account-manage"
	USER_LOGIN_PREFIX = "/user-login"
	POST_USER_CHANGE_PHONE = "/change-phone"
	POST_USER_LOGIN = "/login"
	POST_USER_QUERY_USER_INFO = "/query-user-info"
	POST_USER_WX_LOGIN = "/wxLogin"
	POST_USER_WX_GET_OPEN_ID = "/wxapp-get-openid"
	POST_USER_WX_DNCRYPT = "/wxapp-dncrypt"

	USER_TOPIC_PREFIX = "/user-topic"
	POST_USER_TOPIC_QUERY = "/query-user-topic"
	POST_USER_ANSWER_TOPIC_QUERY = "/query-user-answer-topic"
)
