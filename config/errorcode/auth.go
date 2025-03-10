package errorcode

import "ddd-sample/pkg/errorcode"

var (
	// 未登入
	ErrUnlogin = errorcode.NewErrorCode("PL-01-00000", "未登入", "未登入")
	// 已登入
	ErrAlreadylogin = errorcode.NewErrorCode("PL-01-00001", "已登入", "已登入")
	// 登入失敗
	ErrLoginFailed = errorcode.NewErrorCode("PL-01-00002", "登入失敗", "登入失敗")
	// Token已到期
	ErrAuthTokenExpired = errorcode.NewErrorCode("PL-01-00003", "Token已到期", "Token已到期")
	// 舊密碼錯誤
	ErrOldPasswordError = errorcode.NewErrorCode("PL-01-00004", "舊密碼錯誤", "舊密碼錯誤")
	// 沒有權限
	ErrPermissionDenied = errorcode.NewErrorCode("PL-01-00005", "沒有權限", "沒有權限")
)
