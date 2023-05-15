package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// aesValidate 데이터 유효성 검사
func aesValidate(data string) error {
	if data == "" {
		return fmt.Errorf("data is empty")
	}
	return nil
}

// AESCipherEncrypt AES 암호화
func AESCipherEncrypt(data string, configs ...Configs) (string, error) {
	var key []byte
	cfg, err := SetConfigs(configs...)
	if err != nil {
		return "", err
	}
	if err := aesValidate(data); err != nil {
		return "", err
	}
	key = []byte(cfg.AESCipherKey)
	plainText := []byte(data)

	// AES 대칭키 암호화 블록 생성
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 암호화
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	encText := base64.URLEncoding.EncodeToString(cipherText)
	return encText, nil
}

// AESCipherDecrypt AES 복호화
func AESCipherDecrypt(encText string, configs ...Configs) (string, error) {
	var key []byte
	cfg, err := SetConfigs(configs...)
	if err != nil {
		return "", err
	}
	if err := aesValidate(encText); err != nil {
		return "", err
	}
	key = []byte(cfg.AESCipherKey)
	cipherText, err := base64.URLEncoding.DecodeString(encText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("cipherText too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
