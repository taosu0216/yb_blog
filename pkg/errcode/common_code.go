package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000, "服务内部错误")
	InvalidParams             = NewError(10001, "入参错误")
	NotFound                  = NewError(10002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10003, "鉴权失败,找不到对应的APIKey")
	UnauthorizedTokenError    = NewError(10004, "token错误")
	UnauthorizedTokenTimeout  = NewError(10005, "token已超时")
	UnauthorizedTokenGenerate = NewError(10006, "token生成失败")
	TooManyRequests           = NewError(10007, "请求过多")
)
