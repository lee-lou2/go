package security

import (
	"fmt"
	"os"
)

// Configs 정보
type Configs struct {
	AESCipherKey string
}

func SetConfigs(defaultConfigs ...Configs) (*Configs, error) {
	configs := Configs{
		AESCipherKey: os.Getenv("AES_CIPHER_KEY"),
	}
	if len(defaultConfigs) > 0 {
		configs = defaultConfigs[0]
	}
	if configs.AESCipherKey == "" {
		return nil, fmt.Errorf("AESCipherKey 설정이 잘못되었습니다.")
	}
	return &configs, nil
}
