package oauth2

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
	"strings"
)

// getCodeAtGivenLength 주어진 길이의 코드를 반환
func getCodeAtGivenLength(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[index.Int64()]
	}
	return string(result), nil
}

// getCodeAtRandomLength 랜덤한 길이의 코드를 반환
func getCodeAtRandomLength() (string, error) {
	length, err := rand.Int(rand.Reader, big.NewInt(86))
	if err != nil {
		return "", err
	}
	return getCodeAtGivenLength(43 + int(length.Int64()))
}

// getBase64UrlEncoded base64 url encoding
func getBase64UrlEncoded(input []byte) string {
	result := base64.URLEncoding.EncodeToString(input)
	return strings.TrimRight(result, "=")
}

// getCodeVerifier 코드 베리파이어를 반환
func getCodeVerifier() (string, error) {
	codeVerifier, err := getCodeAtRandomLength()
	if err != nil {
		return "", err
	}
	return getBase64UrlEncoded([]byte(codeVerifier)), nil
}

// getCodeChallenge 코드 챌린지를 반환
func getCodeChallenge(codeVerifier string) string {
	hash := sha256.Sum256([]byte(codeVerifier))
	return getBase64UrlEncoded(hash[:])
}
