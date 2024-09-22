package strategy

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type textRequest struct {
	Metadata string `json:"metadata"`
	Data     string `json:"data"`
}

type TextStrategy struct{}

func (s *TextStrategy) Upload(c *gin.Context) (string, error) {
	var textReq textRequest
	if err := c.BindJSON(&textReq); err != nil {
		return "", err
	}
	return fmt.Sprintf("Text: %s", textReq.Data), nil
}
