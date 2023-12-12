package exercicios

import (
	"fmt"
)

// NeedsLicense determines whether a license is needed to operate a certain type of vehicle.
// Only "car" and "truck" require a license; all others do not.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle between two options based on lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	}
	return option2 + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var resellPrice float64

	if age < 3 {
		resellPrice = 0.8 * originalPrice
	} else if age >= 10 {
		resellPrice = 0.5 * originalPrice
	} else {
		resellPrice = 0.7 * originalPrice
	}
	return resellPrice
}

func main2() {
	// Examples of using the functions

	needLicense := NeedsLicense("car")
	fmt.Println("Needs License for car:", needLicense) // Should print true

	needLicense = NeedsLicense("bike")
	fmt.Println("Needs License for bike:", needLicense) // Should print false

	needLicense = NeedsLicense("truck")
	fmt.Println("Needs License for truck:", needLicense) // Should print true

	vehicle := ChooseVehicle("Wuling Hongguang", "Toyota Corolla")
	fmt.Println(vehicle) // Should print "Toyota Corolla is clearly the better choice."

	vehicle = ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf")
	fmt.Println(vehicle) // Should print "Volkswagen Beetle is clearly the better choice."

	price := CalculateResellPrice(1000, 1)
	fmt.Println("Resell Price for 1-year-old vehicle:", price) // Should print 800

	price = CalculateResellPrice(1000, 5)
	fmt.Println("Resell Price for 5-year-old vehicle:", price) // Should print 700

	price = CalculateResellPrice(1000, 15)
	fmt.Println("Resell Price for 15-year-old vehicle:", price) // Should print 500
}
