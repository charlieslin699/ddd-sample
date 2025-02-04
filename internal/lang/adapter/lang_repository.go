package adapter

import (
	"ddd-sample/internal/lang/repository"
	pkgi18n "ddd-sample/pkg/config/i18n"
)

type langRepository struct {
	i18n pkgi18n.I18N
}

func NewLangRepository(i18n pkgi18n.I18N) repository.LangRepository {
	return &langRepository{
		i18n: i18n,
	}
}

func (repo *langRepository) Get(key, locale string, backupLocales ...string) string {
	v := repo.i18n.Get(key, locale)
	if v != key {
		return v
	}

	// 指定語系key不存在，使用備援語系
	for _, l := range backupLocales {
		v := repo.i18n.Get(key, l)
		if v != key {
			return v
		}
	}

	return key
}
