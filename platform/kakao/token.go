package kakao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lee-lou2/go/pkg/redis"
	"github.com/lee-lou2/go/pkg/requests"
	"github.com/lee-lou2/go/pkg/utils"
	netUrl "net/url"
	"os"
)

// setKaKaOToken 카카오톡 토큰 설정
func setKaKaOToken(resp string) error {
	tokenValue, err := utils.StringToMap(resp)
	if err != nil {
		return err
	}

	if _, hasError := tokenValue["error_code"]; hasError {
		// 오류
		err := fmt.Errorf("토큰 조회를 실패하였습니다 : %s", tokenValue["error_code"].(string))
		return err
	}
	if _, hasAccessKey := tokenValue["access_token"]; hasAccessKey {
		cache, err := redis.Client()
		if err != nil {
			return err
		}
		if _, hasRefreshKey := tokenValue["refresh_token"]; hasRefreshKey {
			// 리프레시 토큰 저장
			if err := cache.SetValue(
				"kakao_refresh_token",
				tokenValue["refresh_token"].(string),
				int(tokenValue["refresh_token_expires_in"].(float64)),
			); err != nil {
				return err
			}
		}
		// 액세스 토큰 저장
		if err := cache.SetValue(
			"kakao_access_token",
			tokenValue["access_token"].(string),
			int(tokenValue["expires_in"].(float64)),
		); err != nil {
			return err
		}
	} else {
		err := fmt.Errorf("토큰 값이 포함되어있지 않습니다")
		return err
	}
	return nil
}

// CreateKaKaOToken 카카오 토큰 생성
func CreateKaKaOToken() error {
	clientId := os.Getenv("KAKAO_API_CLIENT_ID")
	redirectUri := os.Getenv("KAKAO_API_REDIRECT_URI")
	code := os.Getenv("KAKAO_API_CODE")

	url := "https://kauth.kakao.com/oauth/token"

	params := netUrl.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", clientId)
	params.Add("redirect_uri", redirectUri)
	params.Add("code", code)

	resp, err := requests.Http(
		"POST",
		url,
		bytes.NewBufferString(params.Encode()),
		&requests.Header{Key: "Content-Type", Value: "application/x-www-form-urlencoded"},
	)
	if err != nil {
		return err
	}
	if err := setKaKaOToken(resp.Body); err != nil {
		return err
	}
	return nil
}

// RefreshKaKaoToken 카카오 토큰 재발급
func RefreshKaKaoToken() error {
	clientId := os.Getenv("KAKAO_API_CLIENT_ID")
	cache, err := redis.Client()
	if err != nil {
		return err
	}
	refreshToken, err := cache.GetValue("kakao_refresh_token")
	if err != nil {
		return err
	}

	url := "https://kauth.kakao.com/oauth/token"

	params := netUrl.Values{}
	params.Add("grant_type", "refresh_token")
	params.Add("client_id", clientId)
	params.Add("refresh_token", refreshToken)

	resp, err := requests.Http(
		"POST",
		url,
		bytes.NewBufferString(params.Encode()),
	)
	if err != nil {
		return err
	}
	if err := setKaKaOToken(resp.Body); err != nil {
		return err
	}
	return nil
}

// SimpleMessageTemplate 기본 메시지 템플릿
func SimpleMessageTemplate(message string) string {
	templateObject := map[string]interface{}{
		"object_type": "text",
		"text":        message,
		"link": map[string]string{
			"web_url":        "",
			"mobile_web_url": "",
		},
	}
	templateObjectBytes, _ := json.Marshal(templateObject)
	return string(templateObjectBytes)
}

// SendKaKaOToMe 나에게 카카오톡 보내기
func SendKaKaOToMe(message string) error {
	templateObject := SimpleMessageTemplate(message)

	url := "https://kapi.kakao.com/v2/api/talk/memo/default/send"

	cache, err := redis.Client()
	if err != nil {
		return err
	}
	accessToken, err := cache.GetValue("kakao_access_token")
	if err != nil {
		return err
	}

	params := netUrl.Values{}
	params.Add("template_object", templateObject)

	_, err = requests.Http(
		"POST",
		url,
		bytes.NewBufferString(params.Encode()),
		&requests.Header{Key: "Content-Type", Value: "application/x-www-form-urlencoded"},
		&requests.Header{Key: "Authorization", Value: "Bearer " + accessToken},
	)
	if err != nil {
		return err
	}
	return nil
}

// SendKaKaOToFriend 친구에게 카카오톡 보내기
func SendKaKaOToFriend(message, friendUuid string) error {
	templateObject := SimpleMessageTemplate(message)

	url := "https://kapi.kakao.com/v1/api/talk/friends/message/default/send"

	cache, err := redis.Client()
	if err != nil {
		return err
	}
	accessToken, err := cache.GetValue("kakao_access_token")
	if err != nil {
		return err
	}

	params := netUrl.Values{}
	params.Add("receiver_uuids", fmt.Sprintf(`["%s"]`, friendUuid))
	params.Add("template_object", templateObject)

	resp, err := requests.Http(
		"POST",
		url,
		bytes.NewBufferString(params.Encode()),
		&requests.Header{Key: "Content-Type", Value: "application/x-www-form-urlencoded"},
		&requests.Header{Key: "Authorization", Value: "Bearer " + accessToken},
	)
	fmt.Println(resp)
	if err != nil {
		return err
	}
	return nil
}
