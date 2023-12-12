package exercicios

import (
	"testing"
)

// Adicione aqui o código das funções fornecidas

func TestWelcome(t *testing.T) {
	expected := "Welcome to my party, Christiane!"
	if got := Welcome("Christiane"); got != expected {
		t.Errorf("Welcome() = %q, want %q", got, expected)
	}
}

func TestHappyBirthday(t *testing.T) {
	expected := "Happy birthday Frank! You are now 58 years old!"
	if got := HappyBirthday("Frank", 58); got != expected {
		t.Errorf("HappyBirthday() = %q, want %q", got, expected)
	}
}

func TestAssignTable(t *testing.T) {
	expected := "Welcome to my party, Christiane!\nYou have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.\nYou will be sitting next to Frank."
	if got := AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298); got != expected {
		t.Errorf("AssignTable() = %q, want %q", got, expected)
	}
}