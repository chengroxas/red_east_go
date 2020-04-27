package common

var (
	CODE_BAD_PARAM        = 400
	CODE_BAD_AUTH         = 401
	CODE_PERSIMMON_ERROR  = 403
	CODE_NOT_EXIST        = 404
	CODE_METHOD_NOT_ALLOW = 405
	CODE_WRONG_ITEM       = 406
)

var errorCodeMsg = map[int]string{
	CODE_BAD_PARAM:        "请校验参数",
	CODE_BAD_AUTH:         "错误的签名校验",
	CODE_WRONG_ITEM:       "错误的数据",
	CODE_PERSIMMON_ERROR:  "没有权限",
	CODE_NOT_EXIST:        "不存在该数据",
	CODE_METHOD_NOT_ALLOW: "不允许的请求方式",
}

func GetErrorCodeMsg(code int) string {
	if value, ok := errorCodeMsg[code]; ok {
		return value
	} else {
		return ""
	}
}
