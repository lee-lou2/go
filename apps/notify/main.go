package notify

import "github.com/lee-lou2/go/apps/notify/cmd/consumers"

func Run() {
	go func() { _ = consumers.NotifyConsumer() }()
}
