package restapi

import "github.com/gin-gonic/gin"

// DefaultMiddleware Fiber 기본 미들웨어
func DefaultMiddleware(r *gin.Engine) {
	// gin 기본 미들웨어
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// gin static
	r.Static("/static", "./static")
	// cors
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
}
