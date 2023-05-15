package otp

import (
	"github.com/xlzd/gotp"
	"strconv"
)

// CreateSecretKey OTP 시크릿키 생성
func CreateSecretKey(configs ...Configs) (string, error) {
	cfg, err := setConfigs(configs...)
	if err != nil {
		return "", err
	}
	otpSecretLength, _ := strconv.Atoi(cfg.otpSecretLength)
	return gotp.RandomSecret(otpSecretLength), nil
}

// GetOTPUri OTP Uri 조회
func GetOTPUri(secretKey string, configs ...Configs) (string, error) {
	cfg, err := setConfigs(configs...)
	if err != nil {
		return "", err
	}
	return gotp.NewDefaultTOTP(secretKey).ProvisioningUri(cfg.otpAccountName, cfg.otpIssuerName), nil
}
