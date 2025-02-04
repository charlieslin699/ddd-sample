package middleware

import (
	"ddd-sample/application/query/auth"
	"ddd-sample/pkg/httpserver"
	"ddd-sample/userinterface/api/common/context"
	"ddd-sample/userinterface/api/common/cookie"

	"github.com/gin-gonic/gin"
)

func HandleAuthorization(checkTokenQuery auth.CheckTokenQuery) httpserver.HandlerFunc {
	return func(ctx *gin.Context) (httpserver.RestfulResult, error) {
		authToken := cookie.AuthToken.Get(ctx)
		output, err := checkTokenQuery.Execute(ctx, auth.CheckTokenQueryInput{
			AuthToken: authToken,
		})
		if err != nil {
			return nil, err
		}

		// 寫入會員資料
		context.UserUID.Set(ctx, output.UID)
		context.Username.Set(ctx, output.Username)

		return httpserver.Next()
	}
}
