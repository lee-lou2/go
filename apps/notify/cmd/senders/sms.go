package senders

import (
	"github.com/lee-lou2/go/apps/notify/cmd/schemas"
	"github.com/lee-lou2/go/pkg/logger"
	"github.com/lee-lou2/go/platform/twilio"
	"log"
	"time"
)

// SendSMS SMS 전송
func SendSMS(body *schemas.RequestBody[schemas.SMSRequest]) {
	if body.Scheduler.SchedulerType == 0 {
		go SMSProcessing(body)
	} else if body.Scheduler.SchedulerType == 1 {
		if passed, err := checkIfTimePassed(body.Scheduler.OneTime); passed {
			go SMSProcessing(body)
		} else if !passed && err == nil {
			go func() { SendSMS(body) }()
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
			go func() { SendSMS(body) }()
		}
	}
	// 데이터 저장
	if err := insertQue(body); err != nil {
		log.Println(err)
		return
	}
}

// SMSProcessing 문자 전송
func SMSProcessing(q *schemas.RequestBody[schemas.SMSRequest]) {
	var err error

	// 로그 기록
	status := logger.LogRequest{Raw: q}
	defer func() { go func() { logger.LogQue <- status }() }()

	for _, phone := range q.Data.Phones {
		err = twilio.SendSMSTwilio(phone, q.Data.Message)
		if err != nil {
			status.Error = err.Error()
			log.Println(q.RequestID, "문자 전송간 오류 발생", err)
		}
	}

	// 콜백 전송
	if err := requestCallback(q.CallbackUri, len(q.Data.Phones), 0); err != nil {
		log.Println(q.RequestID, "콜백 전송간 오류 발생", err)
	}

	log.Println(q.RequestID, "문자 전송 완료")
}
