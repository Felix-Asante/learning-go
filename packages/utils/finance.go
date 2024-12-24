package utils

import (
	"fmt"
)

const BalanceFileName = "balance.txt"

func WithrawFromAmount(accountBalance float64) float64 {
	withdrawAmount := getAmount("Amount to withdraw: ")
	if withdrawAmount > accountBalance {
		fmt.Println("Insufficient balance")
		return accountBalance
	} else if withdrawAmount <= 0 {
		fmt.Println("Invalid amount. Amount must be greater than 0")
		return accountBalance
	}
	accountBalance = calculateNewBalance(accountBalance, withdrawAmount, "withdraw")
	WriteValueToFile(BalanceFileName, accountBalance)
	fmt.Println("Your balance is:", accountBalance)
	return accountBalance
}

func DepositToAccount(accountBalance float64) float64 {
	amount := getAmount("Amount to deposit: ")
	if amount <= 0 {
		fmt.Println("Invalid amount. Amount must be greater than 0")
		return accountBalance
	}
	accountBalance = calculateNewBalance(accountBalance, amount, "deposit")
	WriteValueToFile(BalanceFileName, accountBalance)
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
