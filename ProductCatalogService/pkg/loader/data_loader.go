package loader

import (
	"encoding/json"
	"io/ioutil"
	"log"

	product "github.com/dungnguyen/ecommerce-demo/ProductCatalogService/pkg/model"
)

type catelog struct {
	Products []product.Product `json:"products"`
}

// LoadProductCatelog loads the product catelog from json file
func LoadProductCatelog(fileLocation string) map[string]product.Product {
	productCatelogMap := make(map[string]product.Product)
	catelogFile, fileErr := ioutil.ReadFile(fileLocation)
	if fileErr != nil {
		log.Fatalln("File err", fileErr)
	}
	var productCatelog catelog
	jsonErr := json.Unmarshal(catelogFile, &productCatelog)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	for _, prod := range productCatelog.Products {
		productCatelogMap[prod.ID] = prod
	}

	return productCatelogMap
}
