package sensitive_records

import (
	"errors"
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
	userID, err := crypto.UserID(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	isOwner, err := h.repo.IsSensitiveRecordOwner(id, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if !isOwner {
		c.JSON(http.StatusForbidden, openapi.ErrorResponse{Errors: "no rights to the sensitive record"})
		return
	}

	data, err := c.Request.GetBody()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer data.Close()

	recordData, err := sensitive_records.NewSensitiveRecordData(id, data)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = h.repo.CreateSensitiveRecordData(recordData)
	var existErr *sensitive_records.ExistSensitiveRecordDataError
	if errors.As(err, &existErr) {
		c.JSON(http.StatusConflict, openapi.ErrorResponse{Errors: err.Error()})
		return
	} else if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}
