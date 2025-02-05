package auth

import (
	"context"
	"ddd-sample/application/command"
	"ddd-sample/internal/auth/enum"
	"ddd-sample/internal/auth/repository"
	"ddd-sample/pkg/localtime"
)

// UpdateAccountCommand 更新帳號
type UpdateAccountCommand command.Command[UpdateAccountCommandInput, UpdateAccountCommandOutput]

type updateAccountCommand struct {
	accountRepository repository.AccountRepository
	localTime         localtime.LocalTime
}

type UpdateAccountCommandInput struct {
	UID          string
	Status       uint
	IsEnabledOTP bool
}

type UpdateAccountCommandOutput struct {
}

// NewUpdateAccountCommand 建立更新帳號命令
func NewUpdateAccountCommand(accountRepository repository.AccountRepository, localTime localtime.LocalTime) UpdateAccountCommand {
	return &updateAccountCommand{
		accountRepository: accountRepository,
		localTime:         localTime,
	}
}

func (c *updateAccountCommand) Execute(
	ctx context.Context, input UpdateAccountCommandInput,
) (UpdateAccountCommandOutput, error) {
	// 取得帳號
	account, err := c.accountRepository.Find(ctx, input.UID)
	if err != nil {
		return UpdateAccountCommandOutput{}, err
	}

	// 轉換列舉
	accountStatus, err := enum.ConvertToAccountStatus(input.Status)
	if err != nil {
		return UpdateAccountCommandOutput{}, err
	}

	// 更新帳號
	account.Update(accountStatus, input.IsEnabledOTP, c.localTime.NowTime())
	err = c.accountRepository.Update(ctx, account)
	if err != nil {
		return UpdateAccountCommandOutput{}, err
	}

	return UpdateAccountCommandOutput{}, nil
}
