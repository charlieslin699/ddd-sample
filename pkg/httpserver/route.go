package httpserver

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type RouteFunc func(*httpServer, *gin.RouterGroup)

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
			r.Use(do(s, fn))
		}
	}
}

// 註冊GET API
func GET(path string, fns ...HandlerFunc) RouteFunc {
	return func(s *httpServer, r *gin.RouterGroup) {
		tempFns := make([]gin.HandlerFunc, len(fns))
		for i, fn := range fns {
			tempFns[i] = do(s, fn)
		}

		r.GET(path, tempFns...)
	}
}

// 註冊POST API
func POST(path string, fns ...HandlerFunc) RouteFunc {
	return func(s *httpServer, r *gin.RouterGroup) {
		tempFns := make([]gin.HandlerFunc, len(fns))
		for i, fn := range fns {
			tempFns[i] = do(s, fn)
		}

		r.POST(path, tempFns...)
	}
}

// 註冊PUT API
func PUT(path string, fns ...HandlerFunc) RouteFunc {
	return func(s *httpServer, r *gin.RouterGroup) {
		tempFns := make([]gin.HandlerFunc, len(fns))
		for i, fn := range fns {
			tempFns[i] = do(s, fn)
		}

		r.PUT(path, tempFns...)
	}
}

// 註冊PATCH API
func PATCH(path string, fns ...HandlerFunc) RouteFunc {
	return func(s *httpServer, r *gin.RouterGroup) {
		tempFns := make([]gin.HandlerFunc, len(fns))
		for i, fn := range fns {
			tempFns[i] = do(s, fn)
		}

		r.PATCH(path, tempFns...)
	}
}

// 註冊DELETE API
func DELETE(path string, fns ...HandlerFunc) RouteFunc {
	return func(s *httpServer, r *gin.RouterGroup) {
		tempFns := make([]gin.HandlerFunc, len(fns))
		for i, fn := range fns {
			tempFns[i] = do(s, fn)
		}

		r.DELETE(path, tempFns...)
	}
}

// pipeline流程處理
func do(s *httpServer, fn HandlerFunc) func(*gin.Context) {
	return func(ginCtx *gin.Context) {
		ctx := buildContext(ginCtx)

		// panic處理
		defer func() {
			if p := recover(); p != nil {
				// 取得stack trace
				traceback := string(debug.Stack())

				s.onPanic(ctx, fmt.Errorf("%v", p), traceback)
			}
		}()

		data, err := fn(buildContext(ginCtx))
		if err != nil {
			s.onError(ctx, err)
			return
		}

		// 是否繼續pipeline
		if ctx.isNexted() {
			return
		}

		// 回傳結果
		s.onResult(ctx, data)
	}
}
