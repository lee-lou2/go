package logger

import (
	"context"
	"encoding/json"
	"github.com/lee-lou2/go/pkg/mongo"
	"github.com/natefinch/lumberjack"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

// LogRequest 로그 요청
type LogRequest struct {
	Raw   interface{}
	Error string
}

var LogQue = make(chan LogRequest)

// RunLoggingSystem 로그 채널 소비
func RunLoggingSystem() {
	// 신규 로거 생성
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "logs.log", // 로그 파일 이름
		MaxSize:    100,        // 파일 당 최대 크기(MB)
		MaxBackups: 3,          // 최대 백업 파일 수
		MaxAge:     120,        // 파일 보관 일수
		Compress:   true,       // 로그 파일 압축 여부
	}
	logger := log.New(lumberjackLogger, "[ACCESS LOG] ", log.Ldate|log.Ltime)

	// Mongo 클라이언트
	client, collection, err := mongo.GetCollection("logs", "messages")
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	// 순차적 로그 생성
	for q := range LogQue {
		if _, err := collection.Insert(
			bson.M{"message": q},
		); err != nil {
			log.Println(err)
		}
		data, _ := json.Marshal(&q)
		logger.Println(string(data))
	}
}
