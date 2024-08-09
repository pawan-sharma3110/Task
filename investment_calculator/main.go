package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectReturnRate float64 = 5.5
	var year float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&year)

	var futureValue = float64(investmentAmount) * math.Pow(1+expectReturnRate/100, year)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, year)

	fmt.Printf("Future Value: %.0f \n", futureValue)
	fmt.Printf("Future Real Value: %.0f \n", futureRealValue)
}
