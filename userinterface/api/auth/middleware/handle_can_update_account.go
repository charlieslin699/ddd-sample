package middleware

import (
	"ddd-sample/config/errorcode"
	"ddd-sample/pkg/httpserver"
	"ddd-sample/userinterface/api/common/context"
)

func HandleCanUpdateAccount() httpserver.HandlerFunc {
	return func(ctx *httpserver.Context) (httpserver.RestfulResult, error) {
		// 登入帳號資料
		userUID, err := context.UserUID.Get(ctx)
		if err != nil {
			return nil, err
		}

		// 取得要更新的帳號UID
		uid := ctx.Param("uid")

		// 檢查是否為自己
		if userUID != uid {
			return nil, errorcode.ErrPermissionDenied
		}

		return ctx.Next()
	}
}
