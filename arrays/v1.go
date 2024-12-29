package main

import "fmt"

func mainFunc() {
	prices := [5]float64{1.99, 2.99, 3.99, 4.99, 5.99}
	var productNames [4]string = [4]string{"apple", "banana", "orange", "grape"}

	// productNames = [4]string{"apple", "banana", "orange", "grape"}
	fmt.Println(prices, productNames)
	fmt.Println("3rd value", productNames[2])
	productNames[2] = "mango"
	fmt.Println("3rd value after mutation", productNames[2])

	fmt.Println("Price slices")
	featuredPrice := prices[1:3] // slice: start index 1 and end index 2
	fmt.Println(featuredPrice)
	featuredPrice = prices[:3] // slice: start index 0 and end index 2
	fmt.Println(featuredPrice)
	featuredPrice = prices[1:] // slice: start index 1 and end length-1
	fmt.Println(featuredPrice)
	highlightedPrice := featuredPrice[:2] // slice: start index 0 and end index 0
	fmt.Println(highlightedPrice)
	// mutating a slice affect the original array
	// slices serve as a reference to the original array, array is not copied when u created a slice
	highlightedPrice[0] = 34
	fmt.Println(highlightedPrice)
	fmt.Println(featuredPrice)
	fmt.Println(prices)

	fmt.Println("Length of prices", len(prices))         // length is the size of the array (total items in the array)
	fmt.Println("Capacity of prices", cap(productNames)) // capacity is  the maximum number of elements it can hold without resizing

}
