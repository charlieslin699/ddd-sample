package system

import (
	"github.com/spf13/viper"
)

func GetDBConfig(path string) (db DB, err error) {
	v := viper.New()
	v.SetConfigName("db")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	err = v.Unmarshal(&db)
	return
}
