package senders

import (
	"context"
	"fmt"
	"github.com/lee-lou2/go/pkg/mongo"
	"github.com/lee-lou2/go/pkg/requests"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// insertQue 스케쥴러에 등록
func insertQue(request interface{}) error {
	client, collection, err := mongo.GetCollection("scheduler", "messages")
	if err != nil {
		return err
	}
	defer func() { _ = client.Disconnect(context.Background()) }()
	_, err = collection.Insert(bson.M{"message": request})
	return err
}

// checkIfTimePassed 시간이 지났는지 확인
func checkIfTimePassed(requestOneTime string) (bool, error) {
	oneTime, err := time.Parse("2006-01-02 15:04:05", requestOneTime)
	if err != nil {
		return false, err
	}
	oneTime = oneTime.Add(-9 * time.Hour)
	now := time.Now().UTC()
	if now.After(oneTime) {
		return true, nil
	} else {
		return false, nil
	}
}

// requestCallback 콜백 요청
func requestCallback(callbackUri string, successCount, failedCount int) error {
	// 전송 후 콜백 및 히스토리 저장
	if callbackUri != "" {
		if _, err := requests.Http(
			"GET",
			fmt.Sprintf("%s?success=%v&failed=%v", callbackUri, successCount, failedCount),
			nil,
		); err != nil {
			return err
		}
	}
	return nil
}
