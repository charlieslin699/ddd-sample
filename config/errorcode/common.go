package errorcode

import "ddd-sample/pkg/errorcode"

var (
	// 系統錯誤，程式邏輯錯誤
	ErrSystemError = errorcode.NewErrorCode("PL-00-00000", "系統錯誤", "系統錯誤")
	// 非預期錯誤
	ErrUnexpected = errorcode.NewErrorCode("PL-00-00001", "非預期錯誤", "非預期錯誤")
	// JSON轉換錯誤
	ErrJSONMarshal = errorcode.NewErrorCode("PL-00-00002", "JSON轉換錯誤", "JSON轉換錯誤")
	// 設定檔讀取失敗
	ErrConfigLoading = errorcode.NewErrorCode("PL-00-00003", "設定檔讀取失敗", "設定檔讀取失敗")
	// 取不到Context保存資料
	ErrContextGetFailed = errorcode.NewErrorCode("PL-00-00004", "取不到Context保存資料", "取不到Context保存資料")
	// 列舉轉換錯誤
	ErrEnumConvert = errorcode.NewErrorCode("PL-00-00005", "列舉轉換錯誤", "列舉轉換錯誤")
	// DB資料不存在
	ErrDBRecordNotFound = errorcode.NewErrorCode("PL-00-00006", "資料不存在", "資料不存在")
)
