package strategy

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type credentialRequest struct {
	Metadata string     `json:"metadata"`
	Data     Credential `json:"data"`
}

type CredentialStrategy struct{}

func (s *CredentialStrategy) Upload(c *gin.Context) (string, error) {
	var credentialReq credentialRequest
	if err := c.BindJSON(&credentialReq); err != nil {
		return "", err
	}
	return fmt.Sprintf("Login credentials received: Username=%s, Password=%s", credentialReq.Data.Username, credentialReq.Data.Password), nil
}