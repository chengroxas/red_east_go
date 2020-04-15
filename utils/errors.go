package common

var (
	CODE_BAD_PARAM       = 400
	CODE_REQUIRE_AUTH    = 401
	CODE_PERSIMMON_ERROR = 403
	CODE_NOT_EXIST       = 404
	CODE_METHOD_NOT_ALL  = 405
	CODE_WRONG_ITEM      = 406
)

var errorCodeMsg = map[int]string{
	CODE_BAD_PARAM:    "请校验参数",
	CODE_REQUIRE_AUTH: "需要验证",
}

func GetErrorCodeMsg(code int) string {
	if value, ok := errorCodeMsg[code]; ok {
		return value
	} else {
		return ""
	}
}
