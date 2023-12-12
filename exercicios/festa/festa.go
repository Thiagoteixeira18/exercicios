package exercicios

import (
	"fmt"
)

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func AssignTable(name string, tableNumber int, companion string, direction string, distance float64) string {
	tableNumberStr := fmt.Sprintf("%03d", tableNumber)
	distanceStr := fmt.Sprintf("%.1f", distance)

	message := fmt.Sprintf("Welcome to my party, %s!\n", name)
	message += fmt.Sprintf("You have been assigned to table %s. Your table is %s, exactly %s meters from here.\n", tableNumberStr, direction, distanceStr)
	message += fmt.Sprintf("You will be sitting next to %s.", companion)

	return message
}

func main() {
	fmt.Println(Welcome("Christiane"))
	fmt.Println(HappyBirthday("Frank", 58))
	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}
