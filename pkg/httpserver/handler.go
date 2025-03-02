package httpserver

type RestfulResult any

type ErrorResult struct {
	ErrorCode string `json:"-"`
	Message   string `json:"message"`
}

var (
	// 非預期錯誤(httpserver本身)
	unexpectedResult ErrorResult = ErrorResult{ErrorCode: "-1", Message: ""}
)

type HandlerFunc func(*Context) (RestfulResult, error)

type ErrorHandlerFunc func(*Context, error) (ErrorResult, error)

type PanicHandlerFunc func(ctx *Context, err error, traceback string) (ErrorResult, error)
