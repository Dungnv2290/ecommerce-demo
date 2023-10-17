package main

import (
	"log"
	"net/http"

	"github.com/dungnguyen/ecommerce-demo/PaymentService/pkg/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	paymentEndpoint := api.PaymentEndpoint{}
	router.HandleFunc("/payment", paymentEndpoint.Charge).Methods("POST")
	log.Fatal(http.ListenAndServe(":8810", router))
}
