package oauth2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 테스트를 위한 길이 범위 정의
const (
	minLength = 43
	maxLength = 128
)

// getCodeAtGivenLength 함수 테스트
func TestGetCodeAtGivenLength(t *testing.T) {
	length := 10
	code, err := getCodeAtGivenLength(length)

	assert.NoError(t, err)
	assert.Len(t, code, length)
}

// getCodeAtRandomLength 함수 테스트
func TestGetCodeAtRandomLength(t *testing.T) {
	code, err := getCodeAtRandomLength()

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(code), minLength)
	assert.LessOrEqual(t, len(code), maxLength)
}

// getBase64UrlEncoded 함수 테스트
func TestGetBase64UrlEncoded(t *testing.T) {
	input := []byte("example input")
	encoded := getBase64UrlEncoded(input)

	assert.NotEmpty(t, encoded)
	assert.NotContains(t, encoded, "+")
	assert.NotContains(t, encoded, "/")
}

// getCodeVerifier 함수 테스트
func TestGetCodeVerifier(t *testing.T) {
	codeVerifier, err := getCodeVerifier()

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(codeVerifier), minLength)
}

// getCodeChallenge 함수 테스트
func TestGetCodeChallenge(t *testing.T) {
	codeVerifier := "example code verifier"
	codeChallenge := getCodeChallenge(codeVerifier)

	assert.NotEmpty(t, codeChallenge)
}
