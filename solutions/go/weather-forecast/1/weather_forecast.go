// Package weather provides one function to
// check for weather conditions in a city.
package weather

// CurrentCondition and CurrentLocation store
// locations and contions in strings to be
// used later on.
var (
	CurrentCondition string
	CurrentLocation  string
)

// Forecast returns the current condition of
// the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
