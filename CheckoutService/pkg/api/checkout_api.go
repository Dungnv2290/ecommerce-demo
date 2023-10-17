package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	cart "github.com/dungnguyen/ecommerce-demo/CartService/pkg/model"
	"github.com/dungnguyen/ecommerce-demo/CheckoutService/pkg/model"
	payment "github.com/dungnguyen/ecommerce-demo/PaymentService/pkg/model"
	product "github.com/dungnguyen/ecommerce-demo/ProductCatalogService/pkg/model"
	shipping "github.com/dungnguyen/ecommerce-demo/ShippingService/pkg/model"
)

type CheckoutEndpoint struct {
}

func (c *CheckoutEndpoint) Checkout(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	defer req.Body.Close()

	checkoutRequest := new(model.Order)
	unmarshalErr := json.Unmarshal(body, &checkoutRequest)

	if unmarshalErr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, unmarshalErr)
		return
	}
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGen := rand.New(randomSource)
	orderID := randomGen.Intn(1000000)
	orderResult := model.OrderResult{}
	orderResult.OrderID = fmt.Sprint(orderID)
	//Get the cart
	cart, cartErr := getCart(checkoutRequest.UserID)
	if cartErr != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, cartErr)
		return
	}
	//Get total cost estimation
	cost, costCalErr := calculateCost(cart)
	if costCalErr != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, costCalErr)
		return
	}
	// Get shipping cost
	shippingCost, shippingCostCalErr := getShippingCostEstimate(checkoutRequest.Address, cart)
	if shippingCostCalErr != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, shippingCostCalErr)
		return
	}
	cost = +shippingCost
	//Charge cc
	_, paymentErr := charge(checkoutRequest.CreditCardInfo, cost)
	if paymentErr != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, paymentErr)
		return
	}
	//Place shipping order
	trackingID, shippingOrderError := placeOrder(checkoutRequest.Address, cart)
	if shippingOrderError != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, shippingOrderError)
		return
	}
	//Empty cart
	emptyCartErr := emptyCart(checkoutRequest.UserID)

	if emptyCartErr != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, emptyCartErr)
		return
	}
	//send the result
	orderResult.Address = checkoutRequest.Address
	orderResult.Cart = cart
	orderResult.Cost = float64(cost)
	orderResult.TrackingID = trackingID
	respBytes, respMarshalErr := json.Marshal(orderResult)

	if respMarshalErr != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, respMarshalErr)
		return
	}
	res.Write(respBytes)
	res.Header().Set("Content-Type", "application/json")

}

func getCart(userID string) (cart.Cart, error) {
	cart := cart.Cart{}
	resp, cartErr := http.Get("http://localhost:8889/cart/" + userID)

	// handle 500
	if cartErr != nil {
		return cart, cartErr
	}

	defer resp.Body.Close()
	decodeErr := json.NewDecoder(resp.Body).Decode(&cart)
	if decodeErr != nil {
		return cart, decodeErr
	}

	return cart, nil
}

func calculateCost(cart cart.Cart) (float32, error) {
	var cost float32
	for _, item := range cart.Items {
		resp, catelogErr := http.Get("http://localhost:8888/product/" + item.ProductID)
		if catelogErr != nil {
			return cost, catelogErr
		}

		defer resp.Body.Close()
		product := product.Product{}
		decodeErr := json.NewDecoder(resp.Body).Decode(&product)
		if decodeErr != nil {
			return cost, decodeErr
		}

		cost = cost + product.Price*float32(item.QUantity)
	}

	return cost, nil
}

func getShippingCostEstimate(adddress shipping.Address, cart cart.Cart) (float32, error) {
	var shippingCost float32
	shippingQuoteRequest := shipping.ShippingQuoteRequest{}
	shippingQuoteRequest.Address = adddress
	shippingQuoteRequest.Cart = cart
	bodyBytes, marshalErr := json.Marshal(shippingQuoteRequest)
	if marshalErr != nil {
		return shippingCost, marshalErr
	}
	reader := bytes.NewReader(bodyBytes)
	resp, nwErr := http.Post("http://localhost:8811/shipping/getquote", "application/json", reader)
	if nwErr != nil {
		return shippingCost, nwErr
	}

	shippingQuoteResponse := shipping.ShippingQuoteResponse{}
	decodeErr := json.NewDecoder(resp.Body).Decode(&shippingQuoteResponse)
	if decodeErr != nil {
		return shippingCost, decodeErr
	}

	shippingCost = float32(shippingQuoteResponse.Cost)
	return shippingCost, nil
}

func charge(creditCard payment.CreditCard, amount float32) (string, error) {
	var transactionID string
	paymentRequest := payment.PaymentRequest{}
	paymentRequest.CreditCardInfo = creditCard
	paymentRequest.Amount = float64(amount)
	bodyBytes, marshalErr := json.Marshal(paymentRequest)
	if marshalErr != nil {
		return transactionID, marshalErr
	}
	reader := bytes.NewReader(bodyBytes)
	resp, nwErr := http.Post("http://localhost:8810/payment", "application/json", reader)
	if nwErr != nil {
		return transactionID, nwErr
	}

	paymentResponse := payment.PaymentResponse{}
	defer resp.Body.Close()
	decodeErr := json.NewDecoder(resp.Body).Decode(&paymentResponse)
	if decodeErr != nil {
		return transactionID, decodeErr
	}
	transactionID = paymentResponse.TransactionID
	return transactionID, nil
}

func placeOrder(address shipping.Address, cart cart.Cart) (string, error) {
	var trackingID string
	shippingOrderRequest := shipping.ShippingOrderRequest{}
	shippingOrderRequest.Cart = cart
	shippingOrderRequest.Address = address
	bodyBytes, marshalErr := json.Marshal(shippingOrderRequest)
	if marshalErr != nil {
		return trackingID, marshalErr
	}
	reader := bytes.NewReader(bodyBytes)
	resp, nwErr := http.Post("http://localhost:8811/shipping/order", "application/json", reader)
	if nwErr != nil {
		return trackingID, nwErr
	}

	shippingOrderResponse := shipping.ShippingOrderResponse{}
	defer resp.Body.Close()
	decodeErr := json.NewDecoder(resp.Body).Decode(&shippingOrderResponse)
	if decodeErr != nil {
		return trackingID, decodeErr
	}
	trackingID = shippingOrderResponse.TrackingID
	return trackingID, nil
}

func emptyCart(userID string) error {
	req, rqErr := http.NewRequest(http.MethodDelete, "http://localhost:8889/cart/"+userID, nil)
	if rqErr != nil {
		return rqErr
	}
	resp, respErr := http.DefaultClient.Do(req)
	if respErr != nil {
		return respErr
	}
	if resp.StatusCode != http.StatusAccepted {
		return errors.New("Error in emptying cart")
	}
	return nil
}
