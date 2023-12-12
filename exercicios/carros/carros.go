package exercicios

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    successPercentage := successRate / 100.0
    return float64(productionRate) * successPercentage
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    successPercentage := successRate / 100.0
    carsPerHour := float64(productionRate) * successPercentage
    return int(carsPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    costPerGroup := uint(95000)
    costPerIndividual := uint(10000)
    
    groups := carsCount / 10
    individualCars := carsCount % 10
    
    totalCost := uint(groups) * costPerGroup + uint(individualCars) * costPerIndividual
    return totalCost
}
