package senders

import (
	"github.com/lee-lou2/go/apps/notify/cmd/schemas"
	"github.com/lee-lou2/go/pkg/email"
	"github.com/lee-lou2/go/pkg/logger"
	"log"
	"time"
)

// SendEmail 이메일 전송
func SendEmail(body *schemas.RequestBody[schemas.EmailRequest]) {
	if body.Scheduler.SchedulerType == 0 {
		go EmailProcessing(body)
	} else if body.Scheduler.SchedulerType == 1 {
		if passed, err := checkIfTimePassed(body.Scheduler.OneTime); passed {
			go EmailProcessing(body)
		} else if !passed && err == nil {
			go func() { SendEmail(body) }()
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
			go func() { SendEmail(body) }()
		}
	}
	// 데이터 저장
	if err := insertQue(body); err != nil {
		log.Println(err)
		return
	}
}

// EmailProcessing 이메일 전송
func EmailProcessing(q *schemas.RequestBody[schemas.EmailRequest]) {
	var err error

	// 로그 기록
	status := logger.LogRequest{Raw: q}
	defer func() { go func() { logger.LogQue <- status }() }()

	// 이메일 전송
	data := q.Data
	failedCount := 0
	for _, emailAddress := range data.Emails {
		err = email.SendEmail(emailAddress, data.Subject, data.Message)
		if err != nil {
			failedCount += 1
			status.Error = err.Error()
			log.Println(q.RequestID, "이메일 전송간 오류 발생", err)
		}
	}

	// 콜백 전송
	if err := requestCallback(q.CallbackUri, len(data.Emails)-failedCount, failedCount); err != nil {
		log.Println(q.RequestID, "콜백 전송간 오류 발생", err)
	}

	log.Println(q.RequestID, "이메일 전송 완료")
}
