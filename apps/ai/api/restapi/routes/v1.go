package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/apps/ai/api/restapi/handlers"
)

// V1Routes AI 라우터
func V1Routes(v1 *gin.RouterGroup) {
	ai := v1.Group("/ai")
	{
		// 데이터셋
		ai.GET("/dataset", handlers.GetDatasetsHandler)
		ai.GET("/dataset/instruction", handlers.GetInstructionHandler)
		ai.POST("/dataset", handlers.CreateDatasetHandler)
		ai.PATCH("/dataset/:id", handlers.UpdateDatasetHandler)
	}
}
