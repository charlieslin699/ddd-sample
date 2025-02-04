package enum

import "ddd-sample/pkg/errorcode"

// AccountStatus 帳號狀態
type AccountStatus uint

const (
	// AccountStatusClose 關閉
	AccountStatusClose AccountStatus = iota
	// AccountStatusNormal 正常
	AccountStatusNormal
	// AccountStatusLocked 鎖定
	AccountStatusLocked
)

// ConvertToAccountStatus 轉換為帳號狀態
func ConvertToAccountStatus(n uint) (AccountStatus, error) {
	e := AccountStatus(n)
	switch e {
	case AccountStatusNormal, AccountStatusLocked:
		return e, nil
	default:
		return AccountStatusClose, errorcode.ErrEnumConvert
	}
}

// Value 取值
func (g AccountStatus) Value() uint {
	return uint(g)
}
