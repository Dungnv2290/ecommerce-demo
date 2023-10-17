package model

import (
	"github.com/dungnguyen/ecommerce-demo/CartService/pkg/model"
	payment "github.com/dungnguyen/ecommerce-demo/PaymentService/pkg/model"
	shipping "github.com/dungnguyen/ecommerce-demo/ShippingService/pkg/model"
)

type Order struct {
	UserID         string             `json:"userID"`
	Email          string             `json:"email"`
	Address        shipping.Address   `json:"address"`
	CreditCardInfo payment.CreditCard `json:"creditCard"`
}

type OrderResult struct {
	OrderID    string           `json:"orderID"`
	TrackingID string           `json:"trackingID"`
	Address    shipping.Address `json:"address"`
	Cart       model.Cart       `json:"cart"`
	Cost       float64          `json:"cost"`
}
