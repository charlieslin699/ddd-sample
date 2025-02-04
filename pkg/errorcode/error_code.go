package errorcode

type ErrorCode interface {
	Error() string
	Code() string
	LangIndex() string
	DebugMessage() string
}

type errorCode struct {
	code         string
	debugMessage string
	langIndex    string
}

func NewErrorCode(code, debugMessage, langIndex string) ErrorCode {
	return errorCode{code, debugMessage, langIndex}
}

func (err errorCode) Error() string {
	return err.code
}

func (err errorCode) Code() string {
	return err.code
}

func (err errorCode) LangIndex() string {
	return err.langIndex
}

func (err errorCode) DebugMessage() string {
	return err.debugMessage
}
