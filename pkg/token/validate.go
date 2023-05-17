package token

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

// AdditionalValues 선택적 검증 데이터
type AdditionalValues struct {
	TokenType string
}

// VerifyToken 토큰 검증
func VerifyToken(token, sessionId string, additional ...AdditionalValues) (*Payload, error) {
	// 1. 토큰 유효성 검사
	claims, err := verifyTokenBindClaims(token)
	if err != nil {
		return nil, err
	}

	// 2. 세션 정보 일치 여부 확인
	if err := verifySession(claims, sessionId); err != nil {
		return nil, err
	}

	// 3. 캐시 존재 여부 확인
	if claims["user_id"] == nil {
		return nil, fmt.Errorf("token missing required info")
	}
	userId := claims["user_id"].(int)
	tokenType := claims["type"].(string)
	if err := ExistsTokenInCache(userId, tokenType, sessionId); err != nil {
		return nil, err
	}

	// 4. 추가 검증
	if len(additional) == 1 {
		if additional[0].TokenType != "" && tokenType != additional[0].TokenType {
			return nil, fmt.Errorf("invalid token type")
		}
	} else {
		if tokenType != AccessTokenString {
			return nil, fmt.Errorf("invalid token type")
		}
	}
	return &Payload{
		UserID:    userId,
		SessionId: sessionId,
		Scope:     claims["scope"].(string),
	}, nil
}

// verifyToken 토큰 검증 후 정보 반환
func verifyTokenBindClaims(tokenValue string, configs ...Configs) (map[string]interface{}, error) {
	cfg, err := setConfigs(configs...)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token validation error")
		}
		return []byte(cfg.JWTSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token validation error")
	}
	return token.Claims.(jwt.MapClaims), nil
}

// verifySession 토큰 내 세션 확인
func verifySession(claims map[string]interface{}, sessionId string) error {
	if claims["session_id"].(string) != sessionId {
		return fmt.Errorf("invalid session id")
	}
	return nil
}
