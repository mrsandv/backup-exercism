package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerHourFloat := float64(productionRate) * successRate / 100
	carsPerMinute := carsPerHourFloat / 60
	return int(carsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const cost = 10000
	const costPer10units = 95000
	integerPart := int(carsCount / 10)
	decimalPart := carsCount % 10

	return uint(costPer10units*integerPart + cost*decimalPart)
}
