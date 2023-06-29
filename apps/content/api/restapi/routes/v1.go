package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/apps/content/api/restapi/handlers"
)

// V1Routes 컨텐츠 라우터
func V1Routes(v1 *gin.RouterGroup) {
	content := v1.Group("/content")
	{
		// 비디오 스트리밍
		content.GET("/video/stream/html5/:filename", handlers.VideoStreamingHandler)
		content.GET("/video/stream/hls/:filename", handlers.HLSStreamingHandler)
	}
}
