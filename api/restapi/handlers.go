package restapi

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// HealthCheckHandler 헬스 체크
func HealthCheckHandler(c *gin.Context) {
	// gin "pong" response
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// TestSleepHandler 테스트용 대기
func TestSleepHandler(c *gin.Context) {
	// 지정된 초 만큼 대기
	s := c.Query("s")
	if s == "" {
		s = "0"
	}
	intValue, _ := strconv.Atoi(s)
	time.Sleep(time.Duration(intValue) * time.Second)
	c.JSON(200, gin.H{
		"message": "wake up!",
	})
}
