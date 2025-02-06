package enum

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
