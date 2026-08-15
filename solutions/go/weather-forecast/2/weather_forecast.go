// Package weather provides one function to
// check for weather conditions in a city.
package weather

var (
	// CurrentCondition stores the current condition in a string for usage.
	CurrentCondition string
	// CurrentLocation stores the current location for later usage.
	CurrentLocation string
)

// Forecast returns the current condition of
// the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
