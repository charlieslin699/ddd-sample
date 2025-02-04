package httpserver

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type RouteFunc func(s *httpServer, r *gin.RouterGroup)

// API路徑分群
func Group(path string, fns ...RouteFunc) RouteFunc {
	return func(s *httpServer, r *gin.RouterGroup) {
		group := r.Group(path)
		for _, fn := range fns {
			fn(s, group)
		}
	}
}

// 註冊middleware
func Use(fns ...HandlerFunc) RouteFunc {
	return func(s *httpServer, r *gin.RouterGroup) {
		for _, fn := range fns {
			r.Use(handler(s, fn))
		}
	}
}

// 註冊GET API
func GET(path string, fns ...HandlerFunc) RouteFunc {
	return func(s *httpServer, r *gin.RouterGroup) {
		tempFns := make([]gin.HandlerFunc, len(fns))
		for i, fn := range fns {
			tempFns[i] = handler(s, fn)
		}

		r.GET(path, tempFns...)
	}
}

// 註冊POST API
func POST(path string, fns ...HandlerFunc) RouteFunc {
	return func(s *httpServer, r *gin.RouterGroup) {
		tempFns := make([]gin.HandlerFunc, len(fns))
		for i, fn := range fns {
			tempFns[i] = handler(s, fn)
		}

		r.POST(path, tempFns...)
	}
}

// 註冊PUT API
func PUT(path string, fns ...HandlerFunc) RouteFunc {
	return func(s *httpServer, r *gin.RouterGroup) {
		tempFns := make([]gin.HandlerFunc, len(fns))
		for i, fn := range fns {
			tempFns[i] = handler(s, fn)
		}

		r.PUT(path, tempFns...)
	}
}

// 註冊PATCH API
func PATCH(path string, fns ...HandlerFunc) RouteFunc {
	return func(s *httpServer, r *gin.RouterGroup) {
		tempFns := make([]gin.HandlerFunc, len(fns))
		for i, fn := range fns {
			tempFns[i] = handler(s, fn)
		}

		r.PATCH(path, tempFns...)
	}
}

// 註冊DELETE API
func DELETE(path string, fns ...HandlerFunc) RouteFunc {
	return func(s *httpServer, r *gin.RouterGroup) {
		tempFns := make([]gin.HandlerFunc, len(fns))
		for i, fn := range fns {
			tempFns[i] = handler(s, fn)
		}

		r.DELETE(path, tempFns...)
	}
}

// pipeline流程處理
func handler(s *httpServer, fn HandlerFunc) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// panic處理
		defer func() {
			if p := recover(); p != nil {
				// 取得stack trace
				traceback := string(debug.Stack())

				s.onPanic(ctx, fmt.Errorf("%v", p), traceback)
				ctx.Abort()
			}
		}()

		data, err := fn(ctx)

		// 是否跳過
		if isNext(err) {
			return
		}

		// error處理
		if err != nil {
			s.onError(ctx, err)
			ctx.Abort()
			return
		}

		// 回傳結果
		s.onResult(ctx, data)
		ctx.Abort()
	}
}

// 繼續 handler
func Next() (RestfulResult, error) {
	return nil, errorNext
}

// 是否繼續 handler
func isNext(err error) bool {
	return errors.Is(errorNext, err)
}
