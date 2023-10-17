package main

import (
	"log"
	"net/http"

	"github.com/dungnguyen/ecommerce-demo/CheckoutService/pkg/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	checkoutEndpoint := api.CheckoutEndpoint{}
	router.HandleFunc("/checkout", checkoutEndpoint.Checkout).Methods("POST")
	log.Fatal(http.ListenAndServe(":8812", router))
}
