package system

type DB struct {
	Auth DBConfig `mapstructure:"auth"`
}

type DBConfig struct {
	DBName      string `mapstructure:"dbName"`
	AddrEnv     string `mapstructure:"addrEnv"`
	UsernameEnv string `mapstructure:"usernameEnv"`
	PasswordEnv string `mapstructure:"passwordEnv"`
}
