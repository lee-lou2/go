package slack

import (
	"github.com/lee-lou2/go/apps/ai/models"
)

// RequestMessage 메시지 가져오기
func RequestMessage(text string) (string, error) {
	request := models.Dataset{
		Instruction: text,
		Category:    "slack",
	}
	_, err := request.Create()
	if err != nil {
		return "요청 실패", err
	}
	return "요청 완료", nil
}
