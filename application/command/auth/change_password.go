package auth

import (
	"context"
	"ddd-sample/application/command"
	"ddd-sample/internal/auth/repository"
	"ddd-sample/pkg/errorcode"
	"ddd-sample/pkg/localtime"
)

type ChangePasswordCommand command.Command[ChangePasswordCommandInput, ChangePasswordCommandOutput]

type changePasswordCommand struct {
	accountRepository repository.AccountRepository
	localTime         localtime.LocalTime
}

type ChangePasswordCommandInput struct {
	UID         string
	OldPassword string
	NewPassword string
}

type ChangePasswordCommandOutput struct {
}

func NewChangePasswordCommand(accountRepository repository.AccountRepository, localTime localtime.LocalTime) ChangePasswordCommand {
	return &changePasswordCommand{
		accountRepository: accountRepository,
		localTime:         localTime,
	}
}

func (c *changePasswordCommand) Execute(ctx context.Context, input ChangePasswordCommandInput) (output ChangePasswordCommandOutput, err error) {
	// 取得帳號
	account, err := c.accountRepository.Find(input.UID)
	if err != nil {
		return
	}

	// 檢查舊密碼
	if !account.PasswordVerify(input.OldPassword) {
		err = errorcode.ErrOldPasswordError
		return
	}

	// 變更密碼
	account.ChangePassword(input.NewPassword, c.localTime.NowTime())
	err = c.accountRepository.ChangePassword(account)

	return
}
