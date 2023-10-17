package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dungnguyen/ecommerce-demo/CartService/pkg/data"
	"github.com/dungnguyen/ecommerce-demo/CartService/pkg/model"
	"github.com/gorilla/mux"
)

type CartAPI struct {
	Repository *data.CartRepository
}

func (c *CartAPI) AddCartHandler(res http.ResponseWriter, req *http.Request) {
	cart := new(model.Cart)
	jsonErr := json.NewDecoder(req.Body).Decode(&cart)
	if jsonErr != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, jsonErr)
		return
	}
	defer req.Body.Close()

	for _, item := range cart.Items {
		c.Repository.AddItemToCart(cart.UserID, item)
	}
	res.WriteHeader(http.StatusAccepted)
}

func (c *CartAPI) GetCartHandler(res http.ResponseWriter, req *http.Request) {
	args := mux.Vars(req)
	userID := args["userID"]
	cart := c.Repository.GetCart(userID)

	if len(cart.Items) != 0 {
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(cart)
	} else {
		res.WriteHeader(http.StatusNotFound)
	}
}

func (c *CartAPI) EmptyCartHandler(res http.ResponseWriter, req *http.Request) {
	args := mux.Vars(req)
	userID := args["userID"]
	c.Repository.EmptyCart(userID)
	res.WriteHeader(http.StatusAccepted)
}
