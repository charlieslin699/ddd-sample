package valueobject

const defaultLimit = 5

type LoginFailedCount int

func (l *LoginFailedCount) Increase() {
	*l++
}

// OverLimit 超過次數鎖定帳號(5次)
func (l *LoginFailedCount) OverLimit() bool {
	return *l >= defaultLimit
}

// Value 取得次數
func (l LoginFailedCount) Value() int {
	return int(l)
}
