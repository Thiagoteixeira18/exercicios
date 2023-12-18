package main

import "fmt"

func clouse() func() {
	texto := "exercicio de closure"
	numero := 50
	funcao := func() {
		fmt.Println(texto)
		fmt.Println(numero)
	}
	return funcao
}
func main() {
	numero := 20
	texto := "exercicio"
	fmt.Println(texto)
	fmt.Println(numero)

	funcaonova := clouse()
	funcaonova()
}
