package i18n

type I18N interface {
	Get(key, locale string) (value string)
}

type i18n struct {
}

func NewI18n() I18N {
	return &i18n{}
}

func (i *i18n) Get(key, locale string) (value string) {
	return key // TODO: 待調整
}
