package exercicios

import (
	"fmt"
)

// Função para estimar o tempo total de preparação com base no número de camadas e no tempo médio de preparação por camada em minutos
func PreparationTime1(layers []string, avgTimePerLayer int) int {
	if avgTimePerLayer == 0 {
		avgTimePerLayer = 1
	}

	return len(layers) * avgTimePerLayer
}

// Função para determinar a quantidade de macarrão e molho necessária com base nas camadas
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 51
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// Função para adicionar o ingrediente secreto à lista de ingredientes
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// Função para escalar as quantidades necessárias para o número desejado de porções
func ScaleRecipe(quantities []float64, servings int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i, qty := range quantities {
		scaledQuantities[i] = qty * float64(servings) / 2
	}
	return scaledQuantities
}

func main7() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println("Preparation time:", PreparationTime1(layers, 3))
	fmt.Println("Preparation time with default avgTimePerLayer:", PreparationTime1(layers, 0))

	noodles, sauce := Quantities(layers)
	fmt.Println("Noodles needed:", noodles, "grams")
	fmt.Println("Sauce needed:", sauce, "liters")

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	AddSecretIngredient(friendsList, myList)
	fmt.Println("My list with secret ingredient:", myList)

	quantities := []float64{1.2, 3.6, 10.4}
	scaledQuantities := ScaleRecipe(quantities, 4)
	fmt.Println("Scaled quantities:", scaledQuantities)
}
