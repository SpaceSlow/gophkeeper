package sensitive_records

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginPasswordStrategy struct{}

type LoginPasswordData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *LoginPasswordStrategy) Upload(req *http.Request) (string, error) {
	var credentials LoginPasswordData
	if err := json.NewDecoder(req.Body).Decode(&credentials); err != nil {
		return "", err
	}
	return fmt.Sprintf("Login credentials received: Username=%s, Password=%s", credentials.Username, credentials.Password), nil
}
