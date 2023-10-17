package exercicios

import "testing"

/// TestWelcomeMessage tests the WelcomeMessage function.
func TestWelcomeMessage(t *testing.T) {
	customer := "John"
	want := "Welcome to the Tech Palace, JOHN"
	got := WelcomeMessage(customer)

	if got != want {
		t.Errorf("WelcomeMessage(%s) = %s; want %s", customer, got, want)
	}
}