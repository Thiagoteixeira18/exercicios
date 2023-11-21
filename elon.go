package exercicios

import "fmt"

// Car represents a remote controlled car.
type Car1 struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// NewCar creates and initializes a new Car instance.
func NewCar1(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
}

// Drive simulates driving the car for a certain distance.
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

// DisplayDistance returns the distance driven as a string.
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns the remaining battery charge as a string.
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish checks if the car can finish a given race track distance without running out of battery.
func (c *Car) CanFinish(trackDistance int) bool {
	possibleDistance := c.battery / c.batteryDrain * c.speed
	return possibleDistance >= trackDistance
}

func main1() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)

	// Drive the car and display status
	car.Drive()
	fmt.Println(car.DisplayDistance())
	fmt.Println(car.DisplayBattery())

	// Check if the car can finish a certain track distance
	trackDistance := 100
	fmt.Println(car.CanFinish(trackDistance))
}
