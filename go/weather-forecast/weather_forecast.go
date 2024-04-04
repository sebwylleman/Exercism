// Package weather provides tools to give a current weather forecast for cities in Goblinocus.
package weather

// CurrentCondition represents a certain condition for the current weather.
var CurrentCondition string

// CurrentLocation represents the current location at where the weather forecast is made.
var CurrentLocation string

// Forecast returns a weather forecast given a city and a condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
