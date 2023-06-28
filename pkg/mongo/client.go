package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewCollection 컬렉션
type NewCollection struct {
	*mongo.Collection
}

// getClient MongoDB 클라이언트 생성
func getClient(configs ...Configs) (*mongo.Client, error) {
	// MongoDB 클라이언트 설정
	cfg, err := setConfigs(configs...)

	clientOptions := options.Client().ApplyURI(cfg.Host)

	// 인증 정보 추가
	credential := options.Credential{
		Username: cfg.UserName,
		Password: cfg.Password,
	}
	clientOptions.Auth = &credential

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetCollection 컬렉션 선택
func GetCollection(dbName string, colName string) (*mongo.Client, *NewCollection, error) {
	// 컬렉션 선택
	client, err := getClient()
	if err != nil {
		return client, nil, err
	}
	collection := &NewCollection{client.Database(dbName).Collection(colName)}
	return client, collection, nil
}
