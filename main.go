package main

import (
	"fmt"
)

type Cart struct {
	CartItems []Item
}

type Item struct {
	prod string
	qty  int
}

type ItemCost struct {
	prod  string
	price float32
}

var itemsCost map[string]float32

func main() {

	itemsCost = SetPrices()

	items := make([]Item, 3)
	items = append(items, Item{"apple", 6}, Item{"orange", 3}, Item{"banana", 4})

	var cart Cart
	cart.CartItems = items

	fmt.Println("Task1=", checkout(cart))
}

func checkout(cart Cart) float32 {
	var result float32

	for _, s := range cart.CartItems {
		result = result + (itemsCost[s.prod] * float32(s.qty))
	}

	return result
}

func SetPrices() map[string]float32 {
	itemsCost := make(map[string]float32)
	itemsCost["apple"] = 0.23
	itemsCost["orange"] = 1.25
	itemsCost["banana"] = 2.36
	itemsCost["mango"] = 0.25

	return itemsCost
}
