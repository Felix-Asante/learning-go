package main

import "fmt"

func main() {
	age := 25
	// var agePointer *int = &age
	agePointer := &age

	fmt.Println("Age:", age)
	fmt.Println("Adult years:", getAdultYears(agePointer)) // pass by reference
	fmt.Println("Age:", age)
}

func getAdultYears(age *int) int {
	// return *age - 18
	*age = *age - 18 // mutate the value of the variable
	return *age
}
