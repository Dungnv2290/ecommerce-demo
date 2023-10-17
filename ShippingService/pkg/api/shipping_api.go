package api

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/dungnguyen/ecommerce-demo/ShippingService/pkg/model"
)

type ShippingEndpoint struct {
}

func (s *ShippingEndpoint) GetShippingQuote(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	shippingQuoteRequest := new(model.ShippingQuoteRequest)
	decodeErr := json.NewDecoder(req.Body).Decode(shippingQuoteRequest)
	if decodeErr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, decodeErr)
		return
	}

	var cost float64
	count := len(shippingQuoteRequest.Cart.Items)
	if count != 0 {
		cost = math.Pow(3, (1 + 0.2*float64(count)))
	}

	shippingQuoteResponse := new(model.ShippingQuoteResponse)
	shippingQuoteResponse.Cost = cost
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(shippingQuoteResponse)
}

func (s *ShippingEndpoint) ProcessShippingOrder(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	shippingOrderRequest := new(model.ShippingOrderRequest)
	decodeErr := json.NewDecoder(req.Body).Decode(shippingOrderRequest)
	if decodeErr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, decodeErr)
		return
	}

	shippingOrderResponse := new(model.ShippingOrderResponse)
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGen := rand.New(randomSource)
	transactionID := randomGen.Intn(1000000)
	log.Println("TransactionID :", transactionID)
	shippingOrderResponse.TrackingID = fmt.Sprint(transactionID)

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(shippingOrderResponse)
}
