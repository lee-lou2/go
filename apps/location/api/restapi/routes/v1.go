package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/apps/location/api/restapi/handlers"
)

func V1Routes(v1 *gin.RouterGroup) {
	location := v1.Group("/location")
	{
		// ip 를 이용해 지역 정보 조회
		location.GET("", handlers.LocationHandler)
	}
}
