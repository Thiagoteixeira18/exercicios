package exercicios

import "testing"

func TestCalculateWorkingCarsPerHour(t *testing.T) {
	productionRate := 100
	successRate := 85.0
	expectedResult := 85.0

	result := CalculateWorkingCarsPerHour(productionRate, successRate)

	if result != expectedResult {
		t.Errorf("Expected %f, but got %f", expectedResult, result)
	}
}
