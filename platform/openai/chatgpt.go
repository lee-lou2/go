package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lee-lou2/go/pkg/requests"
	"os"
)

// Response 응답
type Response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// messages 메세지
var messages []map[string]string

// GetCompletion 챗봇 생성
func getCompletion(message map[string]string) string {
	var content Response

	// 메세지 유지
	messages = append(messages, message)

	body := map[string]interface{}{
		"model":       "gpt-3.5-turbo",
		"messages":    messages,
		"temperature": 0.5,
		"max_tokens":  1000,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("JSON 인코딩 에러: %s\n", err)
		return ""
	}

	// 요청
	resp, _ := requests.Http(
		"POST",
		"https://api.openai.com/v1/chat/completions",
		bytes.NewBuffer(jsonBody),
		&requests.Header{
			Key:   "Content-Type",
			Value: "application/json",
		},
		&requests.Header{
			Key:   "Authorization",
			Value: "Bearer " + os.Getenv("OPENAI_API_KEY"),
		},
	)
	if err := json.Unmarshal([]byte(resp.Body), &content); err != nil {
		fmt.Printf("JSON 디코딩 에러: %s\n", err)
		return ""
	}
	messages = append(messages, map[string]string{
		"role":    "assistant",
		"content": content.Choices[0].Message.Content,
	})
	return content.Choices[0].Message.Content
}
