package auth

import (
	"context"
	"ddd-sample/application/query"
	infradbauth "ddd-sample/infra/db/auth"
)

// GetAccountPermissionQuery 取得帳號權限
type GetAccountPermissionQuery query.Query[GetAccountPermissionQueryInput, GetAccountPermissionQueryOutput]

type getAccountPermissionQuery struct {
	dbAuth infradbauth.DBAuth
}

type GetAccountPermissionQueryInput struct {
	AccountUID string
}

type GetAccountPermissionQueryOutput struct {
	PermissionUIDs []string
}

func NewGetAccountPermissionQuery(dbAuth infradbauth.DBAuth) GetAccountPermissionQuery {
	return &getAccountPermissionQuery{
		dbAuth: dbAuth,
	}
}

func (q *getAccountPermissionQuery) Execute(
	_ context.Context, input GetAccountPermissionQueryInput,
) (output GetAccountPermissionQueryOutput, err error) {
	// 取得帳號權限
	accountPermissionData, err := q.dbAuth.GetAccountPermission(input.AccountUID)
	if err != nil {
		return
	}

	for _, accountPermission := range accountPermissionData {
		output.PermissionUIDs = append(output.PermissionUIDs, accountPermission.PermissionUID)
	}

	return
}
