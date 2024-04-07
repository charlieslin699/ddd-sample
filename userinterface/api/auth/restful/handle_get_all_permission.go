package restful

import (
	"ddd-sample/application/query/auth"
	"ddd-sample/application/query/lang"
	"ddd-sample/pkg/httpserver"
	"ddd-sample/userinterface/api/auth/model"
	"ddd-sample/userinterface/api/common/cookie"

	"github.com/gin-gonic/gin"
)

func HandleGetAllPermission(getAllPermissionQuery auth.GetAllPermissionQuery, getLangQuery lang.GetLangQuery) httpserver.HandlerFunc {
	return func(ctx *gin.Context) (httpserver.RestfulResult, error) {
		// 取帳號語系
		locale := cookie.Locale.Get(ctx)

		// 取得所有權限
		output, err := getAllPermissionQuery.Execute(ctx, auth.GetAllPermissionQueryInput{})
		if err != nil {
			return nil, err
		}

		// 組回傳資料
		var responseData model.GetPermissionResponse
		for _, permission := range output.Permissions {
			// 取字典檔
			name, err := getLangQuery.Execute(ctx, lang.GetLangQueryInput{
				Key:    permission.LangIndex,
				Locale: locale,
			})
			if err != nil {
				return nil, err
			}

			responseData.Permissions = append(responseData.Permissions, model.Permission{
				UID:   permission.UID,
				Group: permission.Group,
				Name:  name,
			})
		}

		return responseData, nil
	}
}
