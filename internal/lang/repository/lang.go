package repository

type LangRepository interface {
	Get(key, locale string, backupLocales ...string) (value string)
}
