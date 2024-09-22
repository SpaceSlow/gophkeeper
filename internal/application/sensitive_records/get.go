package sensitive_records

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type sensitiveRecordType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (h *SensitiveRecordHandlers) ListSensitiveRecordTypes(c *gin.Context) {
	result, err := h.repo.ListSensitiveRecordTypes()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	types := make([]sensitiveRecordType, 0, len(result))
	for _, t := range result {
		types = append(types, sensitiveRecordType{Id: t.Id(), Name: t.Name()})
	}

	c.JSON(http.StatusOK, types)
}

func (h *SensitiveRecordHandlers) ListSensitiveRecords(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *SensitiveRecordHandlers) SensitiveRecordDataWithID(c *gin.Context, id int) {
	//TODO implement me
	panic("implement me")
}
