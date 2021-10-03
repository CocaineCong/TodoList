package e

var MsgFlags = map[int]string{
	SUCCESS:                 "ok",
	ERROR:                   "fail",
	INVALID_PARAMS:          "请求参数错误",

	ERROR_AUTH_CHECK_TOKEN_FAIL:       "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:    "Token已超时",
	ERROR_AUTH_TOKEN:                  "Token生成失败",
	ERROR_AUTH:                        "Token错误",
	ERROR_AUTH_INSUFFICIENT_AUTHORITY: "权限不足",
	ERROR_NOT_COMPARE:"不匹配",
	ERROR_DATABASE: "数据库操作出错,请重试",

}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
