package sensitive_records

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func (h *SensitiveRecordHandlers) setStrategy(recordType string) error {
	switch recordType {
	case "text-file":
		h.strategy = &TextFileStrategy{}
	case "binary-file":
		h.strategy = &BinaryFileStrategy{}
	case "credentials":
		h.strategy = &CredentialStrategy{}
	case "payment-card":
		h.strategy = &PaymentCardStrategy{}
	default:
		return fmt.Errorf("unknown sensitive record type: %s", recordType)
	}
	return nil
}
