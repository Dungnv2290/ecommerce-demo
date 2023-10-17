package api

import (
	"encoding/json"
	"net/http"
	"strings"

	product "github.com/dungnguyen/ecommerce-demo/ProductCatalogService/pkg/model"
	"github.com/gorilla/mux"
)

// ProductAPI is data structure hold the static data
type ProductAPI struct {
	CatelogMap map[string]product.Product
}

func (p *ProductAPI) GetProductHandler(res http.ResponseWriter, req *http.Request) {
	args := mux.Vars(req)
	productID := args["id"]
	product, isExist := p.CatelogMap[productID]
	if isExist {
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(product)
	} else {
		res.WriteHeader(404)
	}
}

func (p *ProductAPI) GetAllProductsHandler(res http.ResponseWriter, req *http.Request) {
	var products []product.Product
	for _, value := range p.CatelogMap {
		products = append(products, value)
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(products)
}

func (p *ProductAPI) SearchProductHandler(res http.ResponseWriter, req *http.Request) {
	queryString := req.URL.Query().Get("query")
	var matchingProducts []product.Product

	for id, prod := range p.CatelogMap {
		if strings.Contains(prod.Name, queryString) {
			matchingProducts = append(matchingProducts, p.CatelogMap[id])
		}
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(matchingProducts)
}
