package validation

import "ddd-sample/pkg/httpserver"

func Validate[T any](ctx *httpserver.Context) (result T, err error) {
	err = ctx.ShouldBind(&result)
	return
}

func ValidateURI[T any](ctx *httpserver.Context) (result T, err error) {
	err = ctx.ShouldBindUri(&result)
	return
}
