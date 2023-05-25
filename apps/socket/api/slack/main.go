package slack

import (
	"github.com/lee-lou2/go/apps/ai/models"
	"github.com/lee-lou2/go/platform/wrtn"
)

// RequestMessage 메시지 가져오기
func RequestMessage(text string) (string, error) {
	response, err := wrtn.Generate(text, "", "", "")
	if err != nil {
		return "요청 실패", err
	}
	request := models.Dataset{
		Instruction: text,
		Output:      *response,
		Status:      2,
	}
	_, err = request.Create()
	if err != nil {
		return "요청 실패", err
	}
	return *response, nil
}
