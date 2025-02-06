package auth

import (
	"context"
	"ddd-sample/application/query"
	infradbauth "ddd-sample/infra/db/auth"
	"ddd-sample/internal/auth/enum"
	"ddd-sample/pkg/env"
	"ddd-sample/pkg/otp"
)

type GetAccountQuery query.Query[GetAccountQueryInput, GetAccountQueryOutput]

type getAccountQuery struct {
	dbAuth infradbauth.DBAuth
	env    env.Env
}

type GetAccountQueryInput struct {
	UID string
}

type GetAccountQueryOutput struct {
	Username string
	Status   enum.AccountStatus
	OTPURL   string
}

func NewGetAccountQuery(dbAuth infradbauth.DBAuth, env env.Env) GetAccountQuery {
	return &getAccountQuery{
		dbAuth: dbAuth,
		env:    env,
	}
}

func (q *getAccountQuery) Execute(ctx context.Context, input GetAccountQueryInput) (output GetAccountQueryOutput, err error) {
	// 取帳號資料
	accountData, err := q.dbAuth.GetAccount(ctx, input.UID)
	if err != nil {
		return
	}

	// 取得帳號的 OTP URL
	url, err := otp.GenerateURL(q.env.GetValue(env.ProjectName), accountData.Username, accountData.Secret)
	if err != nil {
		return GetAccountQueryOutput{}, err
	}

	output.Username = accountData.Username
	output.Status = accountData.Status
	output.OTPURL = url

	return
}
