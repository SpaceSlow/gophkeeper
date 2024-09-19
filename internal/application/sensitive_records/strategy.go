package sensitive_records

import (
	"github.com/gin-gonic/gin"
)

type SensitiveRecordStrategy interface {
	Upload(c *gin.Context) (string, error)
}
