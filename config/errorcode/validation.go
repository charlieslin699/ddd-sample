package errorcode

import "ddd-sample/pkg/errorcode"

var (
	// 參數驗證失敗
	ErrValidation = errorcode.NewErrorCode("PL-03-00000", "參數驗證失敗", "參數驗證失敗")
)
