package auth

import (
	"context"
	"ddd-sample/application/query"
	"ddd-sample/pkg/env"
	"ddd-sample/pkg/errorcode"
	"ddd-sample/pkg/localtime"
	"ddd-sample/pkg/token"
)

type CheckTokenQuery query.Query[CheckTokenQueryInput, CheckTokenQueryOutput]

type checkTokenQuery struct {
	env       env.Env
	localTime localtime.LocalTime
}

type CheckTokenQueryInput struct {
	AuthToken string
}

type CheckTokenQueryOutput struct {
	UID      string
	Username string
}

func NewCheckTokenQuery(env env.Env, localTime localtime.LocalTime) CheckTokenQuery {
	return &checkTokenQuery{
		env:       env,
		localTime: localTime,
	}
}

func (q *checkTokenQuery) Execute(ctx context.Context, input CheckTokenQueryInput) (CheckTokenQueryOutput, error) {
	// 取金鑰
	jwtKey := q.env.GetValue(env.AuthTokenKey)

	// 解析JWT
	claims, err := token.ParseAuthToken(input.AuthToken, []byte(jwtKey))
	if err != nil {
		return CheckTokenQueryOutput{}, errorcode.ErrUnlogin
	}

	// 檢查token是否有效
	if !claims.IsValid(q.localTime.NowTime()) {
		return CheckTokenQueryOutput{}, errorcode.ErrAuthTokenExpired
	}

	return CheckTokenQueryOutput{
		UID:      claims.UID(),
		Username: claims.Username(),
	}, nil
}
