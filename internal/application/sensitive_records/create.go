package sensitive_records

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

func (h *SensitiveRecordHandlers) PostSensitiveRecord(c *gin.Context) {
	var req openapi.CreateSensitiveRecordRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, openapi.ErrorResponse{Errors: err.Error()})
		return
	}

	userID, err := crypto.UserID(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	sensitiveRecord, err := sensitive_records.CreateSensitiveRecord(userID, string(req.Type), req.Metadata)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	sensitiveRecord, err = h.repo.CreateSensitiveRecord(sensitiveRecord)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, openapi.CreateSensitiveRecordResponse{
		Id:       sensitiveRecord.Id(),
		Metadata: sensitiveRecord.Metadata(),
		Type:     openapi.SensitiveRecordTypeEnum(sensitiveRecord.Type()),
	})
}

func (h *SensitiveRecordHandlers) PostSensitiveRecordData(c *gin.Context, id int) {
	//TODO implement me
	panic("implement me")
}
