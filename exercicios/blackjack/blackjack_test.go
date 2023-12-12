package exercicios

import "testing"


func TestParseCard(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"ace", 11},
		{"dois", 2},
		{"três", 3},
		{"quatro", 4},
		{"cinco", 5},
		{"seis", 6},
		{"sete", 7},
		{"oito", 8},
		{"nove", 9},
		{"dez", 10},
		{"jack", 10},
		{"queen", 10},
		{"king", 10},
		{"ás", 11},
		{"treze", 0},
	}

	for _, test := range tests {
		if output := ParseCard(test.input); output != test.expected {
			t.Errorf("Para a entrada %s, esperava-se %d, mas obteve-se %d", test.input, test.expected, output)
		}
	}
}