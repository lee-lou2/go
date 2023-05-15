package twilio

import (
	"fmt"
	"github.com/kevinburke/twilio-go"
	"strings"
)

func ValidateSMSTwilioRequest(to string) error {
	// 010 으로 시작하는지 확인
	if !strings.HasPrefix(to, "010") {
		return fmt.Errorf("invalid phone number")
	}
	// 10자리 또는 11자리인지 확인
	if len(to) != 10 && len(to) != 11 {
		return fmt.Errorf("invalid phone number")
	}
	// 문자가 있는지 확인
	if len(to) == 0 {
		return fmt.Errorf("invalid message")
	}
	return nil
}

// SendSMSTwilio 트윌리오를 이용한 문자 발송
func SendSMSTwilio(to, message string, configs ...Configs) error {
	cfg, err := setConfigs(configs...)
	if err != nil {
		return err
	}
	to = strings.Replace(to, "-", "", -1)
	to = strings.Replace(to, " ", "", -1)
	if ValidateSMSTwilioRequest(to) != nil {
		return fmt.Errorf("invalid request")
	}
	client := twilio.NewClient(cfg.AccountSid, cfg.AuthToken, nil)

	if _, err := client.Messages.SendMessage(
		cfg.FromNumber,
		fmt.Sprintf("+82%s", to[1:]),
		message,
		nil,
	); err != nil {
		return err
	}
	return nil
}
