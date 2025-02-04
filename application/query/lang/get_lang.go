package lang

import (
	"context"
	"ddd-sample/application/query"
	"ddd-sample/pkg/config/i18n"
)

type GetLangQuery query.Query[GetLangQueryInput, string]

type getLangQuery struct {
	i18n i18n.I18N
}

type GetLangQueryInput struct {
	Key    string
	Locale string
}

func NewGetLangQuery(i18n i18n.I18N) GetLangQuery {
	return &getLangQuery{i18n}
}

func (q *getLangQuery) Execute(_ context.Context, input GetLangQueryInput) (string, error) {
	value := q.i18n.Get(input.Key, input.Locale)
	return value, nil
}
