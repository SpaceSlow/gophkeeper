package sensitive_records

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

func (h *SensitiveRecordHandlers) ListSensitiveRecords(c *gin.Context) {
	userID, err := crypto.UserID(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	sensitiveRecords, err := h.repo.ListSensitiveRecords(userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	responseRecords := make([]openapi.SensitiveRecord, 0, len(sensitiveRecords))
	for _, s := range sensitiveRecords {
		responseRecords = append(responseRecords, openapi.SensitiveRecord{
			Id:       s.Id(),
			Metadata: s.Metadata(),
			Type:     openapi.SensitiveRecordTypeEnum(s.Type()),
		})
	}

	c.JSON(http.StatusOK, openapi.ListSensitiveRecordResponse{SensitiveRecords: responseRecords})
}

func (h *SensitiveRecordHandlers) FetchSensitiveRecordWithID(c *gin.Context, id int) {
	//TODO implement me
	panic("implement me")
}
