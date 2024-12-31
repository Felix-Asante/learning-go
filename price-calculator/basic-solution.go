package main

import "fmt"

func runBasic() {
	prices := []float64{10, 30, 20}
	taxRates := []float64{0.1, 0.2, 0.3}

	results := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		var taxIncludedPrice []float64 = make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrice[priceIndex] = price * (1 + taxRate)
		}
		results[taxRate] = taxIncludedPrice
	}

	fmt.Println(results)
}
