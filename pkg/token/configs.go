package token

import (
	"fmt"
	"os"
)

type Configs struct {
	JWTSecretKey string
}

func setConfigs(defaultConfigs ...Configs) (*Configs, error) {
	configs := Configs{
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
	if len(defaultConfigs) > 0 {
		configs = defaultConfigs[0]
	}
	if configs.JWTSecretKey == "" {
		return nil, fmt.Errorf("JWTSecretKey 설정이 잘못되었습니다.")
	}
	return &configs, nil
}
