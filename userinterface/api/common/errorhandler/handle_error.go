package errorhandler

import (
	"ddd-sample/application/query/lang"
	"ddd-sample/pkg/errorcode"
	"ddd-sample/pkg/httpserver"
	"ddd-sample/pkg/log"
	"ddd-sample/userinterface/api/common/cookie"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleError(getLangQuery lang.GetLangQuery) httpserver.ErrorHandlerFunc {
	return func(ctx *gin.Context, err error) (result httpserver.ErrorResult, err1 error) {
		locale := cookie.Locale.Get(ctx)

		switch v := err.(type) {
		case errorcode.ErrorCode:
			handleErrorCode(v)
			result.ErrorCode = v.Code()
			result.Message = getLang(ctx, getLangQuery, v.LangIndex(), locale)
		case validator.ValidationErrors: // gin參數驗證
			result.ErrorCode = errorcode.ErrValidation.Code()
			result.Message = v.Error()
		case validator.FieldError: // gin參數驗證
			result.ErrorCode = errorcode.ErrValidation.Code()
			result.Message = v.Error()
		default:
			result.ErrorCode = errorcode.ErrUnexpected.Code()
			result.Message = getLang(ctx, getLangQuery, errorcode.ErrUnexpected.LangIndex(), locale)
		}

		return
	}
}

func handleErrorCode(errcode errorcode.ErrorCode) {
	switch errcode {
	// 取不到context資料，可能是middleware沒有設定
	case errorcode.ErrContextGetFailed:
		traceback := string(debug.Stack())
		log.Error(errcode, traceback)
	}
}

func getLang(ctx *gin.Context, getLangQuery lang.GetLangQuery, key, locale string) string {
	value, _ := getLangQuery.Execute(ctx, lang.GetLangQueryInput{
		Key:    key,
		Locale: locale,
	})

	return value
}
