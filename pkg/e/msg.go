package e

var MsgFlags = map[int]string {
    SUCCESS: "ok",
    ERROR: "服务器错误",
    INVALID_PARAMS: "请求参数错误",
    INVALID_AUTH_PARAMS: "Auth参数错误",

    ERROR_NOT_EXIST_EXAMTESTING: "该考试不存在",
    ERROR_NOT_EXIST_EXAMTESTING_EXAMINEE: "账号或者考试不存在",
    ERROR_PASSWORD_LOGIN: "账号或密码错误",

    ERROR_AUTH_CHECK_TOKEN_FAIL: "Token鉴权失败",
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
    ERROR_AUTH_TOKEN: "Token生成失败",
    ERROR_AUTH: "Token错误",

    ERROR_JSON_HANDLE: "Json转化处理错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}