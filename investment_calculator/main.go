package main

import (
	"fmt"
	"math"
)

func main() {

	const INFLATION_RATE = 2.6
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	valueAfterInflation := futureValue / math.Pow(1+INFLATION_RATE/100, years)

	fmt.Println("Future Value:", futureValue)
	fmt.Println("Value After Inflation:", valueAfterInflation)
}
