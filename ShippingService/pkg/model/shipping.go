package model

import "github.com/dungnguyen/ecommerce-demo/CartService/pkg/model"

type Address struct {
	StreetAddress string `json:"streetAddress"`
	State         string `json:"state"`
	City          string `json:"city"`
	Country       string `json:"country"`
	ZipCode       int    `json:"zipCode"`
}

type ShippingQuoteRequest struct {
	Address Address    `json:"address"`
	Cart    model.Cart `json:"cart"`
}

type ShippingQuoteResponse struct {
	Cost float64 `json:"cost"`
}

type ShippingOrderRequest struct {
	Address Address    `json:"address"`
	Cart    model.Cart `json:"cart"`
}

type ShippingOrderResponse struct {
	TrackingID string `json:"trackingID"`
}
