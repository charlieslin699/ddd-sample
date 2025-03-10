package env

type EnvKey interface {
	Value() string
}
