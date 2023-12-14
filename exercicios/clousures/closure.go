package main

import "fmt"

func clouse() func() {
	texto := "exercicio de closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}
func main() {
	texto := "exercicio"
	fmt.Println(texto)

	funcaonova := clouse()
	funcaonova()
}
