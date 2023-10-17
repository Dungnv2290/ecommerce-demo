package model

import "fmt"

type CreditCard struct {
	Number          string `json:"number"`
	CSV             int    `json:"csv"`
	ExpirationYear  int    `json:"expirationYear"`
	ExpirationMonth int    `json:"expirationMonth"`
}

func (c CreditCard) String() string {
	creditCardStr := "number:%s,csv: %d,expirationMonth: %d,expirationYear:%d"
	return fmt.Sprintf(creditCardStr, c.Number, c.CSV, c.ExpirationMonth, c.ExpirationYear)
}

type PaymentRequest struct {
	CreditCardInfo CreditCard `json:"creditCard"`
	Amount         float64    `json:"amount"`
}

type PaymentResponse struct {
	TransactionID string
}

func (p PaymentRequest) String() string {
	paymentRequestStr := "creditCard:{%s},amount:%f"
	return fmt.Sprintf(paymentRequestStr, p.CreditCardInfo, p.Amount)
}
