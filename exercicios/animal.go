package exercicios

import (
	"fmt"
	"math/rand"
	"time"
)

func RollADie() int {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	return rand.Intn(20) + 1 // add 1 to ensure the range is 1 to 20
}

func GenerateWandEnergy() float64 {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	return rand.Float64() * 12.0 // generates a float between 0.0 and 12.0
}

func ShuffleAnimals() []string {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}

func main8() {
	fmt.Println("Rolling a die:", RollADie())
	fmt.Println("Generating wand energy:", GenerateWandEnergy())
	fmt.Println("Shuffling animals:", ShuffleAnimals())
}