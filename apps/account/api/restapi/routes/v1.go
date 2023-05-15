package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/apps/account/api/restapi/handlers"
)

// V1Routes 계정 라우터
func V1Routes(v1 *gin.RouterGroup) {
	accounts := v1.Group("/accounts")
	{
		// 로그인
		accounts.GET("/login", handlers.LoginHandler)
		// 콜백
		accounts.GET("/callback", handlers.CallbackHandler)
	}
}
