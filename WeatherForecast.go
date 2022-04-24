// Package weather provides tools to forecast 
// weather.
package weather
// CurrentCondition represent certain weather condition in a city.
var CurrentCondition string
// CurrentLocation represent the city we forcasted the weather for.
var CurrentLocation string

// Forecast return string that represent the city and its current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

