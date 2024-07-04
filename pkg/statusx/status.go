package statusx

type Status int

// GetMsg ...
func GetMsg(status Status) string {
	return StatusMap[status]
}

const (
	StatusOk Status = iota
	StatusUnknownError
	StatusApiNotFount
	StatusUnauthorized
	StatusVerifyTokenError
	StatusPermissionsError
	StatusInvalidRequest
	StatusRequestTimeout
	StatusTooManyRequests
	StatusTemporarilyUnavailable
	StatusInternalServerError
	StatusExternalServerError
	StatusPasswordError
)

// StatusMap ...
var StatusMap = map[Status]string{
	StatusOk:                     "",
	StatusUnknownError:           "未知错误",
	StatusApiNotFount:            "请求的API不存在",
	StatusUnauthorized:           "TOKEN错误",
	StatusVerifyTokenError:       "TOKEN校验错误",
	StatusPermissionsError:       "权限校验错误",
	StatusInvalidRequest:         "请求参数错误",
	StatusRequestTimeout:         "请求超时",
	StatusTooManyRequests:        "请求过于频繁",
	StatusTemporarilyUnavailable: "服务暂时无法访问",
	StatusInternalServerError:    "服务内部异常",
	StatusExternalServerError:    "外部服务异常",
	StatusPasswordError:          "密码错误",
}
