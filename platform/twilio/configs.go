package twilio

import (
	"fmt"
	"os"
)

type Configs struct {
	AccountSid string
	AuthToken  string
	FromNumber string
}

// setConfigs Twilio 설정
func setConfigs(defaultConfigs ...Configs) (*Configs, error) {
	configs := Configs{
		AccountSid: os.Getenv("TWILIO_SID"),
		AuthToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
		FromNumber: os.Getenv("TWILIO_FROM_NUMBER"),
	}
	if len(defaultConfigs) > 0 {
		configs = defaultConfigs[0]
	}
	if configs.AccountSid == "" || configs.FromNumber == "" || configs.AuthToken == "" {
		return nil, fmt.Errorf("Twilio 설정이 잘못되었습니다.")
	}
	return &configs, nil
}
