package errors

const (
	NotFound         ErrorType = 404 //资源不存在
	ParamInvalid     ErrorType = 400 //参数错误
	UnAuthorize      ErrorType = 401 //没有用户信息
	PermissionDenied ErrorType = 403 //没有权限
	SystemError      ErrorType = 500 //系统业务错误
)
