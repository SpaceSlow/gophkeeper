package sensitive_records

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"

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

func (h *SensitiveRecordHandlers) DownloadFile(c *gin.Context, uuid openapi_types.UUID) {
	userID, err := crypto.UserID(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	file, err := h.repo.FetchFile(userID, uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, openapi.ErrorResponse{Errors: errors.New("not found file with current uuid").Error()})
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	if _, err := io.Copy(c.Writer, file); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
