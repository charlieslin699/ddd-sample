package restful //nolint:dupl // 不同用途

import (
	"ddd-sample/application/command/auth"
	"ddd-sample/pkg/httpserver"
	"ddd-sample/userinterface/api/auth/model"
	"ddd-sample/userinterface/api/common/validation"
)

func HandleUpdateAccount(updateAccountCommand auth.UpdateAccountCommand) httpserver.HandlerFunc {
	return func(ctx *httpserver.Context) (httpserver.RestfulResult, error) {
		// 參數驗證
		requestURI, err := validation.ValidateURI[model.PutAccountRequestURI](ctx)
		if err != nil {
			return nil, err
		}
		requestData, err := validation.Validate[model.PutAccountRequest](ctx)
		if err != nil {
			return nil, err
		}

		// 更新帳號
		_, err = updateAccountCommand.Execute(ctx, auth.UpdateAccountCommandInput{
			UID:          requestURI.UID,
			Status:       requestData.Status,
			IsEnabledOTP: requestData.IsEnabledOTP,
		})

		return model.PutAccountResponse{}, err
	}
}
