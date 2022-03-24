package elon

import (
	"fmt"
)

// Drive method on the 'Car' updates the number of meters driven based on the car's speed,
// and reduces the battery according to the battery drainage.
// If a car's battery is below its battery drain percentage, car can't be driven anymore.
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// DisplayDistance method on 'Car' returns the distance as displayed on the LED display as a string.
// Output should be as in example: "Driven 0 meters"
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery method on 'Car' returns the battery percentage as displayed on the LED display as a string.
// Output should be as in example: "Battery at 100%"
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

/*
CanFinish method on 'Car' takes a trackDistance int as its parameter.
To finish a race, a car has to be able to drive the race's distance.
This means not draining its battery before having crossed the finish line.
Returns true if the car can finish the race; otherwise, return false. */
func (car *Car) CanFinish(trackDistance int) bool {
	// for car.battery > 0 {
	// 	car.Drive()
	// }
	// return car.distance >= trackDistance
	maxDistance := (car.battery / car.batteryDrain) * car.speed
	return maxDistance >= trackDistance
}
