package auth

import (
	"context"
	"ddd-sample/application/command"
	"ddd-sample/config/env"
	"ddd-sample/internal/auth/repository"
	pkgenv "ddd-sample/pkg/env"
	"ddd-sample/pkg/localtime"
)

// 系統登入
type LoginCommand command.Command[LoginCommandInput, LoginCommandOutput]

type loginCommand struct {
	identityRepository repository.IdentityRepository
	env                pkgenv.Env
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
func NewLoginCommand(identityRepository repository.IdentityRepository, env pkgenv.Env, localtime localtime.LocalTime) LoginCommand {
	return &loginCommand{
		identityRepository: identityRepository,
		env:                env,
		localtime:          localtime,
	}
}

// 登入
func (c *loginCommand) Execute(ctx context.Context, input LoginCommandInput) (LoginCommandOutput, error) {
	// 取身分資料
	identity, err := c.identityRepository.Find(ctx, input.Username)
	if err != nil {
		return LoginCommandOutput{}, err
	}

	// 檢查帳號密碼
	if isOK := identity.CheckIdentity(input.Username, input.Password); !isOK {
		// 寫入登入失敗紀錄
		err = c.identityRepository.SaveLoginFailedRecord(ctx, identity)
		return LoginCommandOutput{}, err
	}

	// 產生token
	token := identity.CreateToken([]byte(c.env.GetValue(env.AuthTokenKey)), c.localtime.NowTime())

	// 保存登入紀錄
	err = c.identityRepository.SaveLoginRecord(ctx, identity, token)
	if err != nil {
		return LoginCommandOutput{}, err
	}

	return LoginCommandOutput{
		IsLogin: true,
		Token:   token.TokenString,
	}, nil
}
