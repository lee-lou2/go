package wrtn

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Response 응답
type Response struct {
	Result string `json:"result"`
	Data   struct {
		AccessToken string `json:"accessToken"`
	} `json:"data"`
}

// RefreshToken 토큰 갱신
func RefreshToken(token string) (*string, error) {
	url := "https://api.wow.wrtn.ai/auth/refresh"

	data := strings.NewReader("")
	req, err := http.NewRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Refresh", token)
	req.Header.Set("Sec-CH-UA", `"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
	req.Header.Set("Sec-CH-UA-Mobile", "?0")
	req.Header.Set("Sec-CH-UA-Platform", "macOS")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	accessToken := response.Data.AccessToken
	return &accessToken, nil
}
