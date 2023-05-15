package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/apps/notify/api/restapi/handlers"
)

func V1Routes(v1 *gin.RouterGroup) {
	notify := v1.Group("/notify")
	{
		notify.POST(":app", handlers.SendMessageHandler)
	}
}
