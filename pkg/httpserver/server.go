package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPServer interface {
	Route(fn ...RouteFunc)
	Catch(fn ...CatchFunc)
	Engine() http.Handler
}

type httpServer struct {
	engine        *gin.Engine
	errorHandlers []ErrorHandlerFunc
	panicHandlers []PanicHandlerFunc
}

// 工廠
func NewHTTPServer() HTTPServer {
	return &httpServer{
		engine:        gin.New(),
		errorHandlers: make([]ErrorHandlerFunc, 0),
		panicHandlers: make([]PanicHandlerFunc, 0),
	}
}

// 回傳API Engine
func (s *httpServer) Engine() http.Handler {
	return s.engine
}

// 註冊路由
func (s *httpServer) Route(fns ...RouteFunc) {
	for _, fn := range fns {
		fn(s, &s.engine.RouterGroup)
	}
}

// 註冊錯誤處理
func (s *httpServer) Catch(fns ...CatchFunc) {
	for _, fn := range fns {
		fn(s)
	}
}

// 正常回應，http response 200
func (s *httpServer) onResult(ctx *Context, result RestfulResult) {
	ResponseOK(ctx.Context, result)
}

// 錯誤處理，http response 400
func (s *httpServer) onError(ctx *Context, err error) {
	var errResult ErrorResult
	var fnErr error

	for _, fn := range s.errorHandlers {
		// 帶入新的context
		ctx = buildContext(ctx.Context)

		errResult, fnErr = fn(ctx, err)
		if fnErr != nil {
			responseUnexpectedError(ctx.Context) // error handler發生error // TODO: log紀錄錯誤
			return
		}

		if ctx.isNexted() {
			continue
		}

		break
	}

	ResponseFailure(ctx.Context, errResult)
}

// 非預期錯誤處理，http response 400
func (s *httpServer) onPanic(ctx *Context, err error, trace string) {
	var errResult ErrorResult
	var fnErr error

	for _, fn := range s.panicHandlers {
		// 帶入新的context
		ctx = buildContext(ctx.Context)

		errResult, fnErr = fn(ctx, err, trace)
		if fnErr != nil {
			responseUnexpectedError(ctx.Context) // panic handler發生error // TODO: log紀錄錯誤
			return
		}

		if ctx.isNexted() {
			continue
		}

		break
	}

	ResponseFailure(ctx.Context, errResult)
}
