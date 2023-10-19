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
	tests := []struct {
		knightIsAwake   bool
		archerIsAwake   bool
		prisonerIsAwake bool
		expected        bool
	}{
		{false, false, true, true},
		{true, true, true, true},
		{false, true, false, true},
		{true, false, false, true},
	}

	for _, test := range tests {
		result := CanSpy(test.knightIsAwake, test.archerIsAwake, test.prisonerIsAwake)
		if result != test.expected {
			t.Errorf("Para knightIsAwake=%v, archerIsAwake=%v, prisonerIsAwake=%v, esperado %v, mas obtido %v", test.knightIsAwake, test.archerIsAwake, test.prisonerIsAwake, test.expected, result)
		}
	}
}

func TestCanSignalPrisoner(t *testing.T) {
	tests := []struct {
		archerIsAwake   bool
		prisonerIsAwake bool
		expected        bool
	}{
		{false, true, true},
		{true, true, false},
		{false, false, false},
		{true, false, false},
	}

	for _, test := range tests {
		result := CanSignalPrisoner(test.archerIsAwake, test.prisonerIsAwake)
		if result != test.expected {
			t.Errorf("Para archerIsAwake=%v, prisonerIsAwake=%v, esperado %v, mas obtido %v", test.archerIsAwake, test.prisonerIsAwake, test.expected, result)
		}
	}
}

func TestCanFreePrisoner(t *testing.T) {
	tests := []struct {
		knightIsAwake        bool
		archerIsAwake        bool
		prisonerIsAwake      bool
		annalynsDogIsPresent bool
		expected             bool
	}{
		{true, true, true, true, false},
		{true, false, true, false, false},
		{false, true, true, true, false},
		{false, true, false, true, true},
		{false, false, true, false, true},
		{true, true, true, true, false},
	}

	for _, test := range tests {
		result := CanFreePrisoner(test.knightIsAwake, test.archerIsAwake, test.prisonerIsAwake, test.annalynsDogIsPresent)
		if result != test.expected {
			t.Errorf("Para knightIsAwake=%v, archerIsAwake=%v, prisonerIsAwake=%v, annalynsDogIsPresent=%v, esperado %v, mas obtido %v", test.knightIsAwake, test.archerIsAwake, test.prisonerIsAwake, test.annalynsDogIsPresent, test.expected, result)
		}
	}
}