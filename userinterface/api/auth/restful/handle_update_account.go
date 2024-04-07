package restful

import (
	"ddd-sample/application/command/auth"
	"ddd-sample/pkg/httpserver"
	"ddd-sample/userinterface/api/auth/model"
	"ddd-sample/userinterface/api/common/validation"

	"github.com/gin-gonic/gin"
)

func HandleUpdateAccount(updateAccountCommand auth.UpdateAccountCommand) httpserver.HandlerFunc {
	return func(ctx *gin.Context) (httpserver.RestfulResult, error) {
		// 參數驗證
		requestUri, err := validation.ValidateUri[model.PutAccountRequestUri](ctx)
		if err != nil {
			return nil, err
		}
		requestData, err := validation.Validate[model.PutAccountRequest](ctx)
		if err != nil {
			return nil, err
		}

		// 更新帳號
		_, err = updateAccountCommand.Execute(ctx, auth.UpdateAccountCommandInput{
			UID:          requestUri.UID,
			Status:       requestData.Status,
			IsEnabledOTP: requestData.IsEnabledOTP,
		})

		return model.PutAccountResponse{}, err
	}
}
