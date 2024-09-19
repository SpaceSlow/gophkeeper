package sensitive_records

import (
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type textFileRequest struct {
	Preview  string                `form:"preview"`
	Metadata string                `form:"metadata"`
	Data     *multipart.FileHeader `form:"data"`
}

type TextFileStrategy struct{}

func (s *TextFileStrategy) Upload(c *gin.Context) (string, error) {
	var textFileReq textFileRequest
	if err := c.Bind(&textFileReq); err != nil {
		return "", err
	}
	return fmt.Sprintf("Text file"), nil
}
