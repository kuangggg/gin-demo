package e

var MsgFlags = map[int]string {
	SUCCESS : "OK",
	ERROR : "SERVER ERROR",
	INVALID_PARAMS: "无效参数",

	ERROR_EXIST_TAG: "tag 已经存在",
}

func GetMsg(code int) string  {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}

	return MsgFlags[ERROR]
}
