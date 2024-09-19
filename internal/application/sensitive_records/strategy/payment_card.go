package strategy

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type PaymentCard struct {
	Number     string `json:"number"`
	ExpireDate string `json:"expire_date"`
	Cardholder string `json:"cardholder"`
	Code       string `json:"code"`
}

type paymentCardRequest struct {
	Preview  string      `json:"preview"`
	Metadata string      `json:"metadata"`
	Data     PaymentCard `json:"data"`
}

type PaymentCardStrategy struct{}

func (s *PaymentCardStrategy) Upload(c *gin.Context) (string, error) {
	var paymentCardReq paymentCardRequest
	if err := c.BindJSON(&paymentCardReq); err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"Payment card credentials received: Number=%s, Expire date=%s, Cardholder=%s",
		paymentCardReq.Data.Number,
		paymentCardReq.Data.ExpireDate,
		paymentCardReq.Data.Cardholder,
	), nil
}
