package starter

import "fmt"

type transformCb func(int) int

func main() {
	fmt.Println("Advanced functions in Go")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	moreNumbers := []int{12, 10, 1, 5}

	// dbNumbers := doubleNumbers(&numbers)

	dbNumbers := transformNumbers(&numbers, double)
	trpNumbers := transformNumbers(&numbers, triple)
	transformedNumbers := transformNumbers(&moreNumbers, getTransformerFunction("double"))

	fmt.Println(numbers)
	fmt.Println(trpNumbers)
	fmt.Println(dbNumbers)
	fmt.Println(transformedNumbers)
}

func doubleNumbers(numbers *[]int) []int {
	dbNumbers := []int{}
	for _, number := range *numbers {
		dbNumbers = append(dbNumbers, double(number))
	}
	return dbNumbers
}
func transformNumbers(numbers *[]int, cb transformCb) []int {
	transformedNumbers := []int{}
	for _, number := range *numbers {
		transformedNumbers = append(transformedNumbers, cb(number))
	}
	return transformedNumbers
}

// function can return other functions
func getTransformerFunction(transformType string) transformCb {
	if transformType == "double" {
		return double
	}
	return triple
}
func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
