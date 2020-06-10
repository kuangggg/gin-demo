package e

var MsgFlags = map[int]string {
	SUCCESS : "OK",
	ERROR : "SERVER ERROR",
	INVALID_PARAMS: "无效参数",
}

func GetMsg(code int) string  {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
