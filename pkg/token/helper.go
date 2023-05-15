package token

import (
	"fmt"
	"github.com/lee-lou2/go/pkg/redis"
)

// SetTokenInCache 토큰 저장
func SetTokenInCache(userId int, tokenType string, sessionId string) error {
	key := fmt.Sprintf("USER_TOKEN__%s__%d", tokenType, userId)
	client, _ := redis.Client()
	if err := client.PushValueSetLimit(key, sessionId, userTokenSetLimit); err != nil {
		return err
	}
	return nil
}

// ExistsTokenInCache 토큰 존재 여부 확인
func ExistsTokenInCache(userId int, tokenType string, sessionId string) error {
	key := fmt.Sprintf("USER_TOKEN__%s__%d", tokenType, userId)
	client, _ := redis.Client()
	exists, err := client.ExistsValue(key, sessionId)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("token not found")
	}
	return nil
}
