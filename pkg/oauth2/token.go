package oauth2

import (
	"encoding/json"
	"fmt"
	"github.com/lee-lou2/go/pkg/requests"
	"net/url"
	"os"
	"strings"
)

// tokenResponse 토큰 응답
type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
}

// GetToken 토큰 가져오기
func GetToken(code, state string) (*tokenResponse, error) {
	var respData tokenResponse

	data := url.Values{}
	data.Set("client_id", os.Getenv("OAUTH2_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("OAUTH2_CLIENT_SECRET"))
	data.Set("code", code)
	data.Set("code_verifier", state)
	data.Set("redirect_uri", os.Getenv("OAUTH2_CALLBACK_URL"))
	data.Set("grant_type", "authorization_code")
	payload := strings.NewReader(data.Encode())

	// 요청
	resp, _ := requests.Http(
		"POST",
		os.Getenv("OAUTH2_SSO_SERVER_HOST")+"/oauth2/token/",
		payload,
		&requests.Header{
			Key:   "Cache-Control",
			Value: "no-cache",
		},
		&requests.Header{
			Key:   "Content-Type",
			Value: "application/x-www-form-urlencoded",
		},
	)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("토큰 조회를 실패하였습니다 : %s", resp.Body)
	}
	if err := json.Unmarshal([]byte(resp.Body), &respData); err != nil {
		return nil, err
	}
	return &respData, nil
}
