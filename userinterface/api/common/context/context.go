package context

import (
	"ddd-sample/pkg/errorcode"
	"reflect"

	"github.com/gin-gonic/gin"
)

var (
	Username     ContextKey[string] = "username"
	UserUID      ContextKey[string] = "userUID"
	BusinessCode ContextKey[string] = "businessCode"
)

type ContextKey[T any] string

func Context[T any](fns ...ContextKeyOptionfunc) ContextKey[T] {
	var c T
	t := reflect.TypeOf(c)
	key := t.PkgPath() + "." + t.Name()

	for _, fn := range fns {
		fn(&key)
	}

	ck := ContextKey[T](key)

	return ck
}

func (ck ContextKey[T]) Get(ctx *gin.Context) (data T, err error) {
	key := string(ck)
	if value, isExist := ctx.Get(key); isExist {
		return value.(T), errorcode.ErrContextGetFailed
	}

	return
}

func (ck ContextKey[T]) Set(ctx *gin.Context, data T) {
	key := string(ck)
	ctx.Set(key, data)
}

type ContextKeyOptionfunc func(key *string)

func WithKey(k string) ContextKeyOptionfunc {
	return func(key *string) {
		*key = k
	}
}
