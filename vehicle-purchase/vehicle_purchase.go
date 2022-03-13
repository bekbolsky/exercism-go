package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	var recommendedOption string

	if option2 < option1 {
		recommendedOption = option2
	} else {
		recommendedOption = option1
	}
	return fmt.Sprintf("%s is clearly the better choice.", recommendedOption)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var calculatedPrice float64
	const (
		discountLessThan3     float64 = 0.8
		discountBetween3and10 float64 = 0.7
		discountMoreThan10    float64 = 0.5
	)

	if age < 3 {
		calculatedPrice = originalPrice * discountLessThan3
	} else if age >= 3 && age < 10 {
		calculatedPrice = originalPrice * discountBetween3and10
	} else if age >= 10 {
		calculatedPrice = originalPrice * discountMoreThan10
	}
	return calculatedPrice
}
