package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/platform/wrtn"
	"net/http"
)

// WRTNRequestBody 요청 바디
type WRTNRequestBody struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Model        string `json:"model"`
	Text         string `json:"text"`
}

// WRTNHandler 뤼튼 챗봇
func WRTNHandler(c *gin.Context) {
	var request WRTNRequestBody
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := wrtn.Generate(request.Text, request.AccessToken, request.RefreshToken, request.Model)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}
