package sensitive_records

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

func (h *SensitiveRecordHandlers) UploadFile(c *gin.Context) {
	userID, err := crypto.UserID(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if c.GetHeader("Content-Type") != "application/octet-stream" || c.Request.Body == http.NoBody {
		c.JSON(http.StatusBadRequest, openapi.ErrorResponse{Errors: errors.New("invalid data").Error()})
		return
	}

	uuid, err := h.repo.CreateFile(userID, c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, openapi.CreateFileResponse{Uuid: uuid})
}

func (h *SensitiveRecordHandlers) DownloadFile(c *gin.Context) {
	userID, err := crypto.UserID(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	uid, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, openapi.ErrorResponse{Errors: errors.New("invalid uuid").Error()})
		return
	}

	file, err := h.repo.FetchFile(userID, uid)
	if err != nil {
		c.JSON(http.StatusNotFound, openapi.ErrorResponse{Errors: errors.New("not found file with current uuid").Error()})
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	if _, err := io.Copy(c.Writer, file); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
