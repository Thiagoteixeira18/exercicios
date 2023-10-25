package exercicios

// TODO: define the 'Car' type struct

// Define a structure Car
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// Define a function NewCar that returns an instance of Car
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// Define a structure Track
type Track struct {
	distance int
}

// Define a function NewTrack that returns an instance of Track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Define a function Drive that updates the distance and decreases the battery
func Drive(c Car) Car {
	if c.battery >= c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
	return c
}

// Define a function CanFinish that checks if the car can finish the track
func CanFinish(c Car, t Track) bool {
	maxDistance := c.battery / c.batteryDrain
	return maxDistance*c.speed >= t.distance
}
