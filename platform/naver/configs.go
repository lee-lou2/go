package naver

import (
	"fmt"
	"os"
)

// Configs AWS 설정
type Configs struct {
	NaverClientID     string
	NaverClientSecret string
}

// setConfigs AWS 설정
func setConfigs(defaultConfigs ...Configs) (*Configs, error) {
	configs := Configs{
		NaverClientID:     os.Getenv("NAVER_CLIENT_ID"),
		NaverClientSecret: os.Getenv("NAVER_CLIENT_SECRET"),
	}
	if len(defaultConfigs) > 0 {
		configs = defaultConfigs[0]
	}
	if configs.NaverClientID == "" || configs.NaverClientSecret == "" {
		return nil, fmt.Errorf("NAVER 설정이 잘못되었습니다.")
	}
	return &configs, nil
}
