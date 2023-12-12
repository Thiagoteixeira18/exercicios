package exercicios

import "fmt"

func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, count := range birdsPerDay {
		total += count
	}
	return total
}

func BirdsInWeek(birdsPerDay []int, weekNumber int) int {
	startIndex := (weekNumber - 1) * 7
	endIndex := startIndex + 7
	if endIndex > len(birdsPerDay) {
		endIndex = len(birdsPerDay)
	}
	total := 0
	for i := startIndex; i < endIndex; i++ {
		total += birdsPerDay[i]
	}
	return total
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i == 0 || i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}

func main4() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	fmt.Println(TotalBirdCount(birdsPerDay))
	fmt.Println(BirdsInWeek(birdsPerDay, 2))
	fmt.Println(FixBirdCountLog(birdsPerDay))
}