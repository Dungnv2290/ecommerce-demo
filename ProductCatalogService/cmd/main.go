package main

import (
	"log"
	"net/http"

	"github.com/dungnguyen/ecommerce-demo/ProductCatalogService/pkg/api"
	"github.com/dungnguyen/ecommerce-demo/ProductCatalogService/pkg/loader"
	"github.com/gorilla/mux"
)

func main() {
	productAPI := &api.ProductAPI{}
	productAPI.CatelogMap = loader.LoadProductCatelog("products.json")
	router := mux.NewRouter()

	router.HandleFunc("/product/{id}", productAPI.GetProductHandler).Methods("GET")
	router.HandleFunc("/product", productAPI.GetAllProductsHandler).Methods("GET")
	router.HandleFunc("/search", productAPI.SearchProductHandler).Methods("GET").Queries("query", "{query}")

	log.Fatal(http.ListenAndServe(":8888", router))
}
