package errcode

var (
	Success                   = NewError(0, "success")
	ServerError               = NewError(1_0000_000, "server interal error")
	InvalidParams             = NewError(1_0000_001, "invalid params error")
	NotFound                  = NewError(1_0000_002, "not found error")
	UnauthorizedAuthNotExist  = NewError(1_0000_003, "unauthorize error, no appkey and appSecret")
	UnauthorizedTokenError    = NewError(1_0000_004, "unauthorize error, token error")
	UnauthorizedTokenTimeout  = NewError(1_0000_005, "unauthorize error, token time out error")
	UnauthorizedTokenGenerate = NewError(1_0000_006, "unauthorize error, token generate error")
	TooManyRequests           = NewError(1_0000_007, "too many requests error")
)
