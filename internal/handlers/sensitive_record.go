package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/internal/handlers/sensitive_records"
	"github.com/SpaceSlow/gophkeeper/internal/store"
)

func SensitiveRecordUploadHandler(db *store.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		recordType := c.DefaultQuery("type", "credentials")

		strategy, err := sensitive_records.NewSensitiveRecordStrategy(recordType)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := strategy.Upload(c.Request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": result})
	}
}
