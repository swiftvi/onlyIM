package xcode

var codeMsg = map[int]string{
	SERVER_COMMON_ERROR: "Server common error",
	REQUEST_PARAM_ERROR: "Request param error",
	DATABASE_ERROR:      "Database error",
}

func CodeToMsg(code int) string {
	if msg, ok := codeMsg[code]; ok {
		return msg
	}

	return codeMsg[SERVER_COMMON_ERROR]
}