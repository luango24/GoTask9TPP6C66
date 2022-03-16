package main

import "testing"

func TestCheckout(t *testing.T) {
	input := `{"items":[{"prod":"apple","qty":6},{"prod":"orange","qty":3},{"prod":"banana","qty":4}]}`
	cart, _ := GetCart(input)
	total := Checkout(cart)
	if total != 2.85 {
		t.Errorf("Checkout was incorrect, got: %f, want: %f.", total, 2.85)
	}
}
