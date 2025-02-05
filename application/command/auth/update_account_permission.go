package auth

import (
	"context"
	"ddd-sample/application/command"
	"ddd-sample/internal/auth/repository"
	"ddd-sample/pkg/localtime"
)

type UpdateAccountPermissionCommand command.Command[UpdateAccountPermissionCommandInput, UpdateAccountPermissionCommandOutput]

type updateAccountPermissionCommand struct {
	accountPermissionRepository repository.AccountPermissionRepository
	localTime                   localtime.LocalTime
}

type UpdateAccountPermissionCommandInput struct {
	AccountUID     string
	PermissionUIDs []string
}

type UpdateAccountPermissionCommandOutput struct {
}

func NewUpdateAccountPermissionCommand(
	accountPermissionRepository repository.AccountPermissionRepository, localTime localtime.LocalTime,
) UpdateAccountPermissionCommand {
	return &updateAccountPermissionCommand{
		accountPermissionRepository: accountPermissionRepository,
		localTime:                   localTime,
	}
}

func (c *updateAccountPermissionCommand) Execute(
	ctx context.Context, input UpdateAccountPermissionCommandInput,
) (UpdateAccountPermissionCommandOutput, error) {
	// 取aggregate
	accountPermission, err := c.accountPermissionRepository.Find(ctx, input.AccountUID)
	if err != nil {
		return UpdateAccountPermissionCommandOutput{}, err
	}

	// 更新權限
	accountPermission.Update(input.PermissionUIDs, c.localTime.NowTime())

	// 儲存
	err = c.accountPermissionRepository.Update(ctx, accountPermission)

	return UpdateAccountPermissionCommandOutput{}, err
}
