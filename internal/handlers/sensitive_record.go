package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/internal/handlers/sensitive_records"
	"github.com/SpaceSlow/gophkeeper/internal/store"
)

type SensitiveRecordHandler struct {
	db       *store.DB
	strategy sensitive_records.SensitiveRecordStrategy
}

func NewSensitiveRecordHandler(db *store.DB) *SensitiveRecordHandler {
	return &SensitiveRecordHandler{db: db}
}

func (h *SensitiveRecordHandler) Upload() func(c *gin.Context) {
	return func(c *gin.Context) {
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
}

func (h *SensitiveRecordHandler) setStrategy(recordType string) error {
	switch recordType {
	case "text-file":
		h.strategy = &sensitive_records.TextFileStrategy{}
	case "binary-file":
		h.strategy = &sensitive_records.BinaryFileStrategy{}
	case "credentials":
		h.strategy = &sensitive_records.CredentialStrategy{}
	case "payment-card":
		h.strategy = &sensitive_records.PaymentCardStrategy{}
	default:
		return fmt.Errorf("unknown sensitive record type: %s", recordType)
	}
	return nil
}
