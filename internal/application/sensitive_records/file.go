package sensitive_records

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/SpaceSlow/gophkeeper/pkg/crypto"
)

func (h *SensitiveRecordHandlers) UploadFile(c *gin.Context) {
	userID, err := crypto.UserID(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	uuid, err := h.repo.CreateFile(userID, c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, openapi.CreateFileResponse{Hash: uuid.String()}) // TODO fix response needed uuid field
}

func (h *SensitiveRecordHandlers) DownloadFile(c *gin.Context, hash string) {
	//TODO implement me
	panic("implement me")
}
