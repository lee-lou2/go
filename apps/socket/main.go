package socket

import (
	slackApi "github.com/lee-lou2/go/apps/socket/api/slack"
	"github.com/lee-lou2/go/platform/slack"
)

// Run Socket 실행
func Run() {
	go func() { _ = slack.SocketModeSlackBot(slackApi.RequestMessage) }()
}
