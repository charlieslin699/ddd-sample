package adapter

import (
	"context"
	infradbauth "ddd-sample/infra/db/auth"
	infradbauthmodel "ddd-sample/infra/db/auth/model"
	"ddd-sample/internal/auth/aggregate"
	"ddd-sample/internal/auth/entity"
	"ddd-sample/internal/auth/repository"
	"ddd-sample/internal/auth/valueobject"
	coreadapter "ddd-sample/internal/core/adapter"
	corerepository "ddd-sample/internal/core/repository"
)

type accountPermissionRepository struct {
	dbAuth infradbauth.DBAuth
	corerepository.CoreRepository
}

// NewAccountPermissionRepository 建立帳號權限 repository
func NewAccountPermissionRepository(dbAuth infradbauth.DBAuth) repository.AccountPermissionRepository {
	return &accountPermissionRepository{
		dbAuth:         dbAuth,
		CoreRepository: coreadapter.NewCoreRepository(),
	}
}

// 取aggregate
func (repo *accountPermissionRepository) Find(ctx context.Context, accountUID string) (*aggregate.AccountPermission, error) {
	// 取帳號
	accountData, err := repo.dbAuth.GetAccount(ctx, accountUID)
	if err != nil {
		return nil, err
	}

	// 取得全部的權限
	allPermissionData, err := repo.dbAuth.GetAllPermission(ctx)
	if err != nil {
		return nil, err
	}

	// 取得帳號的權限
	accountPermissions, err := repo.dbAuth.GetAccountPermission(ctx, accountData.UID)
	if err != nil {
		return nil, err
	}

	// 轉換value object
	permissions := repo.parseToPermissionValueObject(allPermissionData, accountPermissions)

	// 組aggregate
	a := aggregate.BuildAccountPermission(
		entity.BuildAccount(
			accountData.UID,
			accountData.Username,
			accountData.Password,
			accountData.Secret,
			accountData.Status,
		), permissions)

	return a, nil
}

// 更新aggregate
func (repo *accountPermissionRepository) Update(ctx context.Context, a *aggregate.AccountPermission) error {
	// 轉換成db model
	accountPermissions := repo.parseToAccountPermissionModel(a.Account().UID(), a.Permissions())

	// 更新帳號權限
	if err := repo.dbAuth.UpdateAccountPermission(ctx, a.Account().UID(), accountPermissions); err != nil {
		return err
	}

	return repo.PubEvent(a)
}

// 轉換成value object
func (repo *accountPermissionRepository) parseToPermissionValueObject(
	allPermissionData []infradbauthmodel.Permission, accountPermissionData []infradbauthmodel.AccountPermission,
) valueobject.Permissions {
	// 組全部權限
	var allPermissions []entity.Permission
	for _, permissionData := range allPermissionData {
		allPermissions = append(allPermissions, entity.BuildPermission(permissionData.UID, permissionData.Name, permissionData.Key))
	}

	// 組帳號權限
	var permissionUIDs []string
	for _, accountPermissionData := range accountPermissionData {
		permissionUIDs = append(permissionUIDs, accountPermissionData.PermissionUID)
	}

	return valueobject.BuildPermissions(allPermissions, permissionUIDs)
}

// 轉換成db model
func (repo *accountPermissionRepository) parseToAccountPermissionModel(
	accountUID string, permissions valueobject.Permissions,
) []infradbauthmodel.AccountPermission {
	var accountPermissions []infradbauthmodel.AccountPermission
	for _, permissionUID := range permissions.AccountPermissionUIDs {
		accountPermissions = append(accountPermissions, infradbauthmodel.AccountPermission{
			AccountUID:    accountUID,
			PermissionUID: permissionUID,
		})
	}

	return accountPermissions
}
