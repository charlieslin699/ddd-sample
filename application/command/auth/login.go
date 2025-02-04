package auth

import (
	"context"
	"ddd-sample/application/command"
	"ddd-sample/internal/auth/repository"
	"ddd-sample/pkg/env"
	"ddd-sample/pkg/localtime"
)

// 系統登入
type LoginCommand command.Command[LoginCommandInput, LoginCommandOutput]

type loginCommand struct {
	identityRepository repository.IdentityRepository
	env                env.Env
	localtime          localtime.LocalTime
}

type LoginCommandInput struct {
	Username string
	Password string
}

type LoginCommandOutput struct {
	IsLogin bool
	Token   string // JWT token
}

// 工廠
func NewLoginCommand(identityRepository repository.IdentityRepository, env env.Env, localtime localtime.LocalTime) LoginCommand {
	return &loginCommand{
		identityRepository: identityRepository,
		env:                env,
		localtime:          localtime,
	}
}

// 登入
func (c *loginCommand) Execute(_ context.Context, input LoginCommandInput) (output LoginCommandOutput, err error) {
	// 取身分資料
	identity, err := c.identityRepository.Find(input.Username)
	if err != nil {
		return
	}

	// 檢查帳號密碼
	if isOK := identity.CheckIdentity(input.Username, input.Password); !isOK {
		// 寫入登入失敗紀錄
		err = c.identityRepository.SaveLoginFailedRecord(identity)
		return
	}

	// 產生token
	token := identity.CreateToken([]byte(c.env.GetValue(env.AuthTokenKey)), c.localtime.NowTime())

	// 保存登入紀錄
	err = c.identityRepository.SaveLoginRecord(identity, token)
	if err != nil {
		return
	}

	// 組資料
	output = LoginCommandOutput{
		IsLogin: true,
		Token:   token.TokenString,
	}

	return
}
