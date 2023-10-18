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
