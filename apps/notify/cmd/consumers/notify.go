package consumers

import (
	"context"
	"fmt"
	"github.com/lee-lou2/go/apps/notify/cmd/schemas"
	"github.com/lee-lou2/go/apps/notify/cmd/senders"
	"github.com/lee-lou2/go/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// NotifyQue 알림 큐
var NotifyQue = make(chan interface{})

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

// NotifyConsumer 알림 소비
func NotifyConsumer() error {
	for {
		body := <-NotifyQue
		switch body.(type) {
		case *schemas.RequestBody[schemas.PushRequest]:
			request := body.(*schemas.RequestBody[schemas.PushRequest])
			if request.Scheduler.SchedulerType == 0 {
				go senders.PushProcessing(request)
			} else if request.Scheduler.SchedulerType == 1 {
				if passed, err := checkIfTimePassed(request.Scheduler.OneTime); passed {
					go senders.PushProcessing(request)
				} else if !passed && err == nil {
					go func() { NotifyQue <- request }()
				} else {
					return err
				}
			} else if request.Scheduler.SchedulerType == 2 {
				startTime, err := time.Parse("2006-01-02 15:04:05", request.Scheduler.Interval.StartTime)
				if err != nil {
					return err
				}
				for i := 0; i < request.Scheduler.Interval.Count; i++ {
					oneTime := startTime.Add(time.Duration(request.Scheduler.Interval.Duration) * time.Second * time.Duration(i))
					newRequest := *request
					newRequest.Scheduler.SchedulerType = 1
					newRequest.Scheduler.OneTime = oneTime.Format("2006-01-02 15:04:05")
					go func() { NotifyQue <- &newRequest }()
				}
			}
		case *schemas.RequestBody[schemas.EmailRequest]:
			request := body.(*schemas.RequestBody[schemas.EmailRequest])
			if request.Scheduler.SchedulerType == 0 {
				go senders.EmailProcessing(request)
			} else if request.Scheduler.SchedulerType == 1 {
				if passed, err := checkIfTimePassed(request.Scheduler.OneTime); passed {
					go senders.EmailProcessing(request)
				} else if !passed && err == nil {
					go func() { NotifyQue <- request }()
				} else {
					return err
				}
			} else if request.Scheduler.SchedulerType == 2 {
				startTime, err := time.Parse("2006-01-02 15:04:05", request.Scheduler.Interval.StartTime)
				if err != nil {
					return err
				}
				for i := 0; i < request.Scheduler.Interval.Count; i++ {
					oneTime := startTime.Add(time.Duration(request.Scheduler.Interval.Duration) * time.Second * time.Duration(i))
					newRequest := *request
					newRequest.Scheduler.SchedulerType = 1
					newRequest.Scheduler.OneTime = oneTime.Format("2006-01-02 15:04:05")
					go func() { NotifyQue <- &newRequest }()
				}
			}
		case *schemas.RequestBody[schemas.SMSRequest]:
			request := body.(*schemas.RequestBody[schemas.SMSRequest])
			if request.Scheduler.SchedulerType == 0 {
				go senders.SMSProcessing(request)
			} else if request.Scheduler.SchedulerType == 1 {
				if passed, err := checkIfTimePassed(request.Scheduler.OneTime); passed {
					go senders.SMSProcessing(request)
				} else if !passed && err == nil {
					go func() { NotifyQue <- request }()
				} else {
					return err
				}
			} else if request.Scheduler.SchedulerType == 2 {
				startTime, err := time.Parse("2006-01-02 15:04:05", request.Scheduler.Interval.StartTime)
				if err != nil {
					return err
				}
				for i := 0; i < request.Scheduler.Interval.Count; i++ {
					oneTime := startTime.Add(time.Duration(request.Scheduler.Interval.Duration) * time.Second * time.Duration(i))
					newRequest := *request
					newRequest.Scheduler.SchedulerType = 1
					newRequest.Scheduler.OneTime = oneTime.Format("2006-01-02 15:04:05")
					go func() { NotifyQue <- &newRequest }()
				}
			}
		case *schemas.RequestBody[schemas.SlackRequest]:
			request := body.(*schemas.RequestBody[schemas.SlackRequest])
			if request.Scheduler.SchedulerType == 0 {
				go senders.SlackProcessing(request)
			} else if request.Scheduler.SchedulerType == 1 {
				if passed, err := checkIfTimePassed(request.Scheduler.OneTime); passed {
					go senders.SlackProcessing(request)
				} else if !passed && err == nil {
					go func() { NotifyQue <- request }()
				} else {
					return err
				}
			} else if request.Scheduler.SchedulerType == 2 {
				startTime, err := time.Parse("2006-01-02 15:04:05", request.Scheduler.Interval.StartTime)
				if err != nil {
					return err
				}
				for i := 0; i < request.Scheduler.Interval.Count; i++ {
					oneTime := startTime.Add(time.Duration(request.Scheduler.Interval.Duration) * time.Second * time.Duration(i))
					newRequest := *request
					newRequest.Scheduler.SchedulerType = 1
					newRequest.Scheduler.OneTime = oneTime.Format("2006-01-02 15:04:05")
					go func() { NotifyQue <- &newRequest }()
				}
			}
		default:
			return fmt.Errorf("unknown request type")
		}
		// 데이터 저장
		if err := insertQue(body); err != nil {
			return err
		}
	}
}
