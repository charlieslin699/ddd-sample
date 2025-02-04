package enum

import "ddd-sample/pkg/errorcode"

type ThirdPartyVerification uint

const (
	// ThirdPartyVerificationOTP OTP
	ThirdPartyVerificationOTP ThirdPartyVerification = iota
)

// ConvertToThirdPartyVerification 轉換為第三方驗證
func ConvertToThirdPartyVerification(n uint) (ThirdPartyVerification, error) {
	e := ThirdPartyVerification(n)
	switch e {
	case ThirdPartyVerificationOTP:
		return e, nil
	default:
		return e, errorcode.ErrEnumConvert
	}
}
