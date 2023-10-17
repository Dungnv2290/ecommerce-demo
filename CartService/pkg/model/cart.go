package model

// Item holds selected product and the quantity
type Item struct {
	ProductID string `json:"productID"`
	QUantity  int    `json:"quantity"`
}

// Cart holds purchased products and userID
type Cart struct {
	UserID string `json:"userID"`
	Items  []Item `json:"items"`
}
