package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectReturnRate = 5.5
	var year float64 = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectReturnRate/100, year)
	fmt.Println(futureValue)
}
