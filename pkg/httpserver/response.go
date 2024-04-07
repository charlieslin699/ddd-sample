package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// 正常回傳 code 0
	SuccessCode string = "0"
)

type ResponseData struct {
	Code string `json:"code"` // 狀態碼 0為正常
	Data any    `json:"data"` // 回傳結果
}

// 回傳 http 200
func ResponseOK(ctx *gin.Context, result RestfulResult) {
	data := ResponseData{
		Code: SuccessCode,
		Data: result,
	}
	ctx.JSON(http.StatusOK, data)
}

// 回傳 http 400
func ResponseFailure(ctx *gin.Context, result ErrorResult) {
	data := ResponseData{
		Code: result.ErrorCode,
		Data: result,
	}
	ctx.JSON(http.StatusBadRequest, data)
}

// 回傳非預期錯誤(httpserver本身)
func responseUnexpectedError(ctx *gin.Context) {
	ResponseFailure(ctx, unexpectedResult)
}
