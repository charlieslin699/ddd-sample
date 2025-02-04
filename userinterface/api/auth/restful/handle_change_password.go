package restful //nolint:dupl // 不同用途

import (
	"ddd-sample/application/command/auth"
	"ddd-sample/pkg/httpserver"
	"ddd-sample/userinterface/api/auth/model"
	"ddd-sample/userinterface/api/common/validation"

	"github.com/gin-gonic/gin"
)

func HandleChangePassword(changePasswordCommand auth.ChangePasswordCommand) httpserver.HandlerFunc {
	return func(ctx *gin.Context) (httpserver.RestfulResult, error) {
		// 參數驗證
		requestURI, err := validation.ValidateURI[model.PutAccountPasswordRequestURI](ctx)
		if err != nil {
			return nil, err
		}
		requestData, err := validation.Validate[model.PutAccountPasswordRequest](ctx)
		if err != nil {
			return nil, err
		}

		// 更新密碼
		_, err = changePasswordCommand.Execute(ctx, auth.ChangePasswordCommandInput{
			UID:         requestURI.UID,
			OldPassword: requestData.OldPassword,
			NewPassword: requestData.NewPassword,
		})

		return model.PutAccountPasswordResponse{}, err
	}
}
