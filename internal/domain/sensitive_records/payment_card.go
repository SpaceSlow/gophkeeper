package sensitive_records

type PaymentCard struct {
	Number     string
	ExpMonth   uint8
	ExpYear    uint8
	Cardholder string
	Code       int16
}
