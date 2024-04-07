package valueobject

import "ddd-sample/internal/auth/entity"

type Permissions struct {
	AllPermissions        map[string]entity.Permission // 所有權限 key: UID
	AccountPermissionUIDs []string
}

func BuildPermissions(allPermissions []entity.Permission, accountPermissionUIDs []string) Permissions {
	p := Permissions{}

	for _, permission := range allPermissions {
		p.AllPermissions[permission.UID()] = permission

	}

	p.Update(accountPermissionUIDs)

	return p
}

func (p *Permissions) Update(newPermissionUIDS []string) {
	clear(p.AccountPermissionUIDs)

	for _, permissionUID := range newPermissionUIDS {
		if _, isExist := p.AllPermissions[permissionUID]; isExist {
			p.AccountPermissionUIDs = append(p.AccountPermissionUIDs, permissionUID)
		}
	}
}
