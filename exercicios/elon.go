package exercicios

import "fmt"

type Car2 struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

func NewCar2(speed, batteryDrain int) *Car2 {
	return &Car2{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
}

func (c *Car2) Drive() {
	if c.battery >= c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

func (c *Car2) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c *Car2) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c *Car2) CanFinish(trackDistance int) bool {
	distanceLeft := trackDistance - c.distance
	batteryRequired := distanceLeft * c.batteryDrain / c.speed
	return batteryRequired <= c.battery
}

func main44() {
	speed := 5
	batteryDrain := 2
	car := NewCar2(speed, batteryDrain)

	car.Drive()
	fmt.Println(car.DisplayDistance())
	fmt.Println(car.DisplayBattery())

	trackDistance := 100
	fmt.Println(car.CanFinish(trackDistance))
}