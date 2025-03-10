package env

import (
	"ddd-sample/pkg/log"
	"fmt"
	"os"
)

var (
	GoPath       = newEnvKey("GOPATH", "./")                           // Go目錄
	ConfigPath   = newEnvKey("CONFIG_PATH", "/go/src/ddd-sample")      // 設定檔目錄
	ProjectName  = newEnvKey("PROJECT_NAME", "ddd-sample")             // 專案名稱
	ProjectEnv   = newEnvKey("PROJECT_ENV", "prod")                    // 系統環境 dev/qa/prod
	AuthTokenKey = newEnvKey("AUTH_TOKEN_KEY", "fPbhhYjgqTynF1EGnkgs") // 系統登入JWT金鑰
)

type EnvKey func() string

func newEnvKey(key, defaultValue string) EnvKey {
	value := os.Getenv(key)
	if value == "" {
		log.Debug(fmt.Sprintf("環境變數[%s]未設定, 使用預設值: %s\n", key, defaultValue))
		value = defaultValue
	}

	return func() string {
		return value
	}
}

func (ek EnvKey) Value() string {
	return ek()
}
