package restapi

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// DefaultMiddleware Fiber 기본 미들웨어
func DefaultMiddleware(r *gin.Engine) {
	// gin 기본 미들웨어
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// gin static
	r.Static("/static", "./static")
	// cors
	r.Use(cors.Default())
}
