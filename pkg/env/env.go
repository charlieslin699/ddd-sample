package env

import (
	"os"
	"sync"
)

type Env interface {
	GetValue(e EnvKey) string                                 // 使用EnvKey取字典檔
	GetValueByKey(key string) (value string, isKeyExist bool) // 使用輸入的key取字典檔, 少用、維護性差
	MustGetValueByKey(key string) string                      // 使用輸入的key取字典檔, 少用、維護性差, 不存在時panic
}

type env struct {
	cache map[string]string
	mutex *sync.RWMutex
}

func NewEnv() Env {
	return &env{
		cache: make(map[string]string),
		mutex: new(sync.RWMutex),
	}
}

// GetValue 使用EnvKey取字典檔
func (_ env) GetValue(ek EnvKey) string { //nolint:stylecheck // 保持取資料流程一致
	return ek.Value()
}

// GetValueByKey 使用輸入的key取字典檔, 少用、維護性差
func (e *env) GetValueByKey(key string) (value string, isExist bool) {
	e.mutex.RLock()
	value, isExist = e.cache[key]
	e.mutex.RUnlock()

	if isExist {
		return
	}

	value = os.Getenv(key)
	isExist = value != ""
	e.saveKey(key, value)

	return
}

// MustGetValueByKey 使用輸入的key取字典檔, 少用、維護性差
func (e *env) MustGetValueByKey(key string) string {
	e.mutex.RLock()
	value, isExist := e.cache[key]
	e.mutex.RUnlock()

	if isExist {
		return value
	}

	value = os.Getenv(key)
	if value == "" {
		panic("env: key not found: " + key)
	}

	e.saveKey(key, value)

	return value
}

func (e *env) saveKey(key, value string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.cache[key] = value
}
