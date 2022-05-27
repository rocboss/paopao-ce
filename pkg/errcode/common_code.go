package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000, "服务内部错误")
	InvalidParams             = NewError(10001, "入参错误")
	NotFound                  = NewError(10002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10003, "账户不存在")
	UnauthorizedAuthFailed    = NewError(10004, "账户密码错误")
	UnauthorizedTokenError    = NewError(10005, "鉴权失败，Token 错误或丢失")
	UnauthorizedTokenTimeout  = NewError(10006, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(10007, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(10008, "请求过多")

	GatewayMethodsLimit    = NewError(10109, "网关仅接受GET或POST请求")
	GatewayLostSign        = NewError(10110, "网关请求缺少签名")
	GatewayLostAppKey      = NewError(10111, "网关请求缺少APP KEY")
	GatewayAppKeyInvalid   = NewError(10112, "网关请求无效APP KEY")
	GatewayAppKeyClosed    = NewError(10113, "网关请求APP KEY已停用")
	GatewayParamSignError  = NewError(10114, "网关请求参数签名错误")
	GatewayTooManyRequests = NewError(10115, "网关请求频次超限")

	FileUploadFailed = NewError(10200, "文件上传失败")
	FileInvalidExt   = NewError(10201, "文件类型不合法")
	FileInvalidSize  = NewError(10202, "文件大小超限")
)
