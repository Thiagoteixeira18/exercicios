package main

import "fmt"

func RecuperarExecução() {
	if r := recover(); r != nil {
		fmt.Println("A execução recuperada com sucesso")
	}
}

func AlunoEstaAprovado(n1, n2 float64) bool {
	defer RecuperarExecução()

	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("a media e exatamente 6")
}
