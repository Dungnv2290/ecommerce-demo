package main

import (
	"log"
	"net/http"

	"github.com/dungnguyen/ecommerce-demo/CartService/pkg/api"
	"github.com/dungnguyen/ecommerce-demo/CartService/pkg/data"
	"github.com/gorilla/mux"
)

func main() {
	repository := new(data.CartRepository)
	repository.InitRepository("root:root@tcp(localhost:3306)/cartdb")
	apiHandler := new(api.CartAPI)
	apiHandler.Repository = repository
	router := mux.NewRouter()
	router.HandleFunc("/cart", apiHandler.AddCartHandler).Methods("POST")
	router.HandleFunc("/cart/{userID}", apiHandler.GetCartHandler).Methods("GET")
	router.HandleFunc("/cart/{userID}", apiHandler.EmptyCartHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8889", router))
}
