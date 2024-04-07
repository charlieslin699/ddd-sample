package restful

import (
	"ddd-sample/pkg/httpserver"
	"ddd-sample/userinterface/api/auth/model"
	"ddd-sample/userinterface/api/common/validation"

	"ddd-sample/application/command/auth"

	"github.com/gin-gonic/gin"
)

func HandleCreateAccount(createAccountCommand auth.CreateAccountCommand) httpserver.HandlerFunc {
	return func(ctx *gin.Context) (httpserver.RestfulResult, error) {
		// 參數驗證
		requestData, err := validation.Validate[model.PostAccountRequest](ctx)
		if err != nil {
			return nil, err
		}

		output, err := createAccountCommand.Execute(ctx, auth.CreateAccountCommandInput{
			Username: requestData.Username,
			Password: requestData.Password,
		})
		if err != nil {
			return nil, err
		}

		// 新增成功
		return model.PostAccountResponse{
			UID: output.UID,
		}, nil
	}
}
