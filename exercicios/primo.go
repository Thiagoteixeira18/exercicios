package exercicios

import (
	"fmt"
)

type numeros []int

func NumerosPrimos() {
	x := numeros{1, 2, 3, 4, 5, 6, 7, 8}

	for _, num := range x {
		if num > 1 && isPrimo(num) {
			fmt.Printf("%d é um número primo\n", num)
		} else {
			fmt.Printf("%d não é um número primo\n", num)
		}
	}
}

func isPrimo(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}


func notas() {
	materias := []string{"Matemática", "Ciências", "História"}
	notas := map[string]int{
		"Matemática": 6,
		"Ciências":   7,
		"História":   5,
	}
	base := 6

	for _, materia := range materias {
		nota, found := notas[materia]
		if !found {
			fmt.Println("Nota não encontrada para", materia)
			continue
		}

		if nota < base {
			fmt.Println("Reprovado em", materia)
		} else {
			fmt.Println("Aprovado em", materia)
		}
	}
}
