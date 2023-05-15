package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/pkg/oauth2"
	"os"
	"strconv"
)

// LoginHandler 로그인
func LoginHandler(c *gin.Context) {
	resp, _ := oauth2.GetRedirectUri()
	c.Redirect(302, resp)
}

// CallbackHandler 콜백
func CallbackHandler(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	token, err := oauth2.GetToken(code, state)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 쿠키에 토큰 저장
	serverHost := os.Getenv("SERVER_HOST")
	c.SetCookie("access_token", token.AccessToken, 3600, "/", serverHost, false, true)
	c.SetCookie("refresh_token", token.RefreshToken, 3600, "/", serverHost, false, true)
	c.SetCookie("id_token", token.IdToken, 3600, "/", serverHost, false, true)
	c.SetCookie("expires_in", strconv.Itoa(token.ExpiresIn), 3600, "/", serverHost, false, true)
	c.Redirect(302, "/")
}
