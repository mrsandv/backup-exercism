// Package weather is the pricipal package for this app.
package weather

var (
	// CurrentCondition is a variable that show the actual weather condition.
	CurrentCondition string
	// CurrentLocation is a variable that point the location of the weather information showed.
	CurrentLocation string
)

// Forecast calculates the weather conditions in order to get the current values on the city selected.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
