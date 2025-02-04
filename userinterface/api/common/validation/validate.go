package validation

import "github.com/gin-gonic/gin"

func Validate[T any](ctx *gin.Context) (result T, err error) {
	err = ctx.ShouldBind(&result)
	return
}

func ValidateURI[T any](ctx *gin.Context) (result T, err error) {
	err = ctx.ShouldBindUri(&result)
	return
}
