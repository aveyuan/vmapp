package vconst

type SendMedia int8

const (
	EmailSendMedia SendMedia = 1
	PhoneSendMedia SendMedia = 2
)

type SendType int8

const (
	LoginSendType SendType = iota + 1
	RegisterSendType
	ForgetSendType
	Message
)

var SendMediaMap = map[SendMedia]string{
	EmailSendMedia: "邮箱",
	PhoneSendMedia: "手机",
}

var SendTypeMap = map[SendType]string{
	LoginSendType:    "登录",
	RegisterSendType: "注册",
	ForgetSendType:   "找回",
	Message:          "留言",
}
