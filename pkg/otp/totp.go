package otp

import (
	"github.com/pquerna/otp/totp"
)

// GenerateURL 產生OTP URL
func GenerateURL(Issuer, accountName, secret string) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      Issuer,
		AccountName: accountName,
		Secret:      []byte(secret),
	})
	if err != nil {
		return "", err
	}

	return key.URL(), nil

}

// Verify 驗證OTP
func Verify(code, secret string) bool {
	return totp.Validate(code, secret)
}
