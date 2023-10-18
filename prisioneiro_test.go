package exercicios

import (
	"testing"
)

func TestCanFastAttack(t *testing.T) {
	tests := []struct {
		knightIsAwake bool
		expected      bool
	}{
		{true, false},
		{false, true},
	}

	for _, test := range tests {
		result := CanFastAttack(test.knightIsAwake)
		if result != test.expected {
			t.Errorf("Para knightIsAwake=%v, esperado %v, mas obtido %v", test.knightIsAwake, test.expected, result)
		}
	}
}

func TestCanSpy(t *testing.T) {
	teste := []struct {
		knightIsAwake   bool
		archerIsAwake   bool
		prisonerIsAwake bool
		esperado        bool
	}{
		{false, false, true, false},
		{true, true, true, true},
	}
	for _, test := range teste {
		result := (test.knightIsAwake)
		if result != test.esperado {
			t.Errorf("Para canspy=%v, esperado %v, mas obtido %v", test.knightIsAwake, test.esperado, result)
		}
	}
}
