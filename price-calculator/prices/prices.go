package prices

import (
	"fmt"

	"example.com/price-calculator/conversions"
	"example.com/price-calculator/fileManager"
)

type TaxIncludedPriceJob struct {
	InputPrices       []float64
	TaxRate           float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadPrices()
	results := make(map[string]float64)

	for _, price := range job.InputPrices {
		results[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	job.TaxIncludedPrices = results
	fileName := fmt.Sprintf("results_%.0f.json", job.TaxRate*100)
	fileManager.WriteJSON(fileName, job)
	fmt.Println(results)
}

func (job *TaxIncludedPriceJob) LoadPrices() {
	lines, error := fileManager.ReadLines("prices.txt")

	if error != nil {
		fmt.Println("Failed to open file", error)
		return
	}

	prices := []float64{}

	for _, line := range lines {
		price := conversions.StringToFloat64(line)

		prices = append(prices, price)
	}

	job.InputPrices = prices
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxIncludedPrices: make(map[string]float64),
		InputPrices:       []float64{10, 30, 20},
		TaxRate:           taxRate,
	}
}
