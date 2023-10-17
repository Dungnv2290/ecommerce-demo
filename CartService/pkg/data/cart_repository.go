package data

import (
	"database/sql"
	"log"

	"github.com/dungnguyen/ecommerce-demo/CartService/pkg/model"
	_ "github.com/go-sql-driver/mysql"
)

// CartRepository exposes API for cart data
type CartRepository struct {
	Db *sql.DB
}

func (c *CartRepository) InitRepository(connectionStr string) {
	db, dbConnectionErr := sql.Open("mysql", connectionStr)
	if dbConnectionErr != nil {
		log.Fatal(dbConnectionErr)
	}
	c.Db = db
}

func (c *CartRepository) AddItemToCart(userID string, item model.Item) {
	rows, queryErr := c.Db.Query("insert into cart (userID,productID,quantity) values(?,?,?)", userID, item.ProductID, item.QUantity)
	if queryErr != nil {
		log.Fatal(queryErr)
	}
	defer rows.Close()
}

func (c *CartRepository) GetCart(userID string) model.Cart {
	log.Println("user in repository", userID)
	cart := model.Cart{}
	cart.UserID = userID
	rows, queryErr := c.Db.Query("select productID,quantity from cart where userID=?", userID)
	if queryErr != nil {
		log.Fatal(queryErr)
	}
	items := make([]model.Item, 0)
	for rows.Next() {
		item := model.Item{}
		var productID string
		var quantity int
		if err := rows.Scan(&productID, &quantity); err != nil {
			log.Fatal(err)
		}
		item.ProductID = productID
		item.QUantity = quantity
		log.Println("Item in repository ", item)
		items = append(items, item)
	}
	defer rows.Close()
	cart.Items = items
	log.Println("Cart in repository ", cart)
	return cart
}

func (c *CartRepository) AddOrUpdateCartItem(userID string, item model.Item) {
	log.Println("User in repository ", userID)
	var quantity int
	rows, queryErr := c.Db.Query("select quantity from cart where userID=? productID=?", userID, item.ProductID)
	if queryErr != nil {
		log.Fatal(queryErr)
	}
	if rows.Next() {
		if err := rows.Scan(&quantity); err != nil {
			log.Fatal(err)
		}
	} else {
		c.AddItemToCart(userID, item)
		return
	}

	item.QUantity = item.QUantity + quantity
	rows, updateQueryErr := c.Db.Query("update cart set quantity=? where userID=? productID=?", item.QUantity, userID, item.ProductID)
	if updateQueryErr != nil {
		log.Fatal(updateQueryErr)
	}
	defer rows.Close()
}

func (c *CartRepository) EmptyCart(userID string) {
	rows, queryErr := c.Db.Query("delete from cart where userID=?", userID)
	if queryErr != nil {
		log.Fatal(queryErr)
	}
	defer rows.Close()
}
