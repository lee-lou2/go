package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/apps/notify/cmd/schemas"
	"github.com/lee-lou2/go/apps/notify/cmd/senders"
)

// SendMessageHandler 메세지 전송
func SendMessageHandler(c *gin.Context) {
	app := c.Param("app")
	switch app {
	case "push":
		request := schemas.RequestBody[schemas.PushRequest]{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		go func() { senders.SendPush(&request) }()
	case "email":
		request := schemas.RequestBody[schemas.EmailRequest]{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		go func() { senders.SendEmail(&request) }()
	case "sms":
		request := schemas.RequestBody[schemas.SMSRequest]{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		go func() { senders.SendSMS(&request) }()
	case "slack":
		request := schemas.RequestBody[schemas.SlackRequest]{}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			return
		}
		go func() { senders.SendSlack(&request) }()
	default:
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid app"})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
