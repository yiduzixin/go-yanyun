package libs

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIException struct {
	Code      int    `json:"-"`
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
	Request   string `json:"request"`
}

func (e *APIException) Error() string {
	return e.Msg
}
func newAPIException(code int, errorCode int, msg string) *APIException {
	return &APIException{
		Code:      code,
		ErrorCode: errorCode,
		Msg:       msg,
	}
}

const (
	SERVER_ERROR    = 1000 // 系统错误
	NOT_FOUND       = 1001 // 401错误
	UNKNOWN_ERROR   = 1002 // 未知错误
	PARAMETER_ERROR = 1003 // 参数错误
	AUTH_ERROR      = 1004 // 错误
)

func ServerError() *APIException {
	return newAPIException(http.StatusInternalServerError, SERVER_ERROR, http.StatusText(http.StatusInternalServerError))
}

func NotFound() *APIException {
	return newAPIException(http.StatusNotFound, NOT_FOUND, http.StatusText(http.StatusNotFound))
}

func AccessForbidden() *APIException {
	return newAPIException(http.StatusForbidden, AUTH_ERROR, "验证不通过")
}

func UnknownError(message string) *APIException {
	return newAPIException(http.StatusForbidden, UNKNOWN_ERROR, message)
}

func ParameterError(message string) *APIException {
	return newAPIException(http.StatusBadRequest, PARAMETER_ERROR, message)
}

func HandleNotFound(c *gin.Context) {
	handleErr := NotFound()
	handleErr.Request = c.Request.Method + " " + c.Request.URL.String()
	c.JSON(handleErr.Code, handleErr)
	return
}
