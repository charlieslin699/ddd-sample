package auth

import (
	"context"
	"ddd-sample/application/command"
	"ddd-sample/internal/auth/repository"
	"ddd-sample/pkg/localtime"
)

type CreateAccountCommand command.Command[CreateAccountCommandInput, CreateAccountCommandOutput]

type createAccountCommand struct {
	accountRepository repository.AccountRepository
	localTime         localtime.LocalTime
}

type CreateAccountCommandInput struct {
	Username string
	Password string
}

type CreateAccountCommandOutput struct {
	UID string
}

func NewCreateAccountCommand(accountRepository repository.AccountRepository, localTime localtime.LocalTime) CreateAccountCommand {
	return &createAccountCommand{
		accountRepository: accountRepository,
		localTime:         localTime,
	}
}

func (c *createAccountCommand) Execute(
	ctx context.Context, input CreateAccountCommandInput,
) (CreateAccountCommandOutput, error) {
	// 建立帳號
	account := c.accountRepository.New(input.Username, input.Password, c.localTime.NowTime())

	// 儲存帳號
	err := c.accountRepository.Add(ctx, account)
	if err != nil {
		return CreateAccountCommandOutput{}, err
	}

	return CreateAccountCommandOutput{
		UID: account.Account().UID(),
	}, nil
}
