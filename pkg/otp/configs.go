package otp

import (
	"fmt"
	"os"
)

// Configs OTP 설정
type Configs struct {
	otpSecretLength string
	otpAccountName  string
	otpIssuerName   string
}

// setConfigs OTP 설정
func setConfigs(defaultConfigs ...Configs) (*Configs, error) {
	configs := Configs{
		otpSecretLength: os.Getenv("OTP_SECRET_LENGTH"),
		otpAccountName:  os.Getenv("OTP_ACCOUNT_NAME"),
		otpIssuerName:   os.Getenv("OTP_ISSUER_NAME"),
	}
	if len(defaultConfigs) > 0 {
		configs = defaultConfigs[0]
	}
	if configs.otpSecretLength == "" || configs.otpAccountName == "" || configs.otpIssuerName == "" {
		return nil, fmt.Errorf("OTP 설정이 잘못되었습니다.")
	}
	return &configs, nil
}
