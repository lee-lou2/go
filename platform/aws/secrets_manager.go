package aws

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// SecretManager AWS Secret Manager
func SecretManager(secretName string, configs ...Configs) (map[string]string, error) {
	cfg, err := defaultConfig(configs...)
	if err != nil {
		return nil, err
	}

	svc := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	secretString := *result.SecretString
	var jsonMap map[string]string
	err = json.Unmarshal([]byte(secretString), &jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}
