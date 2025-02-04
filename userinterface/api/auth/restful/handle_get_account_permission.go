package restful

import (
	"ddd-sample/application/query/auth"
	"ddd-sample/pkg/httpserver"
	"ddd-sample/userinterface/api/auth/model"
	"ddd-sample/userinterface/api/common/validation"

	"github.com/gin-gonic/gin"
)

func HandleGetAccountPermission(getAccountPermissionQuery auth.GetAccountPermissionQuery) httpserver.HandlerFunc {
	return func(ctx *gin.Context) (httpserver.RestfulResult, error) {
		// 參數驗證
		requestURI, err := validation.ValidateURI[model.GetAccountPermissionRequestURI](ctx)
		if err != nil {
			return nil, err
		}

		// 取得帳號權限
		output, err := getAccountPermissionQuery.Execute(ctx, auth.GetAccountPermissionQueryInput{
			AccountUID: requestURI.UID,
		})

		responseData := model.GetAccountPermissionResponse{
			PermissionUIDs: output.PermissionUIDs,
		}

		return responseData, err
	}
}
