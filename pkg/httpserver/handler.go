package httpserver

import (
	"github.com/gin-gonic/gin"
)

type RestfulResult any

type ErrorResult struct {
	ErrorCode string `json:"-"`
	Message   string `json:"message"`
}

var (
	// 非預期錯誤(httpserver本身)
	unexpectedResult ErrorResult = ErrorResult{ErrorCode: "-1", Message: ""}
)

type HandlerFunc func(*gin.Context) (RestfulResult, error)

type ErrorHandlerFunc func(*gin.Context, error) (ErrorResult, error)

type PanicHandlerFunc func(ctx *gin.Context, err error, traceback string) (ErrorResult, error)
