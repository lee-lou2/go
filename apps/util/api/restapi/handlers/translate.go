package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/platform/naver"
)

// TranslateHandler 번역
func TranslateHandler(c *gin.Context) {
	text := c.Query("text")
	exportLanguage := c.Query("lang")
	if exportLanguage == "" {
		exportLanguage = "en"
	}
	resp, err := naver.Translate(text, "", exportLanguage)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": resp})
	}
}
