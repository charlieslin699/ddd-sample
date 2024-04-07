package auth

import (
	"fmt"
	commandauth "ddd-sample/application/command/auth"
	queryauth "ddd-sample/application/query/auth"
	querylang "ddd-sample/application/query/lang"
	"ddd-sample/infra/db"
	dbauth "ddd-sample/infra/db/auth"
	"ddd-sample/internal/auth/adapter"
	pkgi18n "ddd-sample/pkg/config/i18n"
	pkgsystem "ddd-sample/pkg/config/system"
	pkgenv "ddd-sample/pkg/env"
	"ddd-sample/pkg/httpserver"
	pkglocaltime "ddd-sample/pkg/localtime"
	"ddd-sample/userinterface/api/auth/middleware"
	"ddd-sample/userinterface/api/auth/restful"
	"ddd-sample/userinterface/api/common/errorhandler"
	commonmiddleware "ddd-sample/userinterface/api/common/middleware"
	"ddd-sample/userinterface/api/common/panichandler"
)

func InitRouter(server httpserver.HttpServer) {
	// pkg
	localTime := pkglocaltime.NewLocalTime()
	env := pkgenv.NewEnv()
	i18n := pkgi18n.NewI18n()

	// config
	dbConfigPath := fmt.Sprintf(`%s/config/system/%s`, env.GetValue(pkgenv.ConfigPath), env.GetValue(pkgenv.ProjectEnv))
	dbConfig, err := pkgsystem.GetDBConfig(dbConfigPath)
	if err != nil {
		panic(err)
	}

	// DBConn
	accountDBConn, err := db.NewMySQLConn(
		db.WithDBName(dbConfig.Auth.DBName),
		db.WithAddr(getEnv(env, dbConfig.Auth.AddrEnv)),
		db.WithUsername(getEnv(env, dbConfig.Auth.UsernameEnv)),
		db.WithPassword(getEnv(env, dbConfig.Auth.PasswordEnv)),
	)
	if err != nil {
		panic(err)
	}

	// infra
	dbAuth := dbauth.NewMySQLAuth(accountDBConn)

	// repository
	identityRepo := adapter.NewIdentityRepository(dbAuth)
	accountRepo := adapter.NewAccountRepository(dbAuth)

	// command 存取資料
	loginCommand := commandauth.NewLoginCommand(identityRepo, env, localTime)
	createAccountCommand := commandauth.NewCreateAccountCommand(accountRepo, localTime)
	updateAccountCommand := commandauth.NewUpdateAccountCommand(accountRepo, localTime)
	changePasswordCommand := commandauth.NewChangePasswordCommand(accountRepo, localTime)

	// query 單純取資料或單純邏輯計算
	checkTokenQuery := queryauth.NewCheckTokenQuery(env, localTime)
	getAllPermission := queryauth.NewGetAllPermissionQuery(dbAuth)
	getAccountPermission := queryauth.NewGetAccountPermissionQuery(dbAuth)
	langQuery := querylang.NewGetLangQuery(i18n)

	// 註冊路由
	server.Route(
		httpserver.POST("/login",
			middleware.HandleCanLogin(checkTokenQuery),
			restful.HandleLogin(loginCommand),
		),

		httpserver.Group("/auth",
			// 身分驗證
			httpserver.Use(commonmiddleware.HandleAuthorization(checkTokenQuery)),

			// 取得所有權限
			httpserver.GET("/permission", restful.HandleGetAllPermission(getAllPermission, langQuery)),

			// 建立帳號
			httpserver.POST("/account", restful.HandleCreateAccount(createAccountCommand)),

			// 編輯帳號
			httpserver.PUT("/account/:uid", middleware.HandleCanUpdateAccount(), restful.HandleUpdateAccount(updateAccountCommand)),

			// 變更密碼
			httpserver.PUT("/account/:uid/password", middleware.HandleCanUpdateAccount(), restful.HandleChangePassword(changePasswordCommand)),

			// 取得帳號權限
			httpserver.GET("/account/:uid/permission", restful.HandleGetAccountPermission(getAccountPermission)),
		),
	)

	// 錯誤處理
	server.Catch(
		httpserver.CatchError(errorhandler.HandleError(langQuery)),
		httpserver.CatchPanic(panichandler.HandlePanic(langQuery)),
	)
}

func getEnv(env pkgenv.Env, key string) string {
	value, isExist := env.GetValueByKey(key)
	if !isExist {
		panic(fmt.Sprintf("環境變數[%s]未設定!", key))
	}

	return value
}
