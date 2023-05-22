package wrtn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Generate AI 생성
func Generate(prompt, accessToken, refreshToken string) (*string, error) {
	if refreshToken == "" {
		refreshToken = os.Getenv("WRTN_DEFAULT_REFRESH_TOKEN")
	}
	if refreshToken != "" {
		token, err := RefreshToken(refreshToken)
		if err != nil {
			return nil, err
		}
		accessToken = *token
	}
	if accessToken == "" {
		return nil, fmt.Errorf("access token is empty")
	}

	user := os.Getenv("WRTN_USER")
	room := os.Getenv("WRTN_ROOM")
	url := fmt.Sprintf("https://william.wow.wrtn.ai/generate/stream/%s?type=big&model=GPT4&platform=web&user=%s", room, user)
	method := "POST"
	payload := []byte(fmt.Sprintf(`{"message":"%s","reroll":false}`, prompt))

	// 요청 객체 생성
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// 헤더 추가
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("authorization", "Bearer "+accessToken)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("referrer", "https://wrtn.ai/")
	req.Header.Add("referrerPolicy", "strict-origin-when-cross-origin")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")

	// 요청 보내고 응답 받기
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	responseBody := buf.String()

	// 데이터를 라인 별로 분리하여 배열에 저장합니다.
	lines := strings.Split(responseBody, "\n")

	result := ""
	for _, line := range lines {
		// 불필요한 앞뒤 공백을 삭제합니다.
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// ":"를 중심으로 라인을 두 부분으로 분리합니다.
		parts := strings.SplitN(line, ":", 2)
		dataStr := strings.TrimSpace(parts[1])

		// data 값을 맵 형태로 변환합니다.
		data := make(map[string]interface{})
		err := json.Unmarshal([]byte(dataStr), &data)
		if err != nil {
			continue
		}

		if data["chunk"] == nil {
			continue
		}
		result += data["chunk"].(string)
	}
	return &result, nil
}
