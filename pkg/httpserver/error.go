package httpserver

import (
	"errors"
)

var (
	// next api handler
	errorNext = errors.New("")
)

type CatchFunc func(*httpServer)

// 註冊error handler，將回傳最後一個ErrorResult
func CatchError(fns ...ErrorHandlerFunc) CatchFunc {
	return func(s *httpServer) {
		s.errorHandlers = append(s.errorHandlers, fns...)
	}
}

// 註冊panic handler，將回傳最後一個ErrorResult
func CatchPanic(fns ...PanicHandlerFunc) CatchFunc {
	return func(s *httpServer) {
		s.panicHandlers = append(s.panicHandlers, fns...)
	}
}
