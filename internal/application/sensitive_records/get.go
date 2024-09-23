package sensitive_records

import (
	"net/http"

	"github.com/SpaceSlow/gophkeeper/generated/openapi"
	"github.com/gin-gonic/gin"
)

func (h *SensitiveRecordHandlers) ListSensitiveRecordTypes(c *gin.Context) {
	result, err := h.repo.ListSensitiveRecordTypes()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	types := make([]openapi.SensitiveRecordType, 0, len(result))
	for _, t := range result {
		types = append(types, openapi.SensitiveRecordType{Id: t.Id(), Name: openapi.SensitiveRecordTypeEnum(t.Name())})
	}

	c.JSON(http.StatusOK, openapi.ListSensitiveRecordTypeResponse{SensitiveRecordTypes: types})
}

func (h *SensitiveRecordHandlers) ListSensitiveRecords(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *SensitiveRecordHandlers) SensitiveRecordDataWithID(c *gin.Context, id int) {
	//TODO implement me
	panic("implement me")
}
