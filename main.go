package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Cart struct {
	CartItems []Item `json:"items"`
}

type Item struct {
	Prod string `json:"prod"`
	Qty  int    `json:"qty"`
}

type ItemCost struct {
	prod  string
	price float32
}

var itemsCost map[string]float32

func main() {
	inputCart := `{"items":[{"prod":"apple","qty":6},{"prod":"orange","qty":3},{"prod":"banana","qty":4}]}`
	cart, error := GetCart(inputCart)
	if error == nil {
		fmt.Println("Task1=", Checkout(cart))
	} else {
		fmt.Println((error))
	}
}

func Checkout(cart Cart) float32 {
	itemsCost = getPrices()
	var result float32
	for _, s := range cart.CartItems {
		if price, found := itemsCost[s.Prod]; found {
			result = result + (price * float32(s.Qty))
		} else {
			fmt.Println(s.Prod, " not found")
		}
	}

	return result
}

//Get a Cart object
func GetCart(inputCart string) (Cart, error) {
	var cart Cart
	error := json.Unmarshal([]byte(inputCart), &cart)
	if error != nil {
		return cart, errors.New("Invalid Cart")
	}

	if cart.CartItems == nil {
		return cart, errors.New("Invalid Cart")
	}

	return cart, nil
}

//Return the current product and prices
func getPrices() map[string]float32 {
	itemsCost := make(map[string]float32)
	itemsCost["apple"] = 0.25
	itemsCost["orange"] = 0.05
	itemsCost["banana"] = 0.30
	itemsCost["mango"] = 1.25
	return itemsCost
}
