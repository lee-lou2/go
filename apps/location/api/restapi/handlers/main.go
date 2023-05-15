package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/pkg/location"
)

// LocationHandler 지역 정보 조회
func LocationHandler(c *gin.Context) {
	ipAddress := c.Query("ip")
	// IP 값이 없으면 현재 사용자의 IP 조회
	ipAddress = c.ClientIP()
	if ipAddress == "" {
		c.JSON(409, gin.H{"message": "아이피를 찾을 수 없습니다"})
	}
	_location, _ := location.GetLocation(ipAddress)
	c.JSON(200, gin.H{"location": _location})
}
