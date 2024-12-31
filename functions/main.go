package main

import "fmt"

type transformCb func(int) int

func main() {
	fmt.Println("Annonymous functions in Go")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Annonymous function
	transformedNumbers := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	dbFnc := func(number int) int {
		return number * 2
	}
	transformedNumbers2 := transformNumbers(&numbers, dbFnc)
	// with closures
	tripleFnc := createTransformer(3)
	transformedNumbers3 := transformNumbers(&numbers, tripleFnc)

	fmt.Println(numbers)
	fmt.Println(transformedNumbers)
	fmt.Println(transformedNumbers2)
	fmt.Println(transformedNumbers3)
}

func transformNumbers(numbers *[]int, transformer transformCb) []int {
	transformedValues := []int{}
	for _, number := range *numbers {
		transformedValues = append(transformedValues, transformer(number))
	}
	return transformedValues
}

// closure

func createTransformer(factor int) transformCb {
	return func(number int) int {
		return number * factor
	}
}
