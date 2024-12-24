package main

import "fmt"

func main() {
	fmt.Println("===PROFIT CALCULATOR===")

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	profit, earningBeforeTax, tax := calculateProfit(revenue, expenses, taxRate)
	ratio := profit / earningBeforeTax

	outputText("===RESULTS===")
	outputText(fmt.Sprintf("Earning before tax (EBT): %.2f", earningBeforeTax))
	outputText(fmt.Sprintf("Earning after tax: %.2f", profit))
	outputText(fmt.Sprintf("Tax: %.2f", tax))
	outputText(fmt.Sprintf("Ratio:%.2f", ratio))
}

func getUserInput(text string) float64 {
	var input float64
	fmt.Print(text)
	fmt.Scan(&input)
	return input
}

func outputText(text string) {
	fmt.Println(text)
}

func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	tax := revenue * (taxRate / 100)
	profit := earningBeforeTax - tax

	return profit, earningBeforeTax, tax
}

// Another alternative
// declare the variables in the return statement and allow go to created them
// inside the function you just need to assign the values and return

// func calculateProfit(revenue, expenses, taxRate float64) (earningBeforeTax float64, tax float64, profit float64) {
// 	earningBeforeTax = revenue - expenses
// 	tax = revenue * (taxRate / 100)
// 	profit = earningBeforeTax - tax

// 	return
// }
