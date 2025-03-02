package panichandler

import (
	"ddd-sample/application/query/lang"
	"ddd-sample/pkg/errorcode"
	"ddd-sample/pkg/httpserver"
	"ddd-sample/pkg/log"
	"ddd-sample/userinterface/api/common/cookie"
)

func HandlePanic(getLangQuery lang.GetLangQuery) httpserver.PanicHandlerFunc {
	return func(ctx *httpserver.Context, errIn error, traceback string) (result httpserver.ErrorResult, errOut error) {
		locale := cookie.Locale.Get(ctx)

		result.ErrorCode = errorcode.ErrUnexpected.Code()
		result.Message = getLang(ctx, getLangQuery, errorcode.ErrUnexpected.LangIndex(), locale)
		log.Debug(errIn.Error(), traceback)
		return
	}
}

func getLang(ctx *httpserver.Context, getLangQuery lang.GetLangQuery, key, locale string) string {
	value, _ := getLangQuery.Execute(ctx, lang.GetLangQueryInput{
		Key:    key,
		Locale: locale,
	})

	return value
}
