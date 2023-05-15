package slack

import (
	"fmt"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"log"
	"sync"
)

// SocketModeSlackBot 슬랙 봇
func SocketModeSlackBot(messageFunction func(string) (string, error), configs ...Configs) error {
	cfg, err := setConfigs(configs...)
	if err != nil {
		return err
	}
	api := slack.New(cfg.BotToken, slack.OptionAppLevelToken(cfg.AppToken))
	client := socketmode.New(api)
	// 봇 자신의 ID 가져오기
	botInfo, err := api.AuthTest()
	if err != nil {
		log.Fatalf("봇 정보를 가져오는데 실패했습니다: %v", err)
	}
	botID := botInfo.UserID

	// 처리 중인 메시지 추적을 위한 맵과 뮤텍스 생성
	var processingMessages sync.Map

	go func() {
		for evt := range client.Events {
			switch evt.Type {
			case socketmode.EventTypeEventsAPI:
				eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
				if !ok {
					fmt.Printf("소켓 모드 이벤트가 아닙니다: %v\n", evt)
					continue
				}

				if eventsAPIEvent.Type == slackevents.CallbackEvent {
					innerEvent := eventsAPIEvent.InnerEvent
					switch ev := innerEvent.Data.(type) {
					case *slackevents.MessageEvent:
						// 봇이 작성한 메시지는 처리하지 않음
						if ev.User == botID {
							continue
						}

						// 이미 처리 중인 메시지는 건너뜀
						if _, isProcessing := processingMessages.LoadOrStore(ev.ClientMsgID, true); isProcessing {
							continue
						}
						defer processingMessages.Delete(ev.ClientMsgID)

						channelID := ev.Channel
						message := ev.Text

						responseMessage, _ := messageFunction(message)
						_, _, _, err := api.SendMessage(channelID, slack.MsgOptionText(responseMessage, false))
						if err != nil {
							fmt.Printf("응답 메시지를 보내는데 실패했습니다: %v\n", err)
						}
					}
				}
			}
		}
	}()

	for {
		err := client.Run()
		if err != nil {
			log.Printf("소켓 모드 클라이언트 실행 중 에러 발생: %v", err)
		}
	}
}
