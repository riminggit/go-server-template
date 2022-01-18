package e

var MsgFlags = map[int]string{
	SUCCESS:                 "成功",
	ERROR:                   "失败",
	INVALID_PARAMS:          "请求参数错误",
	UN_AUTHORIZED:           "没有权限",
	UN_AUTHORIZED_NOT_ADMIN: "没有管理员权限",

	IS_LOGIN:                         "已经登陆",
	USER_NOT_LOGIN:                   "用户未登录",
	USER_LAYOUT_SUCCESS:              "用户登出成功",
	ERROR_AUTH_CHECK_TOKEN_FAIL:      "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:   "Token已超时",
	ERROR_AUTH_TOKEN:                 "Token生成失败",
	ERROR_AUTH:                       "Token错误",
	ACCOUNT_DOES_NOT_EXIST:           "账号信息不存在",
	EMAIL_DOES_NOT_EXIST:             "邮箱不存在",
	NICKNAME_DOES_NOT_EXIST:          "用户名不存在",
	PHONE_DOES_NOT_EXIST:             "手机号不存在",
	PASSWORD_MISTAKE:                 "密码错误",
	NOT_EDUIT_AUTH:                   "用户无编辑权限",
	NOT_DELETE_AUTH:                  "用户无删除权限",
	USER_NEED_TEST_TO_ADD_EXPERIENCE: "用户需要考试才能继续增加经验",

	NO_DATA_EXISTS:   "不存在相关数据",
	CREATE_DATA_FALE: "新建数据失败",
	DELETE_DATA_FALE: "删除数据失败",
	UPDATE_DATA_FALE: "更新数据失败",
	RECORD_FALE:      "记录失败",

	WX_DNCRYPT_FAIL: "微信小程序解密失败",

	FEEDBACK_FALE: "反馈失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
