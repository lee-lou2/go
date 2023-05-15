package naver

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// DetectLanguage 언어 감지
func DetectLanguage(text string, configs ...Configs) (string, error) {
	cfg, err := setConfigs(configs...)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://openapi.naver.com/v1/papago/detectLangs", bytes.NewBufferString("query="+text))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-Naver-Client-Id", cfg.NaverClientID)
	req.Header.Add("X-Naver-Client-Secret", cfg.NaverClientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]string
	json.Unmarshal(body, &result)

	return result["langCode"], nil
}

// Translate 텍스트 번역
func Translate(text, insertLanguage string, exportLanguage string, configs ...Configs) (string, error) {
	cfg, err := setConfigs(configs...)
	if err != nil {
		return "", err
	}

	detectLanguage := insertLanguage
	if detectLanguage == "" {
		detectLanguage, err = DetectLanguage(text, *cfg)
		if err != nil {
			return "", err
		}
	}

	// API 엔드포인트 및 파라미터 설정
	apiURL := "https://openapi.naver.com/v1/papago/n2mt"
	parameters := url.Values{}
	parameters.Add("source", detectLanguage)
	parameters.Add("target", exportLanguage)
	parameters.Add("text", text)

	// HTTP 요청 생성
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(parameters.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Naver-Client-Id", cfg.NaverClientID)
	req.Header.Set("X-Naver-Client-Secret", cfg.NaverClientSecret)

	// HTTP 요청 실행
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// 응답 결과 처리
	var result map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	translatedText := result["message"].(map[string]interface{})["result"].(map[string]interface{})["translatedText"].(string)
	return translatedText, nil
}
