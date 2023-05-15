package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
)

// CreateRSAKeyPair 키 생성
func CreateRSAKeyPair() ([]byte, []byte, error) {
	// 키 생성
	privateKey, publicKey, err := GenerateKeyPair(2048)
	if err != nil {
		return nil, nil, err
	}
	// 키 변환
	privateKeyBytes := PrivateKeyToBytes(privateKey)
	publicKeyBytes, err := PublicKeyToBytes(publicKey)
	if err != nil {
		return nil, nil, err
	}
	return privateKeyBytes, publicKeyBytes, nil
}

// GenerateKeyPair 공개/비밀키 생성
func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// PrivateKeyToBytes 비밀키 타입(key>bytes) 변환
func PrivateKeyToBytes(privateKey *rsa.PrivateKey) []byte {
	privateBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)

	return privateBytes
}

// PublicKeyToBytes 공개키 타입(key>bytes) 변환
func PublicKeyToBytes(pub *rsa.PublicKey) ([]byte, error) {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, err
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes, nil
}

// BytesToPrivateKey 비밀키 타입(bytes>key) 변환
func BytesToPrivateKey(privateKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKey)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// BytesToPublicKey 공개키 타입(bytes>key) 변환
func BytesToPublicKey(pub []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return nil, err
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		panic(err)
	}
	return key, nil
}
