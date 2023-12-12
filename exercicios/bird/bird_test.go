package exercicios

import (
	"testing"
)

func TestTotalBirdCount(t *testing.T) {
	birdsPerDay := []int{2, 3, 5, 1, 6} // exemplo de dados para teste
	expectedTotal := 17
	calculatedTotal := TotalBirdCount(birdsPerDay)

	if expectedTotal != calculatedTotal {
		t.Errorf("Esperado %d, Resultado %d", expectedTotal, calculatedTotal)
	}
}
