package exercicios

import "testing"

func TestCalculateResellPrice(t *testing.T) {
	originalPrice := 100.0
	age := 2.0

	expectedResellPrice := 0.8 * originalPrice
	resellPrice := CalculateResellPrice(originalPrice, age)

	if resellPrice != expectedResellPrice {
		t.Errorf("CalculateResellPrice(%f, %f) = %f; want %f", originalPrice, age, resellPrice, expectedResellPrice)
	}
}