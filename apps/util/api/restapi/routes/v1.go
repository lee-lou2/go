package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/apps/util/api/restapi/handlers"
)

// V1Routes 유틸 라우터
func V1Routes(v1 *gin.RouterGroup) {
	utils := v1.Group("/utils")
	{
		// 번역
		utils.GET("/translate", handlers.TranslateHandler)
	}
}
