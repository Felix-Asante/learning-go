package main

import (
	"fmt"
	"os"
	"strconv"
)

const balanceFileName = "balance.txt"

func main() {

	var accountBalance float64 = getBalanceFromFile()
	endApp := false

	fmt.Println("Welcome to Go bank !")

	for {
		fmt.Println("What do you want to do ?")
		fmt.Println("1.Check your balance")
		fmt.Println("2.Deposit money")
		fmt.Println("3.Withdraw money")
		fmt.Println("4.Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			accountBalance = performDeposit(accountBalance)
		case 3:
			accountBalance = withrawAmount(accountBalance)
		default:
			fmt.Println("Goodbye !")
			endApp = true
		}

		if endApp {
			break
		}
	}

	fmt.Println("Thanks for choosing our bank !")

}

func withrawAmount(accountBalance float64) float64 {
	withdrawAmount := getAmount("Amount to withdraw: ")
	if withdrawAmount > accountBalance {
		fmt.Println("Insufficient balance")
		return accountBalance
	} else if withdrawAmount <= 0 {
		fmt.Println("Invalid amount. Amount must be greater than 0")
		return accountBalance
	}
	accountBalance = calculateNewBalance(accountBalance, withdrawAmount, "withdraw")
	writeBalanceToFile(accountBalance)
	fmt.Println("Your balance is:", accountBalance)
	return accountBalance
}

func performDeposit(accountBalance float64) float64 {
	amount := getAmount("Amount to deposit: ")
	if amount <= 0 {
		fmt.Println("Invalid amount. Amount must be greater than 0")
		return accountBalance
	}
	accountBalance = calculateNewBalance(accountBalance, amount, "deposit")
	writeBalanceToFile(accountBalance)
	fmt.Println("Your balance is:", accountBalance)
	return accountBalance
}

func calculateNewBalance(accountBalance float64, amount float64, action string) float64 {
	if action == "withdraw" {
		return accountBalance - amount
	}
	return accountBalance + amount
}

func getAmount(text string) float64 {
	var amount float64
	fmt.Print(text)
	fmt.Scan(&amount)
	return amount
}

func writeBalanceToFile(accountBalance float64) {
	error := os.WriteFile(balanceFileName, []byte(fmt.Sprintf("%.2f", accountBalance)), 0644)

	if error != nil {
		fmt.Println("Error writing balance to file")
	}
}

func getBalanceFromFile() float64 {
	balance, error := os.ReadFile(balanceFileName)
	if error != nil {
		fmt.Println("Error reading balance from file")
		return 0
	}
	balanceFloat, err := strconv.ParseFloat(string(balance), 64)
	if err != nil {
		fmt.Println("Error converting balance to float")
		return 0
	}
	return balanceFloat
}
