package middleware

import (
	"ddd-sample/application/query/auth"
	"ddd-sample/pkg/httpserver"
	"ddd-sample/userinterface/api/auth/model"
	"ddd-sample/userinterface/api/common/cookie"
)

// 檢查是否可重新登入
func HandleCanLogin(checkTokenQuery auth.CheckTokenQuery) httpserver.HandlerFunc {
	return func(ctx *httpserver.Context) (httpserver.RestfulResult, error) {
		authToken := cookie.AuthToken.Get(ctx)
		if authToken == "" {
			return ctx.Next()
		}

		_, err := checkTokenQuery.Execute(ctx, auth.CheckTokenQueryInput{
			AuthToken: authToken,
		})
		if err != nil {
			return ctx.Next()
		}

		// 若已登入, 回傳原本的token
		return model.PostLoginResponse{AuthToken: authToken}, nil
	}
}
