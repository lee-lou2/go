package models

import (
	"fmt"
	"github.com/lee-lou2/go/configs/db"
	"github.com/lee-lou2/go/pkg/token"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	ServerID int    `json:"sso_id"`
}

// GetOrCreateUser 사용자 생성 또는 조회
func GetOrCreateUser(user *User) (*User, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	// ServerID 가 0 이 아닌지 확인
	if user.ServerID == 0 {
		return nil, fmt.Errorf("ServerID is required")
	}
	// ServerID 로 사용자 조회
	if err := conn.Where("server_id = ?", user.ServerID).First(&user).Error; err != nil {
		// 조회 실패 시 새로 생성
		if err := conn.Create(&user).Error; err != nil {
			return nil, err
		}
	}
	return user, nil
}

// GenerateToken 토큰 생성
func (user *User) GenerateToken() (*token.TokenPair, error) {
	tokenPair, err := token.GetToken(token.Payload{
		UserID:    int(user.ID),
		SessionId: "",
		Scope:     "default",
	})
	if err != nil {
		return nil, err
	}
	return tokenPair, nil
}
