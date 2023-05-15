package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/apps/notify/cmd/consumers"
	"github.com/lee-lou2/go/apps/notify/cmd/schemas"
)

// SendMessageHandler 메세지 전송
func SendMessageHandler(c *gin.Context) {
	app := c.Param("app")
	var request interface{}
	switch app {
	case "push":
		request = schemas.RequestBody[schemas.PushRequest]{}
	case "email":
		request = schemas.RequestBody[schemas.EmailRequest]{}
	case "sms":
		request = schemas.RequestBody[schemas.SMSRequest]{}
	case "slack":
		request = schemas.RequestBody[schemas.SlackRequest]{}
	default:
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid app"})
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	go func() { consumers.NotifyQue <- &request }()
	c.JSON(200, gin.H{"message": "success"})
}
