package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"os"
)

// Configs AWS 설정
type Configs struct {
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	AWSRegion          string
}

// setConfigs AWS 설정
func setConfigs(defaultConfigs ...Configs) (*Configs, error) {
	configs := Configs{
		AWSAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AWSRegion:          os.Getenv("AWS_REGION"),
	}
	if len(defaultConfigs) > 0 {
		configs = defaultConfigs[0]
	}
	if configs.AWSAccessKeyID == "" || configs.AWSSecretAccessKey == "" || configs.AWSRegion == "" {
		return nil, fmt.Errorf("AWS 설정이 잘못되었습니다.")
	}
	return &configs, nil
}

// defaultConfig AWS 설정
func defaultConfig(configs ...Configs) (aws.Config, error) {
	cfg, err := setConfigs(configs...)
	if err != nil {
		return aws.Config{}, err
	}
	return config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.AWSRegion),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     cfg.AWSAccessKeyID,
				SecretAccessKey: cfg.AWSSecretAccessKey,
			},
		}),
	)
}
