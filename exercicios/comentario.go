package exercicios
// Package weather provides functionality to retrieve and display information about current weather conditions. 

// CurrentCondition stores the current weather condition.
var CurrentCondition string

// CurrentLocation stores the current location for weather information.
var CurrentLocation string

// Forecast function forecasts the current weather condition for a given city. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
