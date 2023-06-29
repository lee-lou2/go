package handlers

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

// VideoStreamingHandler 비디오 스트리밍
func VideoStreamingHandler(c *gin.Context) {
	// 파일명 확인 (sample.mp4)
	fileName := c.Param("filename")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "filename is required",
		})
		return
	}
	// 아래 파일이 존재하는지 확인
	filePath := "static/contents/" + fileName
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file not found",
		})
		return
	}
	f, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer f.Close()

	fi, _ := f.Stat()
	c.Header("Content-Type", "video/mp4")
	c.Header("Content-Length", string(fi.Size()))

	if _, err := io.Copy(c.Writer, f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func HLSStreamingHandler(c *gin.Context) {
	fileName := c.Param("filename")
	c.Header("Content-Type", "application/vnd.apple.mpegurl")
	http.ServeFile(c.Writer, c.Request, "./static/contents/"+fileName)
}
