package ecode

var (
	Ok               = New(100, "成功")
	Err              = New(101, "未知错误,请联系管理员")
	JsonUnmarshalErr = New(102, "解析json数据错误")
	PixivErr         = New(300, "")
	SendMessageErr   = New(500, "")
)
