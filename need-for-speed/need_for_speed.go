package speed

// Car type struct with fields
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain, battery: 100}
}

// Track type struct
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	var drivenDistance int
	if car.battery < car.batteryDrain {
		drivenDistance = car.distance
	} else {
		car.battery -= car.batteryDrain
		drivenDistance += (car.distance + car.speed)
	}
	return Car{battery: car.battery, batteryDrain: car.batteryDrain, distance: drivenDistance, speed: car.speed}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	for i := car.battery; i > 0; i -= car.batteryDrain {
		car.distance += car.speed
	}
	return car.distance > track.distance
}
