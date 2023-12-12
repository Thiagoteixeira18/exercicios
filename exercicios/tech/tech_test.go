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
func TestAddBorder(t *testing.T) {
	welcomeMsg := "Welcome to OpenAI"
	numStarsPerLine := 5
	expectedResult := "*****\nWelcome to OpenAI\n*****"
	result := AddBorder(welcomeMsg, numStarsPerLine)
	if result != expectedResult {
		t.Errorf("Expected %s, but got %s", expectedResult, result)
	}
}

func TestCleanupMessage(t *testing.T){
	oldMsg := "*****\nola mundo\n*****"
	expectedResult := "ola mundo"
	result := CleanupMessage(oldMsg)
	if result != expectedResult {
		t.Errorf("esperado %s, but got %s", expectedResult, result)
	}
}



