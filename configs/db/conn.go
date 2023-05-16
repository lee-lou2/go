package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

// GetDB - 데이터베이스 연결
func GetDB() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}
	var err error
	// DB 연결
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
