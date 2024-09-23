package sensitive_records

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *SensitiveRecordHandlers) CreateSensitiveRecord(c *gin.Context, params interface{}) {
	//TODO implement me
	panic("implement me")
}

func (h *SensitiveRecordHandlers) CreateSensitiveRecordDataWithID(c *gin.Context, id int) {
	//TODO implement me
	panic("implement me")
}

func (h *SensitiveRecordHandlers) UploadSensitiveRecord(c *gin.Context) {
	recordType := c.Query("type")

	err := h.setStrategy(recordType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.strategy.Upload(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result})
}