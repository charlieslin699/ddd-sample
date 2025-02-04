package header

import (
	"github.com/gin-gonic/gin"
)

var (
	// 第三方 token
	AccessToken = newHeader("AccessToken", "")
	// 遊戲 token
	GameToken = newHeader("GameToken", "")
)

type Header func(ctx *gin.Context) string

func newHeader(key, defaultValue string) Header {
	return func(ctx *gin.Context) string {
		value := ctx.GetHeader(key)
		if value == "" {
			return defaultValue
		}
		return value
	}
}

func (h Header) Get(ctx *gin.Context) string {
	return h(ctx)
}
