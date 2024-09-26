package sensitive_records

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

func (h *SensitiveRecordHandlers) PostSensitiveRecord(c *gin.Context, _ openapi.PostSensitiveRecordParams) {
	sensitiveRecordTypeID, err := sensitive_records.NewSensitiveRecordTypeID(c.Query("type"))
	if err != nil {
		c.JSON(http.StatusBadRequest, openapi.ErrorResponse{Errors: err.Error()})
		return
	}

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
	sensitiveRecord, err := sensitive_records.CreateSensitiveRecord(userID, int(sensitiveRecordTypeID), req.Metadata)
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
		Id:                    sensitiveRecord.Id(),
		Metadata:              sensitiveRecord.Metadata(),
		SensitiveRecordTypeId: sensitiveRecord.TypeID(),
	})
}

func (h *SensitiveRecordHandlers) CreateSensitiveRecordDataWithID(c *gin.Context, id int) {
	//TODO implement me
	panic("implement me")
}
