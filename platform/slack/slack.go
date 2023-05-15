package slack

import (
	"github.com/slack-go/slack"
)

// SendSlack 슬랙 전송
func SendSlack(channel, msg string, configs ...Configs) error {
	cfg, err := setConfigs(configs...)
	if err != nil {
		return err
	}
	api := slack.New(cfg.BotToken)
	if _, _, err := api.PostMessage(
		channel,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionAsUser(true),
	); err != nil {
		return err
	}
	return nil
}
