package first_version

import (
	"fmt"
	"math"
)

func main(){
	
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow((1 + expectedReturnRate/100) , float64(years))

	fmt.Println("Future Value:" ,futureValue)
	fmt.Println("RoundToEven number:" ,math.RoundToEven((futureValue)))
	fmt.Println("Round number:" ,math.Round((futureValue*1000))/1000) // round to 3 decimal places
	fmt.Println("Ceil number:" ,math.Ceil(futureValue))
	fmt.Println("Whole number:" ,math.Floor(futureValue))
	fmt.Println("Exp number:" ,math.Exp(futureValue))
	fmt.Println("Cube root number:" ,math.Cbrt(futureValue))
	fmt.Println("Trunc number:" ,math.Trunc(futureValue))
	var formattedNumber = fmt.Sprintf("Format number: %.3f" ,futureValue) // format to 3 decimal places
	fmt.Println("formatted number:",formattedNumber)
}