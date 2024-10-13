package sensitive_records

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/internal/domain/sensitive_records"
	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

func (h *SensitiveRecordHandlers) DeleteSensitiveRecordWithID(c *gin.Context, id int) {
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

	err = h.repo.DeleteSensitiveRecord(id)
	var notExistErr *sensitive_records.NotExistSensitiveRecordDataError
	if errors.As(err, &notExistErr) {
		c.JSON(http.StatusNotFound, openapi.ErrorResponse{Errors: err.Error()})
		return
	} else if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}
