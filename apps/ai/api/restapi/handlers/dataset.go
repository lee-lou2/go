package handlers

import "github.com/gin-gonic/gin"

// RequestBody 데이터셋 요청 바디
type RequestBody struct {
	Instruction string   `json:"instruction"`
	Input       string   `json:"input"`
	Output      string   `json:"output"`
	Status      int      `json:"status"`
	Category    []string `json:"category"`
}

func GetDatasetsHandler(c *gin.Context) {
}

func GetInstructionHandler(c *gin.Context) {
}

func CreateDatasetHandler(c *gin.Context) {
	request := RequestBody{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": request})
}

func UpdateDatasetHandler(c *gin.Context) {
}
