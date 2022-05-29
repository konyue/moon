package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParma
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeSererBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParma:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeSererBusy:       "服务忙",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeSererBusy]
	}
	return msg
}
