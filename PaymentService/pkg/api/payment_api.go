package api

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/dungnguyen/ecommerce-demo/PaymentService/pkg/model"
)

type PaymentEndpoint struct {
}

func (p *PaymentEndpoint) Charge(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	paymentRequest := new(model.PaymentRequest)
	decodeErr := json.NewDecoder(req.Body).Decode(paymentRequest)
	if decodeErr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, decodeErr)
		return
	}

	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGen := rand.New(randomSource)
	transactionID := randomGen.Intn(1000000)
	log.Printf("Processing payment: transactionID: %d, %s", transactionID, *paymentRequest)
	paymentResponse := model.PaymentResponse{}
	paymentResponse.TransactionID = fmt.Sprint(transactionID)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(paymentResponse)
}
