package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var workingCarsPerHour float64;
    
    workingCarsPerHour = float64(productionRate) * (successRate / 100);
    return (workingCarsPerHour);
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var workingCarsPerMinute int;
    var workingCarsPerHour int;
    
    workingCarsPerHour = int(CalculateWorkingCarsPerHour(productionRate, successRate));
    workingCarsPerMinute =  workingCarsPerHour/ 60;
    return (workingCarsPerMinute);
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	var cost uint;

    cost = uint(carsCount / 10) * 95000;
    cost = cost + uint(carsCount % 10) * 10000;
    return (cost);
}

