package strategy

import (
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type binaryFileRequest struct {
	Metadata string                `form:"metadata"`
	Data     *multipart.FileHeader `form:"data"`
}

type BinaryFileStrategy struct{}

func (s *BinaryFileStrategy) Upload(c *gin.Context) (string, error) {
	var binaryFileReq binaryFileRequest
	if err := c.Bind(&binaryFileReq); err != nil {
		return "", err
	}
	return fmt.Sprintf("Binary file"), nil
}
