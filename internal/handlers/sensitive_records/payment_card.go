package sensitive_records

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PaymentCardStrategy struct{}

type PaymentCardData struct {
	Number     string `json:"number"`
	ExpireDate string `json:"expire_date"`
	Cardholder string `json:"cardholder"`
	Code       string `json:"code"`
}

func (s *PaymentCardStrategy) Upload(req *http.Request) (string, error) {
	var credentials PaymentCardData
	if err := json.NewDecoder(req.Body).Decode(&credentials); err != nil {
		return "", err
	}
	return fmt.Sprintf("Payment card credentials received: Number=%s, Expire date=%s, Cardholder=%s", credentials.Number, credentials.ExpireDate, credentials.Cardholder), nil
}
