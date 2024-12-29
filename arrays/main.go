package main

import (
	"example.com/structures/maps"
)

func main() {
	prices := []float64{10.99, 8.99}

	prices[1] = 12.99
	// prices[3] = 4.6 // will throw error cos the index is out of range

	updatedPrice := append(prices, 5.99)
	updatedPrice[0] = 9.99
	// fmt.Println(prices)
	// fmt.Println(updatedPrice)
	// practice.RunPractice()
	maps.RunMaps()

}
