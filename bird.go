package exercicios

import (
	"fmt"
)

func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, count := range birdsPerDay {
		total += count
	}
	return total
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	startIndex := (week - 1) * 7
	endIndex := week * 7
	total := 0
	for i := startIndex; i < endIndex; i++ {
		total += birdsPerDay[i]
	}
	return total
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 1; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}

func main4() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}

	// Test TotalBirdCount
	totalBirds := TotalBirdCount(birdsPerDay)
	fmt.Println("Total Bird Count:", totalBirds)

	// Test BirdsInWeek
	weekNumber := 2
	birdsInWeek := BirdsInWeek(birdsPerDay, weekNumber)
	fmt.Printf("Birds in Week %d: %d\n", weekNumber, birdsInWeek)

	// Test FixBirdCountLog
	fixedBirdsPerDay := FixBirdCountLog(birdsPerDay)
	fmt.Println("Fixed Bird Count Log:", fixedBirdsPerDay)
}
