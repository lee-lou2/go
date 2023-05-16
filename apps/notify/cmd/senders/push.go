package senders

import (
	"fmt"
	"github.com/lee-lou2/go/apps/notify/cmd/schemas"
	"github.com/lee-lou2/go/pkg/logger"
	"github.com/lee-lou2/go/platform/firebase"
	"log"
	"time"
)

// SendPush 푸시 메시지 전송
func SendPush(body *schemas.RequestBody[schemas.PushRequest]) {
	if body.Scheduler.SchedulerType == 0 {
		go PushProcessing(body)
	} else if body.Scheduler.SchedulerType == 1 {
		if passed, err := checkIfTimePassed(body.Scheduler.OneTime); passed {
			go PushProcessing(body)
		} else if !passed && err == nil {
			go func() { SendPush(body) }()
		} else {
			log.Println(err)
			return
		}
	} else if body.Scheduler.SchedulerType == 2 {
		startTime, err := time.Parse("2006-01-02 15:04:05", body.Scheduler.Interval.StartTime)
		if err != nil {
			log.Println(err)
			return
		}
		for i := 0; i < body.Scheduler.Interval.Count; i++ {
			oneTime := startTime.Add(time.Duration(body.Scheduler.Interval.Duration) * time.Second * time.Duration(i))
			newRequest := *body
			newRequest.Scheduler.SchedulerType = 1
			newRequest.Scheduler.OneTime = oneTime.Format("2006-01-02 15:04:05")
			go func() { SendPush(body) }()
		}
	}
	// 데이터 저장
	if err := insertQue(body); err != nil {
		log.Println(err)
		return
	}
}

// PushProcessing 푸시 전송
func PushProcessing(q *schemas.RequestBody[schemas.PushRequest]) {
	if len(q.Data.FcmTokens) == 0 {
		log.Println(q.RequestID, "푸시 전송 대상이 없습니다.")
		return
	}
	// 로그 기록
	status := logger.LogRequest{Raw: q}
	defer func() { go func() { logger.LogQue <- status }() }()

	// 푸시 전송
	successCount, failedCount, err := firebase.SendPush(q.Data.Title, q.Data.Message, q.Data.FcmTokens)
	if err != nil {
		errMessage := fmt.Sprintf("[failed count : %d] %s", failedCount, err.Error())
		status.Error = errMessage
		log.Println(q.RequestID, "푸시 전송간 오류 발생", err)
	}

	// 콜백 전송
	if err := requestCallback(q.CallbackUri, successCount, failedCount); err != nil {
		log.Println(q.RequestID, "콜백 전송간 오류 발생", err)
	}

	log.Println(q.RequestID, "푸시 전송 완료")
}
