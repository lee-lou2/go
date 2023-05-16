package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lee-lou2/go/apps/ai/models"
	"strconv"
)

// GetDatasetsHandler 데이터셋 목록 조회
func GetDatasetsHandler(c *gin.Context) {
	datasets, err := models.GetDatasets()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": datasets})
}

// GetInstructionHandler 데이터셋 1개 조회
func GetInstructionHandler(c *gin.Context) {
	dataset, err := models.GetInstruction()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 데이터가 빈 값이면 404 Not Found
	if dataset.ID == 0 {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, gin.H{"data": dataset})
}

// CreateDatasetHandler 데이터셋 생성
func CreateDatasetHandler(c *gin.Context) {
	request := models.Dataset{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	obj, err := request.Save()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": obj})
}

// UpdateDatasetHandler 데이터셋 수정
func UpdateDatasetHandler(c *gin.Context) {
	request := models.Dataset{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	obj, err := request.Update(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": obj})
}
