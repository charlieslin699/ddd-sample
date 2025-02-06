package valueobject

import "ddd-sample/internal/auth/enum"

type ThirdPartyVerification struct {
	OTP bool
}

func NewThirdPartyVerification() ThirdPartyVerification {
	return ThirdPartyVerification{}
}

func Build(enumValues ...enum.ThirdPartyVerification) ThirdPartyVerification {
	result := ThirdPartyVerification{}

	for _, v := range enumValues {
		result.Enable(v)
	}

	return result
}

func (t *ThirdPartyVerification) Enable(e enum.ThirdPartyVerification) {
	if e == enum.ThirdPartyVerificationOTP {
		t.OTP = true
	}
}
