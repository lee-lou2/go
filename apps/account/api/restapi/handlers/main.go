package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/apps/account/models"
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
	userId, err := oauth2.GetUserInfo(token.AccessToken)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 사용자 생성
	user, err := models.GetOrCreateUser(&models.User{ServerID: userId})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// JWT 토큰 생성
	tokenPair, err := user.GenerateToken()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 쿠키에 토큰 저장
	serverHost := os.Getenv("SERVER_HOST")
	c.SetCookie("access_token", tokenPair.AccessToken.Token, 3600, "/", serverHost, false, true)
	c.SetCookie("refresh_token", tokenPair.RefreshToken.Token, 3600, "/", serverHost, false, true)
	c.SetCookie("expires_in", strconv.Itoa(int(tokenPair.AccessToken.ExpiredAt)), 3600, "/", serverHost, false, true)
	c.Redirect(302, "/")
}
