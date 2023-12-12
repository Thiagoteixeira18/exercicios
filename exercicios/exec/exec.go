package exercicios

import (
	"errors"
	"fmt"
)

func exec() (int, int) {
	ex := 10
	teste := 5
	return ex, teste
}

func exec2() error {
	ex, teste := exec()
	total := 15

	if ex != total {
		return errors.New("o valor de 'ex' foi diferente do esperado")
	}
	if teste != total {
		return errors.New("o valor de 'teste' foi diferente do esperado")
	}
	return nil
}

func exec3() {
	err := exec2()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valores esperados coincidem com os valores obtidos.")
	}
}
