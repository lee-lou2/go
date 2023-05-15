package senders

import (
	"fmt"
	"github.com/lee-lou2/go/apps/notify/cmd/schemas"
	"github.com/lee-lou2/go/pkg/email"
	"github.com/lee-lou2/go/pkg/logger"
	"github.com/lee-lou2/go/pkg/requests"
	"github.com/lee-lou2/go/platform/firebase"
	"github.com/lee-lou2/go/platform/slack"
	"github.com/lee-lou2/go/platform/twilio"
	"log"
)

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

// SlackProcessing 슬랙 전송
func SlackProcessing(q *schemas.RequestBody[schemas.SlackRequest]) {
	// 로그 기록
	status := logger.LogRequest{Raw: q}
	defer func() { go func() { logger.LogQue <- status }() }()

	err := slack.SendSlack(q.Data.Channel, q.Data.Message)
	if err != nil {
		status.Error = err.Error()
		log.Println(q.RequestID, "슬랙 전송간 오류 발생", err)
	}

	// 콜백 전송
	if err := requestCallback(q.CallbackUri, 1, 0); err != nil {
		log.Println(q.RequestID, "콜백 전송간 오류 발생", err)
	}

	log.Println(q.RequestID, "슬랙 전송 완료")
}
