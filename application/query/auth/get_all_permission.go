package auth

import (
	"context"
	"ddd-sample/application/query"
	infradbauth "ddd-sample/infra/db/auth"
)

// GetAllPermissionQuery 取得所有權限
type GetAllPermissionQuery query.Query[GetAllPermissionQueryInput, GetAllPermissionQueryOutput]

type getAllPermissionQuery struct {
	dbAuth infradbauth.DBAuth
}

type GetAllPermissionQueryInput struct {
}

type GetAllPermissionQueryOutput struct {
	Permissions []Permission
}

type Permission struct {
	UID       string // 權限UID
	Group     string // 權限群組
	LangIndex string // 字典檔索引
}

func NewGetAllPermissionQuery(dbAuth infradbauth.DBAuth) GetAllPermissionQuery {
	return &getAllPermissionQuery{
		dbAuth: dbAuth,
	}
}

func (q *getAllPermissionQuery) Execute(ctx context.Context, _ GetAllPermissionQueryInput) (output GetAllPermissionQueryOutput, err error) {
	// 取得所有權限
	permissionData, err := q.dbAuth.GetAllPermission(ctx)
	if err != nil {
		return
	}

	for _, permssion := range permissionData {
		output.Permissions = append(output.Permissions, Permission{
			UID:       permssion.UID,
			Group:     permssion.Group,
			LangIndex: permssion.LangIndex,
		})
	}

	return
}
