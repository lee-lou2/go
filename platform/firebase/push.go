package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"google.golang.org/api/option"
)

var client *messaging.Client
var ctx = context.Background()

// FCMClient FCM 클라이언트 생성
func FCMClient() (*messaging.Client, error) {
	if client != nil {
		return client, nil
	}
	opt := option.WithCredentialsFile("configs/keys/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	fmcClient, err := app.Messaging(ctx)
	if err != nil {
		return nil, err
	}
	client = fmcClient
	return client, nil
}

// SendPush 푸시 전송
func SendPush(title, body string, tokens []string) (int, int, error) {
	// 클라이언트 조회
	fmcClient, err := FCMClient()
	if err != nil {
		return 0, 0, err
	}

	// 메시지 지정
	messages := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Tokens: tokens,
	}

	resp, err := fmcClient.SendMulticast(ctx, messages)
	if resp.FailureCount > 0 {
		return resp.SuccessCount, resp.FailureCount, fmt.Errorf("failed to send message")
	}
	if err != nil {
		return resp.SuccessCount, resp.FailureCount, err
	}
	return resp.SuccessCount, resp.FailureCount, nil
}
