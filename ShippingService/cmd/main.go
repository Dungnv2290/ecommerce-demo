package main

import (
	"log"
	"net/http"

	"github.com/dungnguyen/ecommerce-demo/ShippingService/pkg/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	ShippingEndpoint := api.ShippingEndpoint{}
	router.HandleFunc("/shipping/getquote", ShippingEndpoint.GetShippingQuote).Methods("POST")
	router.HandleFunc("/shipping/order", ShippingEndpoint.ProcessShippingOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(":8811", router))
}
