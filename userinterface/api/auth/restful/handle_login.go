package restful

import (
	"ddd-sample/application/command/auth"
	"ddd-sample/pkg/errorcode"
	"ddd-sample/pkg/httpserver"
	"ddd-sample/userinterface/api/auth/model"
	"ddd-sample/userinterface/api/common/cookie"
	"ddd-sample/userinterface/api/common/validation"

	"github.com/gin-gonic/gin"
)

// @summary      登入
// @description  帳號登入
// @tags         auth
// @accept       json
// @produce      json
// @param request body model.PostLoginRequest true "參數"
// @router       /login [post]
func HandleLogin(loginCommand auth.LoginCommand) httpserver.HandlerFunc {
	return func(ctx *gin.Context) (httpserver.RestfulResult, error) {
		// 取參數
		requestData, err := validation.Validate[model.PostLoginRequest](ctx)
		if err != nil {
			return nil, err
		}

		// 登入
		output, err := loginCommand.Execute(ctx, auth.LoginCommandInput{
			Username: requestData.Username,
			Password: requestData.Password,
		})
		if err != nil {
			return nil, err
		}

		// 登入失敗
		if !output.IsLogin {
			return nil, errorcode.ErrLoginFailed
		}

		// 設定Cookie
		cookie.AuthToken.Set(ctx, output.Token)

		responseData := model.PostLoginResponse{
			AuthToken: output.Token,
		}

		return responseData, nil
	}
}
