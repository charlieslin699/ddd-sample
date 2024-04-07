package lang

import (
	"context"
	"ddd-sample/application/query"
	"ddd-sample/pkg/config/i18n"
)

type GetLangsQuery query.Query[GetLangsQueryInput, GetLangsQueryOutput]

type getLangsQuery struct {
	i18n i18n.I18N
}

type GetLangsQueryInput struct {
	Keys   []string
	Locale string
}

type GetLangsQueryOutput struct {
	LangMap map[string]string
}

func NewGetLangsQuery(i18n i18n.I18N) GetLangsQuery {
	return &getLangsQuery{i18n}
}

func (q *getLangsQuery) Execute(ctx context.Context, input GetLangsQueryInput) (GetLangsQueryOutput, error) {
	output := GetLangsQueryOutput{
		LangMap: make(map[string]string),
	}

	for _, key := range input.Keys {
		output.LangMap[key] = q.i18n.Get(key, input.Locale)
	}

	return output, nil
}
