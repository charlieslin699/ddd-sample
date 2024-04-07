package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer interface {
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
func NewHttpServer() HttpServer {
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
func (s *httpServer) onResult(ctx *gin.Context, result RestfulResult) {
	ResponseOK(ctx, result)
}

// 錯誤處理，http response 400
func (s *httpServer) onError(ctx *gin.Context, err error) {
	var errResult ErrorResult
	var fnErr error

	for _, fn := range s.errorHandlers {
		errResult, fnErr = fn(ctx, err)
		if isNext(fnErr) {
			continue
		}

		// error handler發生error // TODO: log紀錄錯誤
		if fnErr != nil {
			responseUnexpectedError(ctx)
			return
		}

		break
	}

	ResponseFailure(ctx, errResult)
}

// 非預期錯誤處理，http response 400
func (s *httpServer) onPanic(ctx *gin.Context, err error, trace string) {
	var errResult ErrorResult
	var fnErr error

	for _, fn := range s.panicHandlers {
		errResult, fnErr = fn(ctx, err, trace)
		if isNext(fnErr) {
			continue
		}

		// panic handler發生error // TODO: log紀錄錯誤
		if fnErr != nil {
			responseUnexpectedError(ctx)
			return
		}

		break
	}

	ResponseFailure(ctx, errResult)
}
